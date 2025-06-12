import React, { useState } from 'react';
import { Layout, Card, Input, Button, Form, Row, Col, Typography, Space } from 'antd';
import { UserOutlined, SafetyOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Header, Content, Footer } = Layout;
const { Title, Paragraph } = Typography;

const StyledLayout = styled(Layout)`
  min-height: 80vh;
  background-image: url('/images/my-car-e.jpg'); /* 请确保图片路径正确 */
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const LoginCard = styled(Card)`
  width: 400px;
  text-align: center;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
`;

const StyledHeader = styled(Header)`
  background: #000;
  padding: 0 50px;
  height: 64px;
  line-height: 64px;
  position: fixed;
  width: 100%;
  z-index: 1;
  top: 0;
  left: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const HeaderTitle = styled(Title)`
  && { /* Use && to increase specificity */
    color: #fff;
    margin: 0;
    display: flex;
    align-items: center;
  }
`;

const FooterSection = styled(Footer)`
  background-color: #fff;
  color: #000;
  padding: 24px 50px;
  text-align: center;
  width: 100%;
  position: absolute;
  bottom: 0;
  left: 0;
`;

const FooterColumns = styled(Row)`
  margin-bottom: 24px;
`;

const FooterLink = styled.a`
  color: #000;
  display: block;
  margin-bottom: 8px;
  &:hover {
    color: #1890ff;
  }
`;

const QrCodeContainer = styled.div`
  text-align: center;
  img {
    width: 100px;
    height: 100px;
    margin-bottom: 8px;
  }
  p {
    color: #000;
    margin: 0;
  }
`;

const FeatureSection = styled.div`
  background-color: #f0f2f5; /* Light gray background */
  padding: 50px 50px; /* Adjust padding as needed */
  text-align: center;
`;

const FeatureCard = styled(Card)`
  text-align: center;
  height: 100%;
  .anticon {
    font-size: 48px;
    color: #FFC107; /* Yellow color for icons */
    margin-bottom: 16px;
  }
  .ant-card-meta-title {
    color: #000; /* Black color for titles */
  }
  .ant-card-meta-description {
    color: #666; /* Darker gray for description */
  }
`;

const LoginRegister = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const onFinish = (values) => {
    setLoading(true);
    console.log('Received values of form: ', values);
    // TODO: Implement actual login/register logic here
    setTimeout(() => {
      setLoading(false);
      // Example: If login successful, redirect
      // history.push('/dashboard'); 
    }, 2000);
  };

  const handleGetVerificationCode = () => {
    const phone = form.getFieldValue('phone');
    if (!phone) {
      alert('请输入手机号');
      return;
    }
    console.log('Getting verification code for:', phone);
    // TODO: Implement send verification code logic here
  };

  return (
    <Layout>
      <StyledHeader>
        <HeaderTitle level={3}>
          ZuLMe
        </HeaderTitle>
        <Space>
          <Button type="primary" style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}>English</Button>
        </Space>
      </StyledHeader>

      <StyledLayout>
        <LoginCard>
          <Title level={3} style={{ marginBottom: 24 }}>登录</Title>
          <Form
            form={form}
            name="login-register"
            initialValues={{ remember: true }}
            onFinish={onFinish}
          >
            <Form.Item
              name="phone"
              rules={[{ required: true, message: '请输入注册手机号!' }]}
            >
              <Input prefix={<UserOutlined />} placeholder="请输入注册手机号" />
            </Form.Item>

            <Form.Item
              name="code"
              rules={[{ required: true, message: '请输入动态验证码!' }]}
            >
              <Input
                prefix={<SafetyOutlined />}
                placeholder="请输入动态验证码"
                addonAfter={<Button type="primary" onClick={handleGetVerificationCode} style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}>获取手机动态验证码</Button>}
              />
            </Form.Item>

            <Form.Item>
              <Button type="primary" htmlType="submit" loading={loading} block style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}>
                登录
              </Button>
            </Form.Item>
          </Form>
        </LoginCard>
      </StyledLayout>

      <FeatureSection>
        <Row gutter={24}>
          <Col span={6}>
            <FeatureCard>
              <img src="/images/car_icon.png" alt="100+车型" style={{ width: 48, height: 48, marginBottom: 16 }} />
              <Title level={4}>100+车型</Title>
              <Paragraph>
                提供多种车型选择，满足不同需求
              </Paragraph>
            </FeatureCard>
          </Col>
          <Col span={6}>
            <FeatureCard>
              <img src="/images/location_icon.png" alt="千家网点" style={{ width: 48, height: 48, marginBottom: 16 }} />
              <Title level={4}>千家网点</Title>
              <Paragraph>
                全国覆盖，取还方便
              </Paragraph>
            </FeatureCard>
          </Col>
          <Col span={6}>
            <FeatureCard>
              <img src="/images/safety_icon.png" alt="100%保障" style={{ width: 48, height: 48, marginBottom: 16 }} />
              <Title level={4}>100%保障</Title>
              <Paragraph>
                安全可靠，全程无忧
              </Paragraph>
            </FeatureCard>
          </Col>
          <Col span={6}>
            <FeatureCard>
              <img src="/images/mileage_icon.png" alt="无限里程" style={{ width: 48, height: 48, marginBottom: 16 }} />
              <Title level={4}>无限里程</Title>
              <Paragraph>
                尽情驰骋，无里程限制
              </Paragraph>
            </FeatureCard>
          </Col>
        </Row>
      </FeatureSection>

    </Layout>
  );
};

export default LoginRegister; 