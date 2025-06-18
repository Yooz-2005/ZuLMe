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
  Empty,
  Tooltip
} from 'antd';
import {
  CarOutlined,
  CalendarOutlined,
  EnvironmentOutlined,
  DollarOutlined,
  ClockCircleOutlined,
  ExclamationCircleOutlined,
  CheckCircleOutlined,
  SyncOutlined,
  CloseCircleOutlined
} from '@ant-design/icons';
import styled from 'styled-components';
import orderService from '../../services/orderService';

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

const OrderList = ({ activeTab = 'all' }) => {
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(false);
  const [detailVisible, setDetailVisible] = useState(false);
  const [selectedOrder, setSelectedOrder] = useState(null);

  // 获取订单列表
  const fetchOrders = async () => {
    setLoading(true);
    try {
      const response = await orderService.getUserOrders();
      if (response && response.code === 200) {
        setOrders(response.data?.orders || []);
      }
    } catch (error) {
      message.error('获取订单列表失败');
      console.error('获取订单列表错误:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  // 根据状态过滤订单
  const getFilteredOrders = () => {
    if (activeTab === 'all') return orders;
    return orders.filter(order => {
      switch (activeTab) {
        case 'pending_payment':
          return order.status === 1; // 待支付
        case 'paid':
          return order.status === 2; // 已支付
        case 'in_use':
          return order.status === 3; // 使用中
        case 'completed':
          return order.status === 4; // 已完成
        case 'cancelled':
          return order.status === 5; // 已取消
        default:
          return true;
      }
    });
  };

  // 获取状态标签
  const getStatusTag = (status) => {
    const statusMap = {
      1: { color: 'warning', text: '待支付', icon: <ClockCircleOutlined /> },
      2: { color: 'success', text: '已支付', icon: <CheckCircleOutlined /> },
      3: { color: 'processing', text: '使用中', icon: <SyncOutlined spin /> },
      4: { color: 'default', text: '已完成', icon: <CheckCircleOutlined /> },
      5: { color: 'error', text: '已取消', icon: <CloseCircleOutlined /> }
    };
    const config = statusMap[status] || { color: 'default', text: '未知', icon: null };
    return (
      <StatusTag color={config.color}>
        {config.icon && <span style={{ marginRight: 4 }}>{config.icon}</span>}
        {config.text}
      </StatusTag>
    );
  };

  // 获取车辆第一张图片
  const getVehicleImage = (vehicle) => {
    if (!vehicle) return '/placeholder-car.jpg';

    // 如果images是字符串，分割获取第一张图片
    if (typeof vehicle.images === 'string' && vehicle.images) {
      const imageArray = vehicle.images.split(',').map(img => img.trim()).filter(img => img);
      if (imageArray.length > 0) {
        return imageArray[0];
      }
    }

    // 如果images是数组，获取第一张图片
    if (Array.isArray(vehicle.images) && vehicle.images.length > 0) {
      return vehicle.images[0];
    }

    // 默认图片
    return '/placeholder-car.jpg';
  };

  // 取消订单
  const handleCancel = (order) => {
    confirm({
      title: '确认取消订单',
      icon: <ExclamationCircleOutlined />,
      content: '取消后无法恢复，确定要取消这个订单吗？',
      okText: '确认取消',
      okType: 'danger',
      cancelText: '我再想想',
      onOk: async () => {
        try {
          await orderService.cancelOrder(order.id);
          message.success('订单已取消');
          fetchOrders();
        } catch (error) {
          message.error('取消订单失败');
        }
      }
    });
  };

  // 查看详情
  const handleViewDetail = (order) => {
    setSelectedOrder(order);
    setDetailVisible(true);
  };

  // 继续支付
  const handleContinuePayment = (order) => {
    const paymentUrl = `http://localhost:3000/mock-payment.html?order_sn=${order.order_sn}&amount=${order.total_amount}&subject=${encodeURIComponent('租车订单-豪华车辆')}&app_id=2021000122671234`;
    window.open(paymentUrl, '_blank');
    message.success('支付链接已打开，请完成支付');
  };

  const filteredOrders = getFilteredOrders();

  return (
    <>
      {filteredOrders.length > 0 ? (
        <List
          loading={loading}
          dataSource={filteredOrders}
          renderItem={(order) => (
            <StyledCard key={order.id}>
              <div style={{ display: 'flex', gap: '20px' }}>
                <VehicleImage
                  width={120}
                  height={80}
                  src={getVehicleImage(order.vehicle)}
                  alt={order.vehicle?.style || order.vehicle?.name}
                  style={{ objectFit: 'cover' }}
                />
                <div style={{ flex: 1 }}>
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: 12 }}>
                    <div>
                      <Title level={5} style={{ margin: 0, marginBottom: 4 }}>
                        <CarOutlined style={{ marginRight: 8, color: '#1890ff' }} />
                        {order.vehicle?.brand} {order.vehicle?.style || order.vehicle?.name}
                      </Title>
                      <Text type="secondary">订单号: {order.order_sn}</Text>
                    </div>
                    {getStatusTag(order.status)}
                  </div>
                  
                  <Space direction="vertical" size={4} style={{ width: '100%' }}>
                    <div>
                      <CalendarOutlined style={{ marginRight: 8, color: '#52c41a' }} />
                      <Text>{order.start_date} 至 {order.end_date}</Text>
                    </div>
                    <div>
                      <EnvironmentOutlined style={{ marginRight: 8, color: '#fa8c16' }} />
                      <Text>取车：{order.pickup_location}</Text>
                      {order.return_location && (
                        <Text style={{ marginLeft: 16 }}>还车：{order.return_location}</Text>
                      )}
                    </div>
                    <div>
                      <DollarOutlined style={{ marginRight: 8, color: '#f5222d' }} />
                      <Text strong>总价: ¥{order.total_amount}</Text>
                    </div>
                  </Space>
                  
                  <Divider style={{ margin: '12px 0' }} />
                  
                  <Space>
                    <Button size="small" onClick={() => handleViewDetail(order)}>
                      查看详情
                    </Button>
                    {order.status === 1 && (
                      <Button type="primary" size="small" onClick={() => handleContinuePayment(order)}>
                        继续支付
                      </Button>
                    )}
                    {order.status === 1 && (
                      <Button danger size="small" onClick={() => handleCancel(order)}>
                        取消订单
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
              <Text>还没有订单记录</Text>
              <br />
              <Button type="link" onClick={() => window.location.href = '/vehicles'}>去预订豪车</Button>
            </div>
          }
        />
      )}

      {/* 订单详情弹窗 */}
      <Modal
        title="订单详情"
        open={detailVisible}
        onCancel={() => setDetailVisible(false)}
        footer={null}
        width={600}
      >
        {selectedOrder && (
          <Descriptions column={1} bordered>
            <Descriptions.Item label="订单号">{selectedOrder.order_sn}</Descriptions.Item>
            <Descriptions.Item label="车辆信息">
              {selectedOrder.vehicle?.brand} {selectedOrder.vehicle?.style || selectedOrder.vehicle?.name}
            </Descriptions.Item>
            <Descriptions.Item label="租赁时间">
              {selectedOrder.start_date} 至 {selectedOrder.end_date}
            </Descriptions.Item>
            <Descriptions.Item label="取车地点">{selectedOrder.pickup_location}</Descriptions.Item>
            {selectedOrder.return_location && (
              <Descriptions.Item label="还车地点">{selectedOrder.return_location}</Descriptions.Item>
            )}
            <Descriptions.Item label="总金额">¥{selectedOrder.total_amount}</Descriptions.Item>
            <Descriptions.Item label="订单状态">{getStatusTag(selectedOrder.status)}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{selectedOrder.created_at}</Descriptions.Item>
            {selectedOrder.paid_at && (
              <Descriptions.Item label="支付时间">{selectedOrder.paid_at}</Descriptions.Item>
            )}
            {selectedOrder.notes && (
              <Descriptions.Item label="备注">{selectedOrder.notes}</Descriptions.Item>
            )}
          </Descriptions>
        )}
      </Modal>
    </>
  );
};

export default OrderList;
