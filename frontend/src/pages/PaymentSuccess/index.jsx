import React, { useEffect, useState } from 'react';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { Result, Button, Card, Descriptions, Spin } from 'antd';
import { CheckCircleOutlined, HomeOutlined, UnorderedListOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import orderService from '../../services/orderService';

const Container = styled.div`
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px 20px;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const ContentCard = styled(Card)`
  max-width: 600px;
  width: 100%;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  
  .ant-card-body {
    padding: 40px;
  }
`;

const PaymentSuccess = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [orderInfo, setOrderInfo] = useState(null);

  const orderSn = searchParams.get('order_sn');
  const tradeNo = searchParams.get('trade_no');

  useEffect(() => {
    if (orderSn) {
      fetchOrderInfo();
    } else {
      setLoading(false);
    }
  }, [orderSn]);

  const fetchOrderInfo = async () => {
    try {
      const response = await orderService.getOrderDetailBySn(orderSn);
      if (response && response.code === 200) {
        setOrderInfo(response.data);
      }
    } catch (error) {
      console.error('获取订单信息失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleGoHome = () => {
    navigate('/');
  };

  const handleViewOrders = () => {
    // 清除可能的缓存，确保显示最新数据
    localStorage.removeItem('reservationListCache');
    localStorage.removeItem('orderListCache');
    navigate('/personal-center');
  };

  if (loading) {
    return (
      <Container>
        <ContentCard>
          <div style={{ textAlign: 'center', padding: '40px' }}>
            <Spin size="large" />
            <p style={{ marginTop: 16, color: '#666' }}>正在获取订单信息...</p>
          </div>
        </ContentCard>
      </Container>
    );
  }

  return (
    <Container>
      <ContentCard>
        <Result
          icon={<CheckCircleOutlined style={{ color: '#52c41a' }} />}
          status="success"
          title="支付成功！"
          subTitle={`您的订单已支付成功，我们将尽快为您安排车辆。${orderSn ? `订单号：${orderSn}` : ''}`}
          extra={[
            <Button type="primary" key="orders" icon={<UnorderedListOutlined />} onClick={handleViewOrders}>
              查看我的订单
            </Button>,
            <Button key="home" icon={<HomeOutlined />} onClick={handleGoHome}>
              返回首页
            </Button>,
          ]}
        />

        {orderInfo && (
          <div style={{ marginTop: 24 }}>
            <Descriptions title="订单详情" bordered column={1}>
              <Descriptions.Item label="订单号">{orderInfo.order_sn}</Descriptions.Item>
              {tradeNo && (
                <Descriptions.Item label="交易号">{tradeNo}</Descriptions.Item>
              )}
              <Descriptions.Item label="车辆信息">
                {orderInfo.vehicle_brand} {orderInfo.vehicle_style}
              </Descriptions.Item>
              <Descriptions.Item label="租赁时间">
                {orderInfo.start_date} 至 {orderInfo.end_date}
              </Descriptions.Item>
              <Descriptions.Item label="取车地点">{orderInfo.pickup_location}</Descriptions.Item>
              <Descriptions.Item label="支付金额">¥{orderInfo.total_amount}</Descriptions.Item>
              <Descriptions.Item label="订单状态">
                <span style={{ color: '#52c41a', fontWeight: 'bold' }}>已支付</span>
              </Descriptions.Item>
            </Descriptions>
          </div>
        )}

        <div style={{ 
          marginTop: 24, 
          padding: 16, 
          background: '#f6ffed', 
          border: '1px solid #b7eb8f', 
          borderRadius: 6,
          color: '#52c41a'
        }}>
          <p style={{ margin: 0, fontWeight: 'bold' }}>温馨提示：</p>
          <p style={{ margin: '8px 0 0 0' }}>
            • 请保持手机畅通，我们会在取车前联系您确认详细信息<br/>
            • 取车时请携带有效身份证件和驾驶证<br/>
            • 如有任何问题，请联系客服：400-123-4567
          </p>
        </div>
      </ContentCard>
    </Container>
  );
};

export default PaymentSuccess;
