import React, { useState, useEffect } from 'react';
import {
  Card,
  List,
  Button,
  Tag,
  Space,
  Modal,
  message,
  Descriptions,
  Image,
  Typography,
  Divider,
  Empty
} from 'antd';
import {
  CarOutlined,
  CalendarOutlined,
  EnvironmentOutlined,
  DollarOutlined,
  ClockCircleOutlined,
  ExclamationCircleOutlined
} from '@ant-design/icons';
import styled from 'styled-components';
import reservationService from '../../services/reservationService';
import orderService from '../../services/orderService';
import paymentService from '../../services/paymentService';

const { Text, Title } = Typography;
const { confirm } = Modal;

const StyledCard = styled(Card)`
  margin-bottom: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  
  .ant-card-body {
    padding: 20px;
  }
`;

const VehicleImage = styled(Image)`
  border-radius: 8px;
`;

const StatusTag = styled(Tag)`
  font-weight: 500;
`;

const ReservationList = ({ activeTab = 'all' }) => {
  const [reservations, setReservations] = useState([]);
  const [loading, setLoading] = useState(false);
  const [detailVisible, setDetailVisible] = useState(false);
  const [selectedReservation, setSelectedReservation] = useState(null);

  // 获取预订列表
  const fetchReservations = async () => {
    setLoading(true);
    try {
      const response = await reservationService.getUserReservations();
      if (response && response.code === 200) {
        setReservations(response.data || []);
      }
    } catch (error) {
      message.error('获取预订列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchReservations();
  }, []);

  // 根据状态过滤预订
  const getFilteredReservations = () => {
    if (activeTab === 'all') return reservations;
    return reservations.filter(reservation => {
      switch (activeTab) {
        case 'processing':
          return reservation.status === 'processing';
        case 'pending_payment':
          return reservation.status === 'pending_payment';
        case 'confirmed':
          return reservation.status === 'confirmed';
        case 'in_use':
          return reservation.status === 'in_use';
        case 'completed':
          return reservation.status === 'completed';
        case 'cancelled':
          return reservation.status === 'cancelled';
        default:
          return true;
      }
    });
  };

  // 获取状态标签
  const getStatusTag = (status) => {
    const statusMap = {
      processing: { color: 'processing', text: '处理中' },
      pending_payment: { color: 'warning', text: '等待付款' },
      confirmed: { color: 'success', text: '预订成功' },
      in_use: { color: 'blue', text: '租赁中' },
      completed: { color: 'default', text: '已完成' },
      cancelled: { color: 'error', text: '已取消' }
    };
    const config = statusMap[status] || { color: 'default', text: '未知' };
    return <StatusTag color={config.color}>{config.text}</StatusTag>;
  };

  // 处理支付
  const handlePayment = async (reservation) => {
    try {
      // 先从预订创建订单
      const orderResponse = await orderService.createOrderFromReservation(reservation.id);
      if (orderResponse && orderResponse.code === 200) {
        const orderId = orderResponse.data.id;
        
        // 获取支付链接
        const paymentResponse = await paymentService.getPaymentUrl(orderId, 'alipay');
        if (paymentResponse && paymentResponse.code === 200) {
          // 打开支付链接
          window.open(paymentResponse.data.payment_url, '_blank');
          message.success('支付链接已打开，请完成支付');
          
          // 刷新预订列表
          setTimeout(() => {
            fetchReservations();
          }, 2000);
        }
      }
    } catch (error) {
      message.error('创建支付失败，请稍后重试');
    }
  };

  // 取消预订
  const handleCancel = (reservation) => {
    confirm({
      title: '确认取消预订',
      icon: <ExclamationCircleOutlined />,
      content: '取消后无法恢复，确定要取消这个预订吗？',
      okText: '确认取消',
      okType: 'danger',
      cancelText: '我再想想',
      onOk: async () => {
        try {
          await reservationService.cancelReservation(reservation.id);
          message.success('预订已取消');
          fetchReservations();
        } catch (error) {
          message.error('取消预订失败');
        }
      }
    });
  };

  // 查看详情
  const handleViewDetail = (reservation) => {
    setSelectedReservation(reservation);
    setDetailVisible(true);
  };

  const filteredReservations = getFilteredReservations();

  return (
    <>
      {filteredReservations.length > 0 ? (
        <List
          loading={loading}
          dataSource={filteredReservations}
          renderItem={(reservation) => (
            <StyledCard key={reservation.id}>
              <div style={{ display: 'flex', gap: '20px' }}>
                <VehicleImage
                  width={120}
                  height={80}
                  src={reservation.vehicle?.images?.[0] || '/placeholder-car.jpg'}
                  alt={reservation.vehicle?.name}
                  style={{ objectFit: 'cover' }}
                />
                <div style={{ flex: 1 }}>
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: 12 }}>
                    <div>
                      <Title level={5} style={{ margin: 0, marginBottom: 4 }}>
                        <CarOutlined style={{ marginRight: 8, color: '#1890ff' }} />
                        {reservation.vehicle?.brand} {reservation.vehicle?.name}
                      </Title>
                      <Text type="secondary">预订单号: {reservation.id}</Text>
                    </div>
                    {getStatusTag(reservation.status)}
                  </div>
                  
                  <Space direction="vertical" size={4} style={{ width: '100%' }}>
                    <div>
                      <CalendarOutlined style={{ marginRight: 8, color: '#52c41a' }} />
                      <Text>{reservation.start_date} 至 {reservation.end_date}</Text>
                    </div>
                    <div>
                      <EnvironmentOutlined style={{ marginRight: 8, color: '#fa8c16' }} />
                      <Text>{reservation.pickup_location}</Text>
                    </div>
                    <div>
                      <DollarOutlined style={{ marginRight: 8, color: '#f5222d' }} />
                      <Text strong>总价: ¥{reservation.total_amount}</Text>
                    </div>
                  </Space>
                  
                  <Divider style={{ margin: '12px 0' }} />
                  
                  <Space>
                    <Button size="small" onClick={() => handleViewDetail(reservation)}>
                      查看详情
                    </Button>
                    {reservation.status === 'pending_payment' && (
                      <Button type="primary" size="small" onClick={() => handlePayment(reservation)}>
                        立即支付
                      </Button>
                    )}
                    {['processing', 'pending_payment'].includes(reservation.status) && (
                      <Button danger size="small" onClick={() => handleCancel(reservation)}>
                        取消预订
                      </Button>
                    )}
                  </Space>
                </div>
              </div>
            </StyledCard>
          )}
        />
      ) : (
        <Empty
          image={Empty.PRESENTED_IMAGE_SIMPLE}
          description={
            <div>
              <Text>还没有预订记录</Text>
              <br />
              <Button type="link" onClick={() => window.location.href = '/vehicles'}>去预订豪车</Button>
            </div>
          }
        />
      )}

      {/* 预订详情弹窗 */}
      <Modal
        title="预订详情"
        open={detailVisible}
        onCancel={() => setDetailVisible(false)}
        footer={null}
        width={600}
      >
        {selectedReservation && (
          <Descriptions column={1} bordered>
            <Descriptions.Item label="预订单号">{selectedReservation.id}</Descriptions.Item>
            <Descriptions.Item label="车辆信息">
              {selectedReservation.vehicle?.brand} {selectedReservation.vehicle?.name}
            </Descriptions.Item>
            <Descriptions.Item label="租赁时间">
              {selectedReservation.start_date} 至 {selectedReservation.end_date}
            </Descriptions.Item>
            <Descriptions.Item label="取车地点">{selectedReservation.pickup_location}</Descriptions.Item>
            <Descriptions.Item label="还车地点">{selectedReservation.return_location}</Descriptions.Item>
            <Descriptions.Item label="总金额">¥{selectedReservation.total_amount}</Descriptions.Item>
            <Descriptions.Item label="预订状态">{getStatusTag(selectedReservation.status)}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{selectedReservation.created_at}</Descriptions.Item>
          </Descriptions>
        )}
      </Modal>
    </>
  );
};

export default ReservationList;
