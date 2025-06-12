import React from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Layout,
  Carousel,
  Card,
  Button,
  Row,
  Col,
  Typography,
  Space
} from 'antd';
import { CarOutlined, SafetyOutlined, ClockCircleOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Header, Content, Footer } = Layout;
const { Title, Paragraph } = Typography;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
`;

const StyledHeader = styled(Header)`
  background: rgba(0, 0, 0, 0.9);
  backdrop-filter: blur(10px);
  padding: 0 50px;
  box-shadow: 0 2px 20px rgba(0,0,0,0.15);
  position: fixed;
  width: 100%;
  z-index: 1000;
  top: 0;
  left: 0;
  height: 70px;
  line-height: 70px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
`;

const StyledNavButton = styled(Button)`
  background: transparent;
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
  font-size: 16px;
  font-weight: 500;
  padding: 8px 24px;
  height: 40px;
  border-radius: 6px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;

  &:before {
    content: '';
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.2), transparent);
    transition: left 0.5s;
  }

  &:hover {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
    border: 1px solid #667eea !important;
    color: #fff !important;
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);

    &:before {
      left: 100%;
    }
  }

  &:active {
    transform: translateY(0);
    box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
  }
`;

const StyledUserButton = styled(Button)`
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #fff;
  font-size: 16px;
  font-weight: 500;
  padding: 8px 24px;
  height: 40px;
  border-radius: 6px;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);

  &:hover {
    background: linear-gradient(135deg, #764ba2 0%, #667eea 100%) !important;
    border: 1px solid rgba(255, 255, 255, 0.25) !important;
    color: #fff !important;
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
  }

  &:active {
    transform: translateY(0);
    box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
  }
`;

const HeroSectionWrapper = styled.div`
  position: relative;
  height: 700px; /* 增加英雄区域的高度 */
  margin-top: 70px; /* 补偿固定导航栏的高度 */
  overflow: hidden; /* 防止内容溢出 */

  .ant-carousel {
    height: 100%;

    .ant-carousel-inner {
      height: 100%;
    }

    .slick-slide {
      height: 700px;

      > div {
        height: 100%;

        img {
          width: 100% !important;
          height: 700px !important;
          object-fit: cover !important;
          object-position: center !important;
          display: block !important;
        }
      }
    }

    .slick-dots {
      bottom: 20px;
      z-index: 10;

      li button {
        background: rgba(255, 255, 255, 0.5);
        border-radius: 50%;
        width: 12px;
        height: 12px;
        border: none;
        opacity: 0.7;
        transition: all 0.3s ease;
      }

      li.slick-active button {
        background: #fff;
        opacity: 1;
        transform: scale(1.2);
      }

      li button:hover {
        background: #fff;
        opacity: 0.9;
      }
    }
  }
`;

const StyledTitle = styled(Title)`
  margin: 0 !important;
  color: #fff !important;
  font-size: 28px !important;
  font-weight: 700 !important;
  letter-spacing: 1px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
`;

const FeatureCard = styled(Card)`
  text-align: center;
  height: 100%;
  .anticon {
    font-size: 48px;
    color: #1890ff;
    margin-bottom: 16px;
  }
`;

const Dashboard = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('userPhone');
    navigate('/login-register');
  };

  const userPhone = localStorage.getItem('userPhone');

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle" style={{ height: '100%' }}>
          <Col>
            <StyledTitle level={3}>ZuLMe</StyledTitle>
          </Col>
          <Col>
            <Space size="large">
              <StyledNavButton onClick={() => navigate('/vehicles')}>
                租豪车
              </StyledNavButton>
              {userPhone ? (
                <>
                  <StyledUserButton type="link" style={{ color: '#fff' }}>{userPhone}</StyledUserButton>
                  <StyledUserButton onClick={() => navigate('/personal-center')}>我的</StyledUserButton>
                  <StyledUserButton onClick={handleLogout}>退出</StyledUserButton>
                </>
              ) : (
                <StyledUserButton onClick={() => navigate('/login-register')}>登录/注册</StyledUserButton>
              )}
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <HeroSectionWrapper>
          {/* 顶部轮播图 */}
          <Carousel autoplay dotPosition="bottom" style={{ height: '100%' }}>
            <div>
              <img src="/images/banner5.png" alt="banner5" />
            </div>
            <div>
              <img src="/images/banner6.png" alt="banner6" />
            </div>
            <div>
              <img src="/images/banner7.png" alt="banner7" />
            </div>
            <div>
              <img src="/images/banner8.png" alt="banner8" />
            </div>
          </Carousel>


        </HeroSectionWrapper>

        <div style={{ padding: '50px 50px' }}>
          <Title level={2} style={{ textAlign: 'center', marginBottom: 40 }}>
            为什么选择我们
          </Title>
          <Row gutter={24}>
            <Col span={8}>
              <FeatureCard>
                <CarOutlined />
                <Title level={4}>丰富车型</Title>
                <Paragraph>
                  提供多种车型选择，满足不同需求
                </Paragraph>
              </FeatureCard>
            </Col>
            <Col span={8}>
              <FeatureCard>
                <SafetyOutlined />
                <Title level={4}>安全保障</Title>
                <Paragraph>
                  车辆定期维护，保险齐全
                </Paragraph>
              </FeatureCard>
            </Col>
            <Col span={8}>
              <FeatureCard>
                <ClockCircleOutlined />
                <Title level={4}>24小时服务</Title>
                <Paragraph>
                  全天候客服支持，随时为您服务
                </Paragraph>
              </FeatureCard>
            </Col>
          </Row>
        </div>

        <div style={{ background: '#f0f2f5', padding: '50px 50px' }}>
          <Title level={2} style={{ textAlign: 'center', marginBottom: 40 }}>
            热门车型
          </Title>
          <Row gutter={24}>
            <Col span={6} key="car-a">
              <Card
                hoverable
                cover={<img alt="car" src="/images/my-car-a.jpg" />}
                onClick={() => navigate('/vehicles')}
                style={{ cursor: 'pointer' }}
              >
                <Card.Meta
                  title="经济轿跑"
                  description="¥599/天起"
                />
              </Card>
            </Col>
            <Col span={6} key="car-b">
              <Card
                hoverable
                cover={<img alt="car" src="/images/my-car-b.jpg" />}
                onClick={() => navigate('/vehicles')}
                style={{ cursor: 'pointer' }}
              >
                <Card.Meta
                  title="舒适轿跑"
                  description="¥899/天起"
                />
              </Card>
            </Col>
            <Col span={6} key="car-c">
              <Card
                hoverable
                cover={<img alt="car" src="/images/my-car-c.jpg" />}
                onClick={() => navigate('/vehicles')}
                style={{ cursor: 'pointer' }}
              >
                <Card.Meta
                  title="豪华轿跑"
                  description="¥1499/天起"
                />
              </Card>
            </Col>
            <Col span={6} key="car-d">
              <Card
                hoverable
                cover={<img alt="car" src="/images/my-car-e.jpg" />}
                onClick={() => navigate('/vehicles')}
                style={{ cursor: 'pointer' }}
              >
                <Card.Meta
                  title="豪华车型"
                  description="¥2999/天起"
                />
              </Card>
            </Col>
          </Row>
        </div>
      </Content>

      <Footer style={{ textAlign: 'center' }}>
        ZuLMe ©2024 Created by Your Company
      </Footer>
    </StyledLayout>
  );
};

export default Dashboard; 