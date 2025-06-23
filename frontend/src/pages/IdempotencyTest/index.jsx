import React, { useState, useEffect } from 'react';
import { Card, Button, Space, message, Typography, Divider, Alert, Tag } from 'antd';
import {
  checkUserUnpaidOrder,
  checkBeforeReservation,
  checkBeforeOrder,
  showUnpaidOrderNotification
} from '../../utils/idempotencyUtils';
import orderService from '../../services/orderService';

const { Title, Text } = Typography;

const IdempotencyTest = () => {
  const [loading, setLoading] = useState(false);
  const [checkResult, setCheckResult] = useState(null);
  const [loginStatus, setLoginStatus] = useState(null);

  // 检查登录状态
  useEffect(() => {
    const token = localStorage.getItem('token');
    setLoginStatus(token ? '已登录' : '未登录');
  }, []);

  // 测试检查未支付订单
  const handleCheckUnpaidOrder = async () => {
    setLoading(true);
    try {
      const result = await checkUserUnpaidOrder();
      setCheckResult(result);
      message.success('检查完成');
    } catch (error) {
      message.error('检查失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 测试预订前检查
  const handleTestReservation = async () => {
    await checkBeforeReservation(() => {
      message.success('预订检查通过，可以进行预订！');
    });
  };

  // 测试下单前检查
  const handleTestOrder = async () => {
    await checkBeforeOrder(() => {
      message.success('下单检查通过，可以创建订单！');
    });
  };

  // 测试显示未支付订单通知
  const handleShowNotification = () => {
    if (checkResult && checkResult.hasUnpaidOrder) {
      showUnpaidOrderNotification(checkResult.unpaidOrder);
    } else {
      message.warning('没有未支付订单可以显示');
    }
  };

  // 直接调用API测试
  const handleDirectApiTest = async () => {
    setLoading(true);
    try {
      const response = await orderService.checkUnpaidOrder();
      console.log('API响应:', response);
      message.success('API调用成功，请查看控制台');
    } catch (error) {
      message.error('API调用失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ padding: '24px', maxWidth: '800px', margin: '0 auto' }}>
      <Title level={2}>接口幂等性测试页面</Title>
      <Text type="secondary">
        此页面用于测试订单接口的幂等性功能，确保用户在有未支付订单时不能创建新的预订或订单。
      </Text>

      <Divider />

      <Card title="基础检查功能" style={{ marginBottom: 16 }}>
        <Space direction="vertical" style={{ width: '100%' }}>
          <Button 
            type="primary" 
            onClick={handleCheckUnpaidOrder}
            loading={loading}
            block
          >
            检查用户未支付订单
          </Button>
          
          <Button 
            onClick={handleDirectApiTest}
            loading={loading}
            block
          >
            直接调用API测试
          </Button>

          {checkResult && (
            <Card size="small" style={{ backgroundColor: '#f6f6f6' }}>
              <Text strong>检查结果：</Text>
              <br />
              <Text>有未支付订单：{checkResult.hasUnpaidOrder ? '是' : '否'}</Text>
              <br />
              {checkResult.hasUnpaidOrder && checkResult.unpaidOrder && (
                <>
                  <Text>订单号：{checkResult.unpaidOrder.order_sn}</Text>
                  <br />
                  <Text>订单金额：¥{checkResult.unpaidOrder.total_amount}</Text>
                  <br />
                  <Text>创建时间：{new Date(checkResult.unpaidOrder.created_at).toLocaleString()}</Text>
                </>
              )}
              <br />
              <Text>消息：{checkResult.message}</Text>
            </Card>
          )}
        </Space>
      </Card>

      <Card title="幂等性检查测试" style={{ marginBottom: 16 }}>
        <Space direction="vertical" style={{ width: '100%' }}>
          <Button 
            type="primary" 
            onClick={handleTestReservation}
            block
          >
            测试预订前检查
          </Button>
          
          <Button 
            type="primary" 
            onClick={handleTestOrder}
            block
          >
            测试下单前检查
          </Button>

          <Button 
            onClick={handleShowNotification}
            block
          >
            显示未支付订单通知
          </Button>
        </Space>
      </Card>

      <Card title="使用说明" size="small">
        <Space direction="vertical">
          <Text>1. 点击"检查用户未支付订单"查看当前用户是否有未支付的订单</Text>
          <Text>2. 点击"测试预订前检查"模拟用户尝试预订车辆的场景</Text>
          <Text>3. 点击"测试下单前检查"模拟用户尝试创建订单的场景</Text>
          <Text>4. 如果有未支付订单，系统会阻止新的预订/下单操作</Text>
          <Text>5. 用户需要先完成当前订单的支付才能进行新的操作</Text>
        </Space>
      </Card>
    </div>
  );
};

export default IdempotencyTest;
