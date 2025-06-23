import React, { useState } from 'react';
import { Layout, Card, Input, Button, Form, Row, Col, Typography, Space, message } from 'antd';
import { UserOutlined, SafetyOutlined, MobileOutlined, LockOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import instance from '../../utils/axiosConfig';
import { useNavigate } from 'react-router-dom';

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
  padding: 20px; /* 增加内边距 */

  // General height for all inputs and buttons
  .ant-input,
  .ant-btn {
    height: 45px;
  }

  // Apply border-radius and overflow hidden to input wrappers
  .ant-input-affix-wrapper, // For input with prefix/suffix
  .ant-input-group-wrapper { // For input with addonBefore/addonAfter
    border-radius: 8px;
    overflow: hidden;
  }

  // Styles specifically for Input with addonAfter to make it seamless
  .ant-input-group {
    display: flex;
    width: 100%;
    border: none; // Remove default border from the input group
    box-shadow: none; // Remove default shadow from the input group
  }

  .ant-input-group .ant-input {
    flex: 1; // Allow input to expand and take available space
    border-right: none; // Remove the border between input and addon
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
  }

  .ant-input-group-addon {
    background-color: transparent !important; // Ensure addon background is transparent
    border: none !important; // Remove default border from addon
    padding: 0 !important; // Remove default padding from addon
    display: flex !important; // Make addon a flex container for its button
    align-items: stretch !important; // Stretch button to full height
    white-space: nowrap !important; // Prevent text wrapping
    flex-shrink: 0 !important; // Prevent addon from shrinking
    min-width: fit-content !important; // Ensure addon has minimum width to fit content
  }

  .ant-input-group-addon .ant-btn {
    border-top-left-radius: 0 !important;
    border-bottom-left-radius: 0 !important;
    border-left: none !important; // Remove border from button side touching input
    margin-left: -1px !important; // Small negative margin to correct subpixel rendering gaps
    height: 100% !important; // Ensure button fills addon height
    width: auto !important; // Allow button to take natural width based on content
    padding: 0 15px !important; // Add some horizontal padding to ensure text is visible
  }
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
    cursor: pointer;
    transition: color 0.3s ease;

    &:hover {
      color: #1890ff;
    }
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
  const [sendingCode, setSendingCode] = useState(false);
  const [countdown, setCountdown] = useState(0);
  const navigate = useNavigate();

  // 处理点击ZuLMe标题返回首页
  const handleLogoClick = () => {
    navigate('/');
  };

  // 發送驗證碼
  const handleGetVerificationCode = async () => {
    try {
      const phone = form.getFieldValue('phone');
      if (!phone) {
        message.error('请输入手机号码');
        return;
      }

      // 驗證手機號碼格式
      if (!/^1[3-9]\d{9}$/.test(phone)) {
        message.error('请输入正确的手机号码格式');
        return;
      }

      setSendingCode(true);
      const response = await instance.post('/user/sendCode', { phone, source: "register" });
      
      if (response.data.code === 200) {
        message.success('验证码已发送，请注意查收');
        // 開始倒計時
        setCountdown(60);
        const timer = setInterval(() => {
          setCountdown((prev) => {
            if (prev <= 1) {
              clearInterval(timer);
              return 0;
            }
            return prev - 1;
          });
        }, 1000);
      } else {
        message.error(response.data.message || '发送验证码失败');
      }
    } catch (error) {
      message.error(error.response?.data?.message || '发送验证码失败，请稍后重试');
    } finally {
      setSendingCode(false);
    }
  };

  // 提交表單
  const onFinish = async (values) => {
    try {
      setLoading(true);
      const response = await instance.post('/user/register', {
        phone: values.phone,
        code: values.code,
        source: "register"
      });

      if (response.data.code === 200) {
        message.success('登入成功！');
        console.log('登入成功，后端返回:', response.data);
        console.log('準備儲存的 Token 值:', response.data.data.token);
        localStorage.setItem('token', response.data.data.token);
        console.log('登入成功後，localStorage 中的 token:', localStorage.getItem('token'));
        localStorage.setItem('userPhone', values.phone);
        navigate('/dashboard');
      } else {
        message.error(response.data.msg || '登入失败，请重试');
      }
    } catch (error) {
      message.error(error.response?.data?.msg || '登入失败，请重试');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Layout>
      <StyledHeader>
        <HeaderTitle level={3} onClick={handleLogoClick}>
          ZuLMe
        </HeaderTitle>
        <Space>
          <Button type="primary" style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}>English</Button>
        </Space>
      </StyledHeader>

      <StyledLayout>
        <LoginCard>
          <Title level={3} style={{ marginBottom: 24 }}>登入</Title>
          <Form
            form={form}
            name="login-register"
            initialValues={{ remember: true }}
            onFinish={onFinish}
          >
            <Form.Item
              name="phone"
              rules={[
                { required: true, message: '请输入手机号码！' },
                { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码格式!' }
              ]}
            >
              <Input prefix={<UserOutlined />} placeholder="请输入手机号码" />
            </Form.Item>

            <Form.Item
              name="code"
              rules={[
                { required: true, message: '请输入验证码!' },
                { len: 5, message: '验证码应为5位数字!' }
              ]}
            >
              <Input
                prefix={<SafetyOutlined />}
                placeholder="请输入验证码"
                addonAfter={
                  <Button 
                    type="primary" 
                    onClick={handleGetVerificationCode} 
                    loading={sendingCode}
                    disabled={countdown > 0}
                    style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}
                  >
                    {countdown > 0 ? `${countdown}秒后重试` : '获取验证码'}
                  </Button>
                }
              />
            </Form.Item>

            <Form.Item>
              <Button type="primary" htmlType="submit" loading={loading} block style={{ backgroundColor: '#1890ff', borderColor: '#1890ff', color: '#fff' }}>
                登入
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