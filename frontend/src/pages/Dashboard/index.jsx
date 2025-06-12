import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { 
  Layout, 
  Carousel, 
  Card, 
  Input, 
  DatePicker, 
  Button, 
  Row, 
  Col,
  Typography,
  Space,
  Select
} from 'antd';
import { SearchOutlined, CarOutlined, SafetyOutlined, ClockCircleOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Header, Content, Footer } = Layout;
const { Title, Paragraph } = Typography;
const { RangePicker } = DatePicker;
const { Option } = Select;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
`;

const StyledHeader = styled(Header)`
  background: #fff;
  padding: 0 50px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  position: fixed;
  width: 100%;
  z-index: 1;
  top: 0;
  left: 0;
`;

const StyledLoginRegisterButton = styled(Button)`
  background-color: #1890ff;
  border-color: #1890ff;
  color: #fff;

  &:hover {
    background-color: #40a9ff !important;
    border-color: #40a9ff !important;
    color: #fff !important;
  }

  &:active {
    background-color: #096dd9 !important;
    border-color: #096dd9 !important;
    color: #fff !important;
  }
`;

const HeroSectionWrapper = styled.div`
  position: relative;
  height: 700px; /* 增加英雄区域的高度 */
  margin-top: 64px; /* 补偿固定导航栏的高度 */
  overflow: hidden; /* 防止内容溢出 */
`;

const SearchContainer = styled.div`
  background: white;
  padding: 30px;
  border-radius: 8px;
  width: 80%;
  max-width: 1000px;
  position: absolute; /* 绝对定位 */
  bottom: 50px; /* 距离底部50px */
  left: 50%; /* 水平居中 */
  transform: translateX(-50%);
  z-index: 2; /* 确保在轮播图之上 */
  box-shadow: 0 4px 16px rgba(0,0,0,0.1); /* 添加阴影效果 */
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
  const [searchParams, setSearchParams] = useState({
    location: '',
    dates: null,
    carType: undefined
  });

  const navigate = useNavigate();

  const handleSearch = () => {
    console.log('Search params:', searchParams);
    // TODO: Implement search functionality
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('userPhone'); // 移除保存的手机号
    navigate('/login-register');
  };

  const userPhone = localStorage.getItem('userPhone'); // 从 localStorage 获取手机号

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={3} style={{ margin: 0 }}>ZuLMe</Title>
          </Col>
          <Col>
            <Space>
              {userPhone ? (
                <>
                  <Button type="link" style={{ color: '#000' }}>{userPhone}</Button>
                  <Button type="primary" onClick={() => navigate('/personal-center')}>我的</Button>
                  <Button type="default" onClick={handleLogout}>退出</Button>
                </>
              ) : (
                <StyledLoginRegisterButton type="primary" onClick={() => navigate('/login-register')}>登录/注册</StyledLoginRegisterButton>
              )}
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <HeroSectionWrapper>
          {/* 顶部轮播图 */}
          <Carousel autoplay dotPosition="bottom" style={{ height: '50%' }}>
            <div>
              <img src="/images/banner1.jpg" alt="banner1" style={{ width: '100%', height: '50%', objectFit: 'cover' }} />
            </div>
            <div>
              <img src="/images/banner2.jpg" alt="banner2" style={{ width: '100%', height: '50%', objectFit: 'cover' }} />
            </div>
            <div>
              <img src="/images/banner3.jpg" alt="banner3" style={{ width: '100%', height: '50%', objectFit: 'cover' }} />
            </div>
            <div>
              <img src="/images/banner4.jpg" alt="banner4" style={{ width: '100%', height: '50%', objectFit: 'cover' }} />
            </div>
          </Carousel>

          {/* 搜索框 */}
          <SearchContainer>
            <Title level={2} style={{ textAlign: 'center', marginBottom: 24 }}>
              找到您的理想座驾
            </Title>
            <Row gutter={16}>
              <Col span={8}>
                <Input
                  placeholder="取车地点"
                  prefix={<SearchOutlined />}
                  value={searchParams.location}
                  onChange={e => setSearchParams({...searchParams, location: e.target.value})}
                />
              </Col>
              <Col span={8}>
                <RangePicker
                  style={{ width: '100%' }}
                  onChange={dates => setSearchParams({...searchParams, dates})}
                />
              </Col>
              <Col span={4}>
                <Select
                  style={{ width: '100%' }}
                  placeholder="车型"
                  onChange={value => setSearchParams({...searchParams, carType: value})}
                >
                  <Option value="economy">经济型</Option>
                  <Option value="comfort">舒适型</Option>
                  <Option value="luxury">豪华型</Option>
                  <Option value="suv">SUV</Option>
                </Select>
              </Col>
              <Col span={4}>
                <Button type="primary" block onClick={handleSearch}>
                  搜索
                </Button>
              </Col>
            </Row>
          </SearchContainer>
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
                cover={<img alt="car" src="/images/my-car-d.jpg" />}
              >
                <Card.Meta
                  title="SUV"
                  description="¥1299/天起"
                />
              </Card>
            </Col>
          </Row>
        </div>
      </Content>

      <Footer style={{ textAlign: 'center', background: '#001529', color: '#fff', padding: '24px 50px' }}>
        ZuLMe ©2023 Created by Ant Design
      </Footer>
    </StyledLayout>
  );
};

export default Dashboard; 