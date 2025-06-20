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
import { parseImages, getDefaultImageByBrand } from '../../utils/imageUtils';

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
          return order.payment_status === 1; // 待支付
        case 'paid':
          return order.payment_status === 2; // 已支付
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

  // 获取状态标签（基于支付状态）
  const getStatusTag = (order) => {
    const paymentStatus = order.payment_status;
    const orderStatus = order.status;

    // 优先显示支付状态
    if (paymentStatus === 1) {
      return (
        <StatusTag color="warning">
          <ClockCircleOutlined style={{ marginRight: 4 }} />
          待支付
        </StatusTag>
      );
    } else if (paymentStatus === 2) {
      return (
        <StatusTag color="success">
          <CheckCircleOutlined style={{ marginRight: 4 }} />
          已支付
        </StatusTag>
      );
    } else if (paymentStatus === 3) {
      return (
        <StatusTag color="error">
          <CloseCircleOutlined style={{ marginRight: 4 }} />
          已取消
        </StatusTag>
      );
    }

    // 其他状态
    const statusMap = {
      3: { color: 'processing', text: '使用中', icon: <SyncOutlined spin /> },
      4: { color: 'default', text: '已完成', icon: <CheckCircleOutlined /> },
      5: { color: 'error', text: '已取消', icon: <CloseCircleOutlined /> }
    };
    const config = statusMap[orderStatus] || { color: 'default', text: '未知', icon: null };
    return (
      <StatusTag color={config.color}>
        {config.icon && <span style={{ marginRight: 4 }}>{config.icon}</span>}
        {config.text}
      </StatusTag>
    );
  };

  // 从Notes中解析车辆信息
  const parseVehicleInfo = (notes) => {
    if (!notes) return { brand: '未知车辆', style: '', images: '' };

    const vehicleMatch = notes.match(/车辆:\s*([^;]+)/);
    if (vehicleMatch) {
      const vehicleInfo = vehicleMatch[1].trim();
      const parts = vehicleInfo.split(' ');
      return {
        brand: parts[0] || '未知品牌',
        style: parts.slice(1).join(' ') || '',
        images: ''
      };
    }
    return { brand: '未知车辆', style: '', images: '' };
  };

  // 从Notes中解析地点信息
  const parseLocationInfo = (notes) => {
    if (!notes) return { pickup: '', return: '' };

    const pickupMatch = notes.match(/取车:\s*([^;]+)/);
    const returnMatch = notes.match(/还车:\s*([^;]+)/);

    return {
      pickup: pickupMatch ? pickupMatch[1].trim() : '',
      return: returnMatch ? returnMatch[1].trim() : ''
    };
  };

  // 获取车辆第一张图片
  const getVehicleImage = (order) => {
    console.log('Order notes:', order.notes); // 调试日志

    // 获取车辆品牌信息
    const vehicleInfo = parseVehicleInfo(order.notes);
    const brand = vehicleInfo.brand;

    if (order.notes) {
      // 从notes中解析图片信息
      const imageMatch = order.notes.match(/图片:\s*([^;]+)/);
      console.log('Image match:', imageMatch); // 调试日志
      if (imageMatch) {
        const imagesString = imageMatch[1].trim();
        console.log('Images string found:', imagesString); // 调试日志

        // 使用工具函数解析图片
        const images = parseImages(imagesString, brand);
        if (images && images.length > 0) {
          console.log('Using image:', images[0]); // 调试日志
          return images[0];
        }
      }
    }

    // 如果没有图片信息，根据品牌使用默认图片
    const defaultImage = getDefaultImageByBrand(brand);
    console.log('Using default image for brand', brand, ':', defaultImage); // 调试日志
    return defaultImage;
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
    const paymentUrl = order.payment_url;

    if (!paymentUrl) {
      message.error('支付链接不存在，请重新创建订单');
      return;
    }

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
                  src={getVehicleImage(order)}
                  alt="车辆图片"
                  style={{ objectFit: 'cover' }}
                />
                <div style={{ flex: 1 }}>
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: 12 }}>
                    <div>
                      <Title level={5} style={{ margin: 0, marginBottom: 4 }}>
                        <CarOutlined style={{ marginRight: 8, color: '#1890ff' }} />
                        {(() => {
                          const vehicleInfo = parseVehicleInfo(order.notes);
                          return `${vehicleInfo.brand} ${vehicleInfo.style}`;
                        })()}
                      </Title>
                      <Text type="secondary">订单号: {order.order_sn}</Text>
                    </div>
                    {getStatusTag(order)}
                  </div>
                  
                  <Space direction="vertical" size={4} style={{ width: '100%' }}>
                    <div>
                      <CalendarOutlined style={{ marginRight: 8, color: '#52c41a' }} />
                      <Text>{order.pickup_time} 至 {order.return_time}</Text>
                    </div>
                    {(() => {
                      const locationInfo = parseLocationInfo(order.notes);
                      return (
                        <div>
                          <EnvironmentOutlined style={{ marginRight: 8, color: '#fa8c16' }} />
                          {locationInfo.pickup && <Text>取车：{locationInfo.pickup}</Text>}
                          {locationInfo.return && (
                            <Text style={{ marginLeft: 16 }}>还车：{locationInfo.return}</Text>
                          )}
                        </div>
                      );
                    })()}
                    <div>
                      <DollarOutlined style={{ marginRight: 8, color: '#f5222d' }} />
                      <Text strong>总价: ¥{order.total_amount}</Text>
                    </div>
                    <div>
                      <Text type="secondary">订单创建时间: {new Date(order.created_at).toLocaleString()}</Text>
                    </div>
                  </Space>
                  
                  <Divider style={{ margin: '12px 0' }} />
                  
                  <Space>
                    <Button size="small" onClick={() => handleViewDetail(order)}>
                      查看详情
                    </Button>
                    {order.payment_status === 1 && order.payment_url && (
                      <Button type="primary" size="small" onClick={() => handleContinuePayment(order)}>
                        继续支付
                      </Button>
                    )}
                    {order.payment_status === 1 && (
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
        {selectedOrder && (() => {
          const vehicleInfo = parseVehicleInfo(selectedOrder.notes);
          const locationInfo = parseLocationInfo(selectedOrder.notes);
          return (
            <Descriptions column={1} bordered>
              <Descriptions.Item label="订单号">{selectedOrder.order_sn}</Descriptions.Item>
              <Descriptions.Item label="车辆信息">
                {vehicleInfo.brand} {vehicleInfo.style}
              </Descriptions.Item>
              <Descriptions.Item label="租赁时间">
                {selectedOrder.pickup_time} 至 {selectedOrder.return_time}
              </Descriptions.Item>
              <Descriptions.Item label="租赁天数">{selectedOrder.rental_days}天</Descriptions.Item>
              {locationInfo.pickup && (
                <Descriptions.Item label="取车地点">{locationInfo.pickup}</Descriptions.Item>
              )}
              {locationInfo.return && (
                <Descriptions.Item label="还车地点">{locationInfo.return}</Descriptions.Item>
              )}
              <Descriptions.Item label="总金额">¥{selectedOrder.total_amount}</Descriptions.Item>
              <Descriptions.Item label="订单状态">{getStatusTag(selectedOrder)}</Descriptions.Item>
              <Descriptions.Item label="创建时间">{new Date(selectedOrder.created_at).toLocaleString()}</Descriptions.Item>
              {selectedOrder.notes && selectedOrder.notes.includes('备注:') && (
                <Descriptions.Item label="备注">
                  {selectedOrder.notes.split('备注:')[1]?.trim()}
                </Descriptions.Item>
              )}
            </Descriptions>
          );
        })()}
      </Modal>
    </>
  );
};

export default OrderList;
