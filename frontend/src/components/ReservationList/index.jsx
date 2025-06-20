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
  Select,
  Row,
  Col
} from 'antd';
import {
  CarOutlined,
  CalendarOutlined,
  EnvironmentOutlined,
  DollarOutlined,
  ClockCircleOutlined,
  ExclamationCircleOutlined,
  CreditCardOutlined,
  CheckCircleOutlined
} from '@ant-design/icons';
import styled from 'styled-components';
import reservationService from '../../services/reservationService';
import orderService from '../../services/orderService';
import vehicleService from '../../services/vehicleService';

const { Text, Title } = Typography;
const { confirm } = Modal;
const { Option } = Select;

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
  const [paymentModalVisible, setPaymentModalVisible] = useState(false);
  const [selectedReturnLocation, setSelectedReturnLocation] = useState(null);
  const [paymentLoading, setPaymentLoading] = useState(false);
  const [returnLocations, setReturnLocations] = useState([]);
  const [locationsLoading, setLocationsLoading] = useState(false);

  // 获取预订列表
  const fetchReservations = async () => {
    setLoading(true);
    try {
      const response = await reservationService.getUserReservations();
      if (response && response.code === 200) {
        // 修复：从response.data.reservations获取预订列表数组
        setReservations(response.data?.reservations || []);
      }
    } catch (error) {
      message.error('获取预订列表失败');
      console.error('获取预订列表错误:', error);
    } finally {
      setLoading(false);
    }
  };

  // 获取网点列表（还车地点）
  const fetchReturnLocations = async () => {
    setLocationsLoading(true);
    try {
      const response = await vehicleService.getLocationList();
      if (response && response.code === 200) {
        // 将商户信息转换为还车地点格式
        const locations = response.data?.merchants?.map(merchant => ({
          id: merchant.id,
          name: merchant.name,
          address: merchant.location,
          business_time: merchant.business_time,
          longitude: merchant.longitude,
          latitude: merchant.latitude
        })) || [];
        setReturnLocations(locations);
      }
    } catch (error) {
      console.error('获取网点列表失败:', error);
      // 如果获取失败，使用默认的模拟数据
      setReturnLocations([
        { id: 1, name: '北京首都国际机场T3航站楼', address: '北京市朝阳区首都机场' },
        { id: 2, name: '北京大兴国际机场', address: '北京市大兴区大兴机场' },
        { id: 3, name: '北京西站', address: '北京市西城区莲花池东路' },
        { id: 4, name: '北京南站', address: '北京市丰台区永外大街' }
      ]);
    } finally {
      setLocationsLoading(false);
    }
  };

  useEffect(() => {
    fetchReservations();
    fetchReturnLocations();
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

  // 处理支付 - 打开支付选择弹窗
  const handlePayment = (reservation) => {
    console.log('准备支付的预订信息:', reservation);
    console.log('预订ID:', reservation.id, '类型:', typeof reservation.id);
    setSelectedReservation(reservation);
    setPaymentModalVisible(true);
    setSelectedReturnLocation(null); // 重置还车地点选择
  };

  // 确认支付
  const confirmPayment = async () => {
    if (!selectedReturnLocation) {
      message.error('请选择还车地点');
      return;
    }

    const reservation = selectedReservation;
    console.log('开始处理支付，预订信息:', reservation);
    console.log('选择的还车地点:', selectedReturnLocation);

    // 检查预订信息是否有效
    if (!reservation || !reservation.id) {
      console.error('预订信息无效:', reservation);
      message.error('预订信息无效，请刷新页面重试');
      return;
    }

    // 验证预订ID并提取数字部分（订单服务期望数字格式）
    let reservationId = reservation.id;
    let numericReservationId;

    // 检查预订ID是否为空
    if (!reservationId) {
      console.error('预订ID为空:', reservation.id);
      message.error('预订ID为空，请刷新页面重试');
      return;
    }

    // 如果是字符串格式的ID（如 "RES123"），提取数字部分
    if (typeof reservationId === 'string') {
      const match = reservationId.match(/RES(\d+)/);
      if (match) {
        numericReservationId = parseInt(match[1], 10);
        console.log('从字符串ID提取数字:', reservationId, '->', numericReservationId);
      } else {
        // 尝试直接解析为数字
        numericReservationId = parseInt(reservationId, 10);
        if (isNaN(numericReservationId)) {
          console.error('无法从预订ID提取数字:', reservationId);
          message.error('预订ID格式无效，请刷新页面重试');
          return;
        }
      }
    } else if (typeof reservationId === 'number') {
      numericReservationId = reservationId;
      console.log('使用数字格式的预订ID:', numericReservationId);
    } else {
      console.error('预订ID类型无效:', reservation.id, typeof reservation.id);
      message.error('预订ID类型无效，请刷新页面重试');
      return;
    }

    if (isNaN(numericReservationId) || numericReservationId <= 0) {
      console.error('预订ID数字无效:', numericReservationId);
      message.error('预订ID格式无效，请刷新页面重试');
      return;
    }

    console.log('预订ID验证通过:', numericReservationId);

    // 检查用户是否已登录
    const token = localStorage.getItem('token');
    console.log('用户token:', token);
    if (!token) {
      message.error('请先登录后再进行支付');
      return;
    }

    setPaymentLoading(true);
    try {
      // 创建订单时包含还车地点信息（取车地点由后端自动从车辆信息获取）
      const orderData = {
        reservation_id: numericReservationId, // 使用提取的数字ID
        return_location_id: selectedReturnLocation,
        payment_method: "alipay", // 使用字符串格式
        notes: `还车地点：${returnLocations.find(loc => loc.id === selectedReturnLocation)?.name}`,
        expected_total_amount: parseFloat(reservation.total_amount) || 0
      };

      console.log('调用创建订单接口，订单数据:', orderData);
      const orderResponse = await orderService.createOrderFromReservation(orderData);
      console.log('创建订单响应:', orderResponse);

      if (orderResponse && orderResponse.code === 200) {
        console.log('订单创建成功，订单数据:', orderResponse.data);

        // 关闭支付弹窗
        setPaymentModalVisible(false);

        // 使用创建订单时返回的支付链接
        const paymentUrl = orderResponse.data.payment_url;

        console.log('支付链接:', paymentUrl);

        if (!paymentUrl) {
          message.error('支付链接生成失败，请重试');
          return;
        }

        // 打开支付链接
        window.open(paymentUrl, '_blank');
        message.success('支付链接已打开，请完成支付');

        // 刷新预订列表
        setTimeout(() => {
          fetchReservations();
        }, 2000);
      } else {
        console.error('创建订单失败:', orderResponse);
        const errorMsg = orderResponse?.message || '创建订单失败';
        message.error(`创建订单失败: ${errorMsg}`);
      }
    } catch (error) {
      console.error('支付处理错误:', error);
      console.error('错误详情:', {
        message: error.message,
        response: error.response?.data,
        status: error.response?.status
      });

      let errorMessage = '创建支付失败，请稍后重试';
      if (error.response?.data?.message) {
        errorMessage = error.response.data.message;
      } else if (error.message) {
        errorMessage = error.message;
      }

      message.error(errorMessage);
    } finally {
      setPaymentLoading(false);
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
                  src={getVehicleImage(reservation.vehicle)}
                  alt={reservation.vehicle?.style || reservation.vehicle?.name}
                  style={{ objectFit: 'cover' }}
                />
                <div style={{ flex: 1 }}>
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: 12 }}>
                    <div>
                      <Title level={5} style={{ margin: 0, marginBottom: 4 }}>
                        <CarOutlined style={{ marginRight: 8, color: '#1890ff' }} />
                        {reservation.vehicle?.brand} {reservation.vehicle?.style || reservation.vehicle?.name}
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
              {selectedReservation.vehicle?.brand} {selectedReservation.vehicle?.style || selectedReservation.vehicle?.name}
            </Descriptions.Item>
            <Descriptions.Item label="租赁时间">
              {selectedReservation.start_date} 至 {selectedReservation.end_date}
            </Descriptions.Item>
            <Descriptions.Item label="取车地点">{selectedReservation.pickup_location}</Descriptions.Item>
            {selectedReservation.return_location && (
              <Descriptions.Item label="还车地点">{selectedReservation.return_location}</Descriptions.Item>
            )}
            <Descriptions.Item label="总金额">¥{selectedReservation.total_amount}</Descriptions.Item>
            <Descriptions.Item label="预订状态">{getStatusTag(selectedReservation.status)}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{selectedReservation.created_at}</Descriptions.Item>
            {selectedReservation.status === 'pending_payment' && (
              <Descriptions.Item label="提示">
                <Text type="warning">请完成支付以确认预订，支付时可选择还车地点</Text>
              </Descriptions.Item>
            )}
          </Descriptions>
        )}
      </Modal>

      {/* 支付选择弹窗 */}
      <Modal
        title="选择还车地点并支付"
        open={paymentModalVisible}
        onCancel={() => setPaymentModalVisible(false)}
        footer={[
          <Button key="cancel" onClick={() => setPaymentModalVisible(false)}>
            取消
          </Button>,
          <Button
            key="confirm"
            type="primary"
            loading={paymentLoading}
            onClick={confirmPayment}
            disabled={!selectedReturnLocation}
            icon={<CreditCardOutlined />}
          >
            确认支付
          </Button>
        ]}
        width={600}
      >
        {selectedReservation && (
          <div>
            <Card size="small" style={{ marginBottom: 20, background: '#f8f9fa' }}>
              <Row gutter={16}>
                <Col span={6}>
                  <Image
                    width={80}
                    height={60}
                    src={getVehicleImage(selectedReservation.vehicle)}
                    alt={selectedReservation.vehicle?.style}
                    style={{ objectFit: 'cover', borderRadius: 6 }}
                  />
                </Col>
                <Col span={18}>
                  <div>
                    <Text strong>
                      {selectedReservation.vehicle?.brand} {selectedReservation.vehicle?.style}
                    </Text>
                    <br />
                    <Text type="secondary">
                      {selectedReservation.start_date} 至 {selectedReservation.end_date}
                    </Text>
                    <br />
                    <Text type="secondary">
                      取车地点：{selectedReservation.pickup_location}
                    </Text>
                    <br />
                    <Text strong style={{ color: '#f5222d' }}>
                      总金额：¥{selectedReservation.total_amount}
                    </Text>
                  </div>
                </Col>
              </Row>
            </Card>

            <Divider orientation="left">选择还车地点</Divider>
            <Select
              style={{ width: '100%' }}
              placeholder={locationsLoading ? "正在加载网点..." : "请选择还车地点"}
              value={selectedReturnLocation}
              onChange={setSelectedReturnLocation}
              size="large"
              loading={locationsLoading}
              disabled={locationsLoading}
            >
              {returnLocations.map(location => (
                <Option key={location.id} value={location.id}>
                  <div>
                    <div style={{ fontWeight: 500 }}>{location.name}</div>
                    <div style={{ fontSize: '12px', color: '#999' }}>
                      {location.address}
                      {location.business_time && (
                        <span style={{ marginLeft: 8, color: '#52c41a' }}>
                          营业时间: {location.business_time}
                        </span>
                      )}
                    </div>
                  </div>
                </Option>
              ))}
            </Select>

            <div style={{ marginTop: 16, padding: 12, background: '#fff7e6', borderRadius: 6, border: '1px solid #ffd591' }}>
              <Text type="warning">
                <ExclamationCircleOutlined style={{ marginRight: 8 }} />
                请确认还车地点，支付成功后将无法更改
              </Text>
            </div>
          </div>
        )}
      </Modal>
    </>
  );
};

export default ReservationList;
