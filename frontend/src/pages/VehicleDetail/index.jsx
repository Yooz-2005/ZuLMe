import React, { useState, useEffect } from 'react';
import { 
  Layout, 
  Row, 
  Col, 
  Typography, 
  Spin, 
  Alert, 
  Button,
  Card,
  Tag,
  Descriptions,
  Image,
  Space,
  Breadcrumb,
  Divider
} from 'antd';
import { 
  HomeOutlined, 
  CarOutlined, 
  UserOutlined, 
  SettingOutlined,
  SafetyOutlined,
  EnvironmentOutlined,
  CalendarOutlined
} from '@ant-design/icons';
import { useParams, useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import vehicleService from '../../services/vehicleService';
import { VEHICLE_STATUS_LABELS, VEHICLE_TYPE_LABELS } from '../../utils/constants';

const { Header, Content } = Layout;
const { Title, Text, Paragraph } = Typography;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
`;

const StyledHeader = styled(Header)`
  background: #000;
  padding: 0 50px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  position: fixed;
  width: 100%;
  z-index: 1;
  top: 0;
  left: 0;
`;

const ContentWrapper = styled.div`
  margin-top: 64px;
  padding: 24px 50px;
`;

const ImageSection = styled.div`
  .ant-image {
    width: 100%;
    border-radius: 8px;
    overflow: hidden;
  }
`;

const PriceCard = styled(Card)`
  position: sticky;
  top: 88px;
  
  .price-text {
    font-size: 28px;
    font-weight: bold;
    color: #ff4d4f;
    margin: 0;
  }
  
  .price-unit {
    font-size: 16px;
    color: #666;
    margin-left: 8px;
  }
`;

const FeatureTag = styled(Tag)`
  margin: 4px;
  padding: 4px 12px;
  border-radius: 16px;
`;

const VehicleDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [vehicle, setVehicle] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchVehicleDetail = async () => {
      if (!id) {
        setError('车辆ID不存在');
        setLoading(false);
        return;
      }

      setLoading(true);
      setError(null);
      
      try {
        const response = await vehicleService.getVehicleDetail(id);
        if (response && response.data) {
          setVehicle(response.data);
        } else {
          setError('车辆信息不存在');
        }
      } catch (err) {
        setError('获取车辆详情失败，请稍后重试');
        console.error('获取车辆详情失败:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchVehicleDetail();
  }, [id]);

  const getStatusColor = (status) => {
    switch (status) {
      case 'available':
        return 'green';
      case 'rented':
        return 'red';
      case 'maintenance':
        return 'orange';
      default:
        return 'default';
    }
  };

  if (loading) {
    return (
      <StyledLayout>
        <StyledHeader>
          <Row justify="space-between" align="middle">
            <Col>
              <Title level={3} style={{ margin: 0, cursor: 'pointer', color: '#fff' }} onClick={() => navigate('/')}>
                ZuLMe
              </Title>
            </Col>
          </Row>
        </StyledHeader>
        <Content>
          <ContentWrapper>
            <div style={{ textAlign: 'center', padding: '100px 0' }}>
              <Spin size="large" />
            </div>
          </ContentWrapper>
        </Content>
      </StyledLayout>
    );
  }

  if (error || !vehicle) {
    return (
      <StyledLayout>
        <StyledHeader>
          <Row justify="space-between" align="middle">
            <Col>
              <Title level={3} style={{ margin: 0, cursor: 'pointer', color: '#fff' }} onClick={() => navigate('/')}>
                ZuLMe
              </Title>
            </Col>
          </Row>
        </StyledHeader>
        <Content>
          <ContentWrapper>
            <Alert
              message="错误"
              description={error}
              type="error"
              showIcon
              action={
                <Space>
                  <Button onClick={() => window.location.reload()}>重试</Button>
                  <Button onClick={() => navigate('/vehicles')}>返回列表</Button>
                </Space>
              }
            />
          </ContentWrapper>
        </Content>
      </StyledLayout>
    );
  }

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={3} style={{ margin: 0, cursor: 'pointer', color: '#fff' }} onClick={() => navigate('/')}>
              ZuLMe
            </Title>
          </Col>
          <Col>
            <Space>
              <Button onClick={() => navigate('/vehicles')}>返回列表</Button>
              <Button onClick={() => navigate('/')}>返回首页</Button>
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          <Breadcrumb style={{ marginBottom: 16 }}>
            <Breadcrumb.Item href="/">
              <HomeOutlined />
            </Breadcrumb.Item>
            <Breadcrumb.Item href="/vehicles">
              <CarOutlined />
              <span>车辆列表</span>
            </Breadcrumb.Item>
            <Breadcrumb.Item>
              {vehicle.name || '车辆详情'}
            </Breadcrumb.Item>
          </Breadcrumb>

          <Row gutter={24}>
            <Col span={16}>
              <Card>
                <div style={{ marginBottom: 24 }}>
                  <Row justify="space-between" align="middle">
                    <Col>
                      <Title level={2} style={{ margin: 0 }}>
                        {vehicle.name || `${vehicle.brand} ${vehicle.style}` || '未知车型'}
                      </Title>
                      <Text type="secondary" style={{ fontSize: '16px' }}>
                        {vehicle.brand} {vehicle.style}
                      </Text>
                    </Col>
                    <Col>
                      <Tag color={getStatusColor(vehicle.status)} style={{ fontSize: '14px', padding: '4px 12px' }}>
                        {VEHICLE_STATUS_LABELS[vehicle.status] || vehicle.status}
                      </Tag>
                    </Col>
                  </Row>
                </div>

                <ImageSection>
                  <Image
                    src={vehicle.images || '/images/my-car-a.jpg'}
                    alt={vehicle.name}
                    fallback="/images/my-car-a.jpg"
                    style={{ width: '100%', maxHeight: '400px', objectFit: 'cover' }}
                  />
                </ImageSection>

                <Divider />

                <Descriptions title="车辆信息" column={2}>
                  <Descriptions.Item label="车辆类型">
                    <FeatureTag color="blue">
                      {VEHICLE_TYPE_LABELS[vehicle.vehicle_type] || vehicle.vehicle_type}
                    </FeatureTag>
                  </Descriptions.Item>
                  <Descriptions.Item label="座位数">
                    <UserOutlined style={{ marginRight: 4 }} />
                    {vehicle.seats || 5}座
                  </Descriptions.Item>
                  <Descriptions.Item label="变速箱">
                    <SettingOutlined style={{ marginRight: 4 }} />
                    {vehicle.transmission || '自动'}
                  </Descriptions.Item>
                  <Descriptions.Item label="燃料类型">
                    {vehicle.fuel_type || '汽油'}
                  </Descriptions.Item>
                  <Descriptions.Item label="车牌号">
                    {vehicle.license_plate || '暂无'}
                  </Descriptions.Item>
                  <Descriptions.Item label="年份">
                    <CalendarOutlined style={{ marginRight: 4 }} />
                    {vehicle.year || '未知'}年
                  </Descriptions.Item>
                  <Descriptions.Item label="位置" span={2}>
                    <EnvironmentOutlined style={{ marginRight: 4 }} />
                    {vehicle.location || '暂无位置信息'}
                  </Descriptions.Item>
                </Descriptions>

                {vehicle.description && (
                  <>
                    <Divider />
                    <div>
                      <Title level={4}>车辆描述</Title>
                      <Paragraph>{vehicle.description}</Paragraph>
                    </div>
                  </>
                )}
              </Card>
            </Col>

            <Col span={8}>
              <PriceCard>
                <div style={{ textAlign: 'center', marginBottom: 24 }}>
                  <div>
                    <span className="price-text">¥{vehicle.price || 0}</span>
                    <span className="price-unit">/天</span>
                  </div>
                </div>

                <Space direction="vertical" style={{ width: '100%' }} size="middle">
                  <Button 
                    type="primary" 
                    size="large" 
                    block
                    disabled={vehicle.status !== 'available'}
                  >
                    立即预订
                  </Button>
                  
                  <Button size="large" block>
                    加入收藏
                  </Button>
                  
                  <Button size="large" block>
                    联系客服
                  </Button>
                </Space>

                <Divider />

                <div>
                  <Title level={5}>
                    <SafetyOutlined style={{ marginRight: 8 }} />
                    安全保障
                  </Title>
                  <ul style={{ paddingLeft: 20, margin: 0 }}>
                    <li>全车保险覆盖</li>
                    <li>24小时道路救援</li>
                    <li>定期维护保养</li>
                    <li>安全检测合格</li>
                  </ul>
                </div>
              </PriceCard>
            </Col>
          </Row>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default VehicleDetail;
