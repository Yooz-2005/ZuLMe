import React, { useState, useEffect } from 'react';
import { Layout, Menu, Button, Card, Typography, Row, Col, Input, Space, Tabs } from 'antd';
import { UserOutlined, FileTextOutlined, AccountBookOutlined, WalletOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

import MyInfoPage from './components/MyInfoPage'; // 导入 MyInfoPage
import ReservationList from '../../components/ReservationList'; // 导入预订列表组件
import OrderList from '../../components/OrderList'; // 导入订单列表组件
import styled from 'styled-components';

const { Header, Sider, Content, Footer } = Layout; // 添加 Footer
const { Title, Text } = Typography;
const { Search } = Input;
const { TabPane } = Tabs;

// 定义 styled-components
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

const PersonalCenter = () => {
    const navigate = useNavigate();
    const [currentPhoneNumber, setCurrentPhoneNumber] = useState(localStorage.getItem('userPhone')); // 将手机号作为状态管理
    const [selectedContentKey, setSelectedContentKey] = useState('my_orders_short_term'); // 新增状态来管理右侧内容

    // 监听 localStorage 中的 userPhone 变化，并更新 state
    useEffect(() => {
        const handleStorageChange = () => {
            setCurrentPhoneNumber(localStorage.getItem('userPhone'));
        };

        window.addEventListener('storage', handleStorageChange);

        return () => {
            window.removeEventListener('storage', handleStorageChange);
        };
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('userPhone');
        setCurrentPhoneNumber(null); // 清空手机号状态
        navigate('/login-register'); // 返回登录/注册页面
    };

    const handlePhoneUpdate = (newPhoneNumber) => {
        setCurrentPhoneNumber(newPhoneNumber);
    };



    const renderContent = () => {
        switch (selectedContentKey) {
            case 'my_orders_short_term':
                return (
                    <Card title="短租自驾" extra={(
                        <Space>
                            <Search placeholder="请输入订单号" onSearch={() => {}} style={{ width: 200 }} />
                            <Button>查询</Button>
                            <Button>清除</Button>
                        </Space>
                    )}>
                        <Tabs
                            defaultActiveKey="all"
                            style={{ marginBottom: 24 }}
                            items={[
                                {
                                    label: '全部',
                                    key: 'all',
                                    children: <ReservationList activeTab="all" />
                                },
                                {
                                    label: '处理中',
                                    key: 'processing',
                                    children: <ReservationList activeTab="processing" />
                                },
                                {
                                    label: '等待付款',
                                    key: 'pending_payment',
                                    children: <ReservationList activeTab="pending_payment" />
                                },
                                {
                                    label: '预订成功',
                                    key: 'confirmed',
                                    children: <ReservationList activeTab="confirmed" />
                                },
                                {
                                    label: '租赁中',
                                    key: 'in_use',
                                    children: <ReservationList activeTab="in_use" />
                                },
                                {
                                    label: '已完成',
                                    key: 'completed',
                                    children: <ReservationList activeTab="completed" />
                                },
                                {
                                    label: '已取消',
                                    key: 'cancelled',
                                    children: <ReservationList activeTab="cancelled" />
                                },
                            ]}
                        />
                        <div style={{ marginTop: '40px', borderTop: '1px solid #f0f0f0', paddingTop: '20px' }}>
                            <Title level={5}>预订状态说明</Title>
                            <Row gutter={16}>
                                <Col span={6}>
                                    <Text strong>预订成功</Text><br/>
                                    <Text type="secondary">预订已确认，尚未提车</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>等待付款</Text><br/>
                                    <Text type="secondary">预订尚未付款，请及时完成支付</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>处理中</Text><br/>
                                    <Text type="secondary">您的预订已提交，正在处理中</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>租赁中</Text><br/>
                                    <Text type="secondary">车辆正在租赁过程中</Text>
                                </Col>
                            </Row>
                        </div>
                    </Card>
                );
            case 'my_orders_management':
                return (
                    <Card title="订单管理" extra={(
                        <Space>
                            <Search placeholder="请输入订单号" onSearch={() => {}} style={{ width: 200 }} />
                            <Button>查询</Button>
                            <Button>清除</Button>
                        </Space>
                    )}>
                        <Tabs
                            defaultActiveKey="all"
                            style={{ marginBottom: 24 }}
                            items={[
                                {
                                    label: '全部订单',
                                    key: 'all',
                                    children: <OrderList activeTab="all" />
                                },
                                {
                                    label: '待支付',
                                    key: 'pending_payment',
                                    children: <OrderList activeTab="pending_payment" />
                                },
                                {
                                    label: '已支付',
                                    key: 'paid',
                                    children: <OrderList activeTab="paid" />
                                },
                                {
                                    label: '使用中',
                                    key: 'in_use',
                                    children: <OrderList activeTab="in_use" />
                                },
                                {
                                    label: '已完成',
                                    key: 'completed',
                                    children: <OrderList activeTab="completed" />
                                },
                                {
                                    label: '已取消',
                                    key: 'cancelled',
                                    children: <OrderList activeTab="cancelled" />
                                },
                            ]}
                        />
                        <div style={{ marginTop: '40px', borderTop: '1px solid #f0f0f0', paddingTop: '20px' }}>
                            <Title level={5}>订单状态说明</Title>
                            <Row gutter={16}>
                                <Col span={6}>
                                    <Text strong>待支付</Text><br/>
                                    <Text type="secondary">订单已创建，等待支付</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>已支付</Text><br/>
                                    <Text type="secondary">支付成功，等待取车</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>使用中</Text><br/>
                                    <Text type="secondary">车辆正在使用中</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>已完成</Text><br/>
                                    <Text type="secondary">订单已完成</Text>
                                </Col>
                            </Row>
                        </div>
                    </Card>
                );
            case 'my_account_info':
                return <MyInfoPage onPhoneUpdate={handlePhoneUpdate} />;
            // 可以添加更多 case 来渲染其他内容
            default:
                return null;
        }
    };

    return (
        <StyledLayout>
            <StyledHeader>
                <Row justify="space-between" align="middle" style={{ height: '100%' }}>
                    <Col>
                        <StyledTitle level={3}>ZuLMe</StyledTitle>
                    </Col>
                    <Col>
                        <Space size="large">
                            <StyledNavButton onClick={() => navigate('/dashboard')}>
                                首页
                            </StyledNavButton>
                            {currentPhoneNumber ? (
                                <>
                                    <StyledUserButton type="link" style={{ color: '#fff' }}>{currentPhoneNumber}</StyledUserButton>
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
            <Content style={{ padding: '24px 50px', marginTop: '70px' }}> {/* 添加 marginTop 以避免被固定 Header 遮挡 */}
                <Layout style={{ background: '#fff' }}>
                    <Sider width={200} style={{ background: '#fff', borderRight: '1px solid #f0f0f0' }}>
                        <Menu
                            mode="inline"
                            defaultSelectedKeys={['my_orders_short_term']}
                            style={{ height: '100%', borderRight: 0 }}
                            onSelect={({ key }) => setSelectedContentKey(key)} // 添加 onSelect 事件
                            items={[
                                {
                                    key: 'my_orders_group',
                                    label: '我的订单',
                                    type: 'group',
                                    children: [
                                        {
                                            key: 'my_orders_short_term',
                                            icon: <FileTextOutlined />,
                                            label: '预订管理',
                                        },
                                        {
                                            key: 'my_orders_management',
                                            icon: <AccountBookOutlined />,
                                            label: '订单管理',
                                        },
                                    ],
                                },
                                {
                                    key: 'my_assets_group',
                                    label: '我的资产',
                                    type: 'group',
                                    children: [
                                        {
                                            key: 'my_assets_points',
                                            icon: <WalletOutlined />,
                                            label: '可用积分',
                                        },
                                        {
                                            key: 'my_assets_coupons',
                                            icon: <WalletOutlined />,
                                            label: '优惠券(0张)',
                                        },
                                        {
                                            key: 'my_assets_balance',
                                            icon: <WalletOutlined />,
                                            label: '账户余额(0元)',
                                        },
                                    ],
                                },
                                {
                                    key: 'my_account_group',
                                    label: '我的账户',
                                    type: 'group',
                                    children: [
                                        {
                                            key: 'my_account_info',
                                            icon: <UserOutlined />,
                                            label: '我的信息',
                                        },
                                        {
                                            key: 'my_account_level',
                                            icon: <UserOutlined />,
                                            label: '我的等级',
                                        },
                                        {
                                            key: 'my_account_payment_code',
                                            icon: <UserOutlined />,
                                            label: '支付密码',
                                        },
                                        {
                                            key: 'my_account_driving_license',
                                            icon: <UserOutlined />,
                                            label: '驾照认证',
                                        },
                                    ],
                                },
                                {
                                    key: 'invoice_management_group',
                                    label: '发票管理',
                                    type: 'group',
                                    children: [
                                        {
                                            key: 'invoice_apply',
                                            icon: <MailOutlined />,
                                            label: '根据订单开发票',
                                        },
                                        {
                                            key: 'invoice_history',
                                            icon: <MailOutlined />,
                                            label: '开票历史记录',
                                        },
                                    ],
                                },
                            ]}
                        />
                    </Sider>
                    <Content
                        style={{
                            padding: '24px',
                            margin: 0,
                            minHeight: 280,
                            borderLeft: '1px solid #f0f0f0',
                        }}
                    >
                        {renderContent()}{/* 调用渲染函数 */}
                    </Content>
                </Layout>
            </Content>
        </StyledLayout>
    );
};

export default PersonalCenter; 