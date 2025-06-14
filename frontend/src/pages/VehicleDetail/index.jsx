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
  Image,
  Space,
  Breadcrumb,
  Divider,
  Carousel,
  Rate
} from 'antd';
import {
  HomeOutlined,
  CarOutlined,
  SafetyOutlined,
  EnvironmentOutlined,
  CalendarOutlined,
  LeftOutlined,
  RightOutlined,
  HeartOutlined,
  ShareAltOutlined,
  PhoneOutlined,
  MessageOutlined,
  StarFilled,
  CheckCircleOutlined,
  DashboardOutlined,
  SettingOutlined
} from '@ant-design/icons';
import { useParams, useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import vehicleService from '../../services/vehicleService';
import { VEHICLE_STATUS_LABELS } from '../../utils/constants';
import { getAllImages, handleImageError } from '../../utils/imageUtils';

const { Header, Content } = Layout;
const { Title, Text, Paragraph } = Typography;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
`;

const StyledHeader = styled(Header)`
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 0 50px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.15);
  position: fixed;
  width: 100%;
  z-index: 1000;
  top: 0;
  left: 0;
  backdrop-filter: blur(10px);
`;

const ContentWrapper = styled.div`
  margin-top: 64px;
  padding: 32px 50px;
  max-width: 1400px;
  margin-left: auto;
  margin-right: auto;
`;

const ImageSection = styled.div`
  margin-bottom: 32px;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
  position: relative;

  .ant-image {
    width: 100%;
    border-radius: 16px;
    overflow: hidden;
  }

  .ant-carousel {
    border-radius: 16px;
    overflow: hidden;

    .slick-dots {
      bottom: 20px;

      li button {
        background: rgba(255, 255, 255, 0.6);
        border-radius: 50%;
        width: 12px;
        height: 12px;
        border: 2px solid rgba(255, 255, 255, 0.8);
        transition: all 0.3s ease;
      }

      li.slick-active button {
        background: #667eea;
        transform: scale(1.2);
      }
    }

    .slick-prev,
    .slick-next {
      z-index: 10;
      width: 50px;
      height: 50px;
      background: rgba(255, 255, 255, 0.9);
      border-radius: 50%;
      display: flex !important;
      align-items: center;
      justify-content: center;
      transition: all 0.3s ease;

      &:hover {
        background: rgba(255, 255, 255, 1);
        transform: scale(1.1);
      }

      &:before {
        font-size: 18px;
        color: #667eea;
        font-weight: bold;
      }
    }

    .slick-prev {
      left: 20px;
    }

    .slick-next {
      right: 20px;
    }
  }

  .carousel-image {
    width: 100%;
    height: 500px;
    object-fit: cover;
    transition: transform 0.3s ease;

    &:hover {
      transform: scale(1.02);
    }
  }

  .single-image {
    width: 100%;
    max-height: 500px;
    object-fit: cover;
    border-radius: 16px;
    transition: transform 0.3s ease;

    &:hover {
      transform: scale(1.02);
    }
  }
`;

const PriceCard = styled(Card)`
  position: sticky;
  top: 100px;
  border-radius: 20px;
  box-shadow: 0 12px 40px rgba(0,0,0,0.15);
  border: none;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%);

  .ant-card-body {
    padding: 32px;
  }

  .price-text {
    font-size: 36px;
    font-weight: 700;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    margin: 0;
  }

  .price-unit {
    font-size: 18px;
    color: #64748b;
    margin-left: 8px;
    font-weight: 500;
  }
`;

const VehicleCard = styled(Card)`
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
  border: none;
  background: white;
  margin-bottom: 24px;

  .ant-card-body {
    padding: 32px;
  }
`;

const ActionButton = styled(Button)`
  border-radius: 12px;
  height: 48px;
  font-weight: 600;
  font-size: 16px;
  transition: all 0.3s ease;

  &.ant-btn-primary {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border: none;

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
    }
  }

  &:not(.ant-btn-primary) {
    border: 2px solid #e2e8f0;
    color: #64748b;

    &:hover {
      border-color: #667eea;
      color: #667eea;
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(102, 126, 234, 0.2);
    }
  }
`;

const StatusTag = styled(Tag)`
  padding: 8px 16px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 14px;
  border: none;

  &.status-available {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: white;
  }

  &.status-rented {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    color: white;
  }

  &.status-maintenance {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
    color: white;
  }
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
        if (response && response.code === 200 && response.data && response.data.vehicle) {
          setVehicle(response.data.vehicle);
        } else {
          setError(response?.data || '车辆信息不存在');
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
          <div style={{
            background: 'white',
            padding: '16px 24px',
            borderRadius: 16,
            marginBottom: 24,
            boxShadow: '0 2px 8px rgba(0,0,0,0.06)',
            border: '1px solid #e2e8f0'
          }}>
            <Breadcrumb
              separator=">"
              style={{
                fontSize: 16,
                fontWeight: 500
              }}
            >
              <Breadcrumb.Item href="/">
                <HomeOutlined style={{ color: '#667eea', marginRight: 4 }} />
                <span style={{ color: '#667eea' }}>首页</span>
              </Breadcrumb.Item>
              <Breadcrumb.Item href="/vehicles">
                <CarOutlined style={{ color: '#667eea', marginRight: 4 }} />
                <span style={{ color: '#667eea' }}>车辆列表</span>
              </Breadcrumb.Item>
              <Breadcrumb.Item>
                <span style={{ color: '#64748b' }}>
                  {vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '车辆详情'}
                </span>
              </Breadcrumb.Item>
            </Breadcrumb>
          </div>

          <Row gutter={32}>
            <Col span={16}>
              <VehicleCard>
                <div style={{ marginBottom: 32 }}>
                  <Row justify="space-between" align="middle">
                    <Col>
                      <Title level={1} style={{
                        margin: 0,
                        background: 'linear-gradient(135deg, #1e293b 0%, #475569 100%)',
                        WebkitBackgroundClip: 'text',
                        WebkitTextFillColor: 'transparent',
                        backgroundClip: 'text',
                        fontSize: '2.5rem',
                        fontWeight: 700
                      }}>
                        {vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '未知车型'}
                      </Title>
                      <Space size="middle" style={{ marginTop: 8 }}>
                        <Text style={{ fontSize: '18px', color: '#64748b', fontWeight: 500 }}>
                          <CalendarOutlined style={{ marginRight: 8, color: '#667eea' }} />
                          {vehicle.year}年
                        </Text>
                        <Text style={{ fontSize: '18px', color: '#64748b', fontWeight: 500 }}>
                          {vehicle.color || '未知颜色'}
                        </Text>
                        <Rate disabled defaultValue={4.5} style={{ fontSize: 16 }} />
                        <Text style={{ color: '#64748b' }}>(4.5分)</Text>
                      </Space>
                    </Col>
                    <Col>
                      <StatusTag className={`status-${vehicle.status}`}>
                        {VEHICLE_STATUS_LABELS[vehicle.status] || vehicle.status}
                      </StatusTag>
                    </Col>
                  </Row>
                </div>

                <ImageSection>
                  {(() => {
                    const images = getAllImages(vehicle.images);

                    if (images.length === 1) {
                      // 单张图片直接显示
                      return (
                        <Image
                          src={images[0]}
                          alt={vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '车辆图片'}
                          fallback="/images/my-car-a.jpg"
                          className="single-image"
                          onError={handleImageError}
                        />
                      );
                    } else {
                      // 多张图片使用轮播
                      return (
                        <Carousel
                          autoplay
                          autoplaySpeed={4000}
                          dots={true}
                          arrows={true}
                          prevArrow={<LeftOutlined />}
                          nextArrow={<RightOutlined />}
                        >
                          {images.map((imageUrl, index) => (
                            <div key={index}>
                              <img
                                src={imageUrl}
                                alt={`${vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '车辆图片'} ${index + 1}`}
                                className="carousel-image"
                                onError={handleImageError}
                              />
                            </div>
                          ))}
                        </Carousel>
                      );
                    }
                  })()}
                </ImageSection>

                <Divider style={{ margin: '32px 0', borderColor: '#e2e8f0' }} />

                <div style={{ marginBottom: 32 }}>
                  <Title level={3} style={{
                    marginBottom: 24,
                    color: '#1e293b',
                    fontWeight: 600
                  }}>
                    <CarOutlined style={{ marginRight: 12, color: '#667eea' }} />
                    车辆详细信息
                  </Title>

                  <Row gutter={[24, 16]}>
                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <CarOutlined style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>品牌</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.brand || '未知品牌'}
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>

                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #10b981 0%, #059669 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <SettingOutlined style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>型号</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.style || '未知型号'}
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>

                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #f59e0b 0%, #d97706 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <CalendarOutlined style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>年份</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.year || '未知'}年
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>

                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <StarFilled style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>颜色</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.color || '未知颜色'}
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>

                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #06b6d4 0%, #0891b2 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <DashboardOutlined style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>里程数</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.mileage ? `${vehicle.mileage}公里` : '暂无数据'}
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>

                    <Col span={12}>
                      <Card size="small" style={{
                        borderRadius: 12,
                        border: '1px solid #e2e8f0',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                      }}>
                        <Space>
                          <div style={{
                            width: 40,
                            height: 40,
                            borderRadius: '50%',
                            background: 'linear-gradient(135deg, #ef4444 0%, #dc2626 100%)',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center'
                          }}>
                            <EnvironmentOutlined style={{ color: 'white', fontSize: 18 }} />
                          </div>
                          <div>
                            <Text type="secondary" style={{ fontSize: 12 }}>位置</Text>
                            <div style={{ fontWeight: 600, color: '#1e293b' }}>
                              {vehicle.location || '暂无位置信息'}
                            </div>
                          </div>
                        </Space>
                      </Card>
                    </Col>
                  </Row>

                  {vehicle.contact && (
                    <Card style={{
                      marginTop: 16,
                      borderRadius: 12,
                      border: '1px solid #e2e8f0',
                      background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)'
                    }}>
                      <Space>
                        <div style={{
                          width: 40,
                          height: 40,
                          borderRadius: '50%',
                          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                          display: 'flex',
                          alignItems: 'center',
                          justifyContent: 'center'
                        }}>
                          <PhoneOutlined style={{ color: 'white', fontSize: 18 }} />
                        </div>
                        <div>
                          <Text type="secondary" style={{ fontSize: 12 }}>联系方式</Text>
                          <div style={{ fontWeight: 600, color: '#1e293b' }}>
                            {vehicle.contact}
                          </div>
                        </div>
                      </Space>
                    </Card>
                  )}
                </div>

                {vehicle.description && (
                  <>
                    <Divider style={{ margin: '32px 0', borderColor: '#e2e8f0' }} />
                    <div style={{
                      padding: '24px',
                      background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)',
                      borderRadius: 16,
                      border: '1px solid #e2e8f0'
                    }}>
                      <Title level={3} style={{
                        marginBottom: 16,
                        color: '#1e293b',
                        fontWeight: 600
                      }}>
                        <MessageOutlined style={{ marginRight: 12, color: '#667eea' }} />
                        车辆描述
                      </Title>
                      <Paragraph style={{
                        fontSize: 16,
                        lineHeight: 1.8,
                        color: '#475569',
                        margin: 0
                      }}>
                        {vehicle.description}
                      </Paragraph>
                    </div>
                  </>
                )}
              </VehicleCard>
            </Col>

            <Col span={8}>
              <PriceCard>
                <div style={{ textAlign: 'center', marginBottom: 32 }}>
                  <div style={{ marginBottom: 16 }}>
                    <span className="price-text">¥{vehicle.price || 0}</span>
                    <span className="price-unit">/天</span>
                  </div>
                  <Text style={{ color: '#64748b', fontSize: 14 }}>
                    价格包含基础保险
                  </Text>
                </div>

                <Space direction="vertical" style={{ width: '100%' }} size="large">
                  <ActionButton
                    type="primary"
                    size="large"
                    block
                    disabled={vehicle.status !== 'available'}
                    icon={<CheckCircleOutlined />}
                  >
                    {vehicle.status === 'available' ? '立即预订' : '暂不可用'}
                  </ActionButton>

                  <Row gutter={12}>
                    <Col span={12}>
                      <ActionButton size="large" block icon={<HeartOutlined />}>
                        收藏
                      </ActionButton>
                    </Col>
                    <Col span={12}>
                      <ActionButton size="large" block icon={<ShareAltOutlined />}>
                        分享
                      </ActionButton>
                    </Col>
                  </Row>

                  <ActionButton size="large" block icon={<MessageOutlined />}>
                    <PhoneOutlined style={{ marginRight: 8 }} />
                    联系客服
                  </ActionButton>
                </Space>

                <Divider style={{ margin: '32px 0', borderColor: '#e2e8f0' }} />

                <div>
                  <Title level={4} style={{
                    marginBottom: 20,
                    color: '#1e293b',
                    fontWeight: 600
                  }}>
                    <SafetyOutlined style={{ marginRight: 12, color: '#10b981' }} />
                    安全保障
                  </Title>

                  <Space direction="vertical" style={{ width: '100%' }} size="middle">
                    {[
                      { icon: <CheckCircleOutlined />, text: '全车保险覆盖', color: '#10b981' },
                      { icon: <CheckCircleOutlined />, text: '24小时道路救援', color: '#667eea' },
                      { icon: <CheckCircleOutlined />, text: '定期维护保养', color: '#f59e0b' },
                      { icon: <CheckCircleOutlined />, text: '安全检测合格', color: '#ef4444' }
                    ].map((item, index) => (
                      <div key={index} style={{
                        display: 'flex',
                        alignItems: 'center',
                        padding: '12px 16px',
                        background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)',
                        borderRadius: 12,
                        border: '1px solid #e2e8f0'
                      }}>
                        <div style={{
                          color: item.color,
                          marginRight: 12,
                          fontSize: 16
                        }}>
                          {item.icon}
                        </div>
                        <Text style={{ fontWeight: 500, color: '#1e293b' }}>
                          {item.text}
                        </Text>
                      </div>
                    ))}
                  </Space>
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
