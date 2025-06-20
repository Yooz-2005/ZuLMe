import React, { useState } from 'react';
import { Card, Button, Form, Input, Select, message, Space, Typography, Divider } from 'antd';
import { CreditCardOutlined, CheckCircleOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import orderService from '../../services/orderService';
import reservationService from '../../services/reservationService';

const { Title, Text } = Typography;
const { Option } = Select;

const Container = styled.div`
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
`;

const TestCard = styled(Card)`
  margin-bottom: 20px;
  
  .ant-card-head {
    background: #f8f9fa;
  }
`;

const PaymentTest = () => {
  const [loading, setLoading] = useState(false);
  const [testResults, setTestResults] = useState([]);

  // 测试创建预订
  const testCreateReservation = async () => {
    setLoading(true);
    try {
      const reservationData = {
        vehicle_id: 1,
        start_date: '2024-01-20',
        end_date: '2024-01-22',
        pickup_location_id: 1,
        total_amount: 1200.00,
        notes: '测试预订'
      };

      const response = await reservationService.createReservation(reservationData);
      
      if (response && response.code === 200) {
        message.success('预订创建成功');
        setTestResults(prev => [...prev, {
          type: 'success',
          title: '预订创建测试',
          content: `预订ID: ${response.data.id}, 状态: ${response.data.status}`
        }]);
        return response.data;
      } else {
        throw new Error(response?.message || '创建预订失败');
      }
    } catch (error) {
      message.error('创建预订失败: ' + error.message);
      setTestResults(prev => [...prev, {
        type: 'error',
        title: '预订创建测试',
        content: error.message
      }]);
      return null;
    } finally {
      setLoading(false);
    }
  };

  // 测试从预订创建订单
  const testCreateOrderFromReservation = async (reservationId) => {
    setLoading(true);
    try {
      const orderData = {
        reservation_id: reservationId,
        return_location_id: 2,
        payment_method: 1, // 支付宝
        notes: '测试订单'
      };

      const response = await orderService.createOrderFromReservation(orderData);
      
      if (response && response.code === 200) {
        message.success('订单创建成功');
        setTestResults(prev => [...prev, {
          type: 'success',
          title: '订单创建测试',
          content: `订单号: ${response.data.order_sn}, 支付链接: ${response.data.payment_url ? '已生成' : '未生成'}`
        }]);
        return response.data;
      } else {
        throw new Error(response?.message || '创建订单失败');
      }
    } catch (error) {
      message.error('创建订单失败: ' + error.message);
      setTestResults(prev => [...prev, {
        type: 'error',
        title: '订单创建测试',
        content: error.message
      }]);
      return null;
    } finally {
      setLoading(false);
    }
  };

  // 测试完整支付流程
  const testFullPaymentFlow = async () => {
    setTestResults([]);
    
    // 1. 创建预订
    const reservation = await testCreateReservation();
    if (!reservation) return;

    // 2. 从预订创建订单
    const order = await testCreateOrderFromReservation(reservation.id);
    if (!order) return;

    // 3. 模拟支付
    if (order.payment_url) {
      setTestResults(prev => [...prev, {
        type: 'info',
        title: '支付链接测试',
        content: `支付链接已生成，可以点击测试支付: ${order.payment_url}`
      }]);
    }
  };

  // 测试获取订单列表
  const testGetOrders = async () => {
    setLoading(true);
    try {
      const response = await orderService.getUserOrders();
      
      if (response && response.code === 200) {
        message.success('获取订单列表成功');
        setTestResults(prev => [...prev, {
          type: 'success',
          title: '订单列表测试',
          content: `获取到 ${response.data?.orders?.length || 0} 个订单`
        }]);
      } else {
        throw new Error(response?.message || '获取订单列表失败');
      }
    } catch (error) {
      message.error('获取订单列表失败: ' + error.message);
      setTestResults(prev => [...prev, {
        type: 'error',
        title: '订单列表测试',
        content: error.message
      }]);
    } finally {
      setLoading(false);
    }
  };

  // 测试获取预订列表
  const testGetReservations = async () => {
    setLoading(true);
    try {
      const response = await reservationService.getUserReservations();
      
      if (response && response.code === 200) {
        message.success('获取预订列表成功');
        setTestResults(prev => [...prev, {
          type: 'success',
          title: '预订列表测试',
          content: `获取到 ${response.data?.reservations?.length || 0} 个预订`
        }]);
      } else {
        throw new Error(response?.message || '获取预订列表失败');
      }
    } catch (error) {
      message.error('获取预订列表失败: ' + error.message);
      setTestResults(prev => [...prev, {
        type: 'error',
        title: '预订列表测试',
        content: error.message
      }]);
    } finally {
      setLoading(false);
    }
  };

  const clearResults = () => {
    setTestResults([]);
  };

  return (
    <Container>
      <Title level={2}>支付流程测试</Title>
      <Text type="secondary">测试预订创建、订单生成、支付流程等功能</Text>

      <Divider />

      <TestCard title="测试功能">
        <Space wrap>
          <Button 
            type="primary" 
            icon={<CreditCardOutlined />}
            loading={loading}
            onClick={testFullPaymentFlow}
          >
            测试完整支付流程
          </Button>
          <Button 
            loading={loading}
            onClick={testGetReservations}
          >
            测试获取预订列表
          </Button>
          <Button 
            loading={loading}
            onClick={testGetOrders}
          >
            测试获取订单列表
          </Button>
          <Button onClick={clearResults}>
            清除结果
          </Button>
        </Space>
      </TestCard>

      {testResults.length > 0 && (
        <TestCard title="测试结果">
          <Space direction="vertical" style={{ width: '100%' }}>
            {testResults.map((result, index) => (
              <Card 
                key={index}
                size="small"
                style={{ 
                  borderLeft: `4px solid ${
                    result.type === 'success' ? '#52c41a' : 
                    result.type === 'error' ? '#ff4d4f' : '#1890ff'
                  }`
                }}
              >
                <div>
                  <Text strong>{result.title}</Text>
                  <br />
                  <Text>{result.content}</Text>
                </div>
              </Card>
            ))}
          </Space>
        </TestCard>
      )}

      <TestCard title="使用说明">
        <Space direction="vertical">
          <Text>1. 点击"测试完整支付流程"会依次创建预订、生成订单、生成支付链接</Text>
          <Text>2. 测试前请确保已登录并且后端服务正常运行</Text>
          <Text>3. 支付链接生成后可以点击进行模拟支付测试</Text>
          <Text>4. 支付成功后可以在个人中心查看订单状态变化</Text>
        </Space>
      </TestCard>
    </Container>
  );
};

export default PaymentTest;
