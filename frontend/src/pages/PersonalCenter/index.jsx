import React, { useState } from 'react';
import { Layout, Menu, Button, Card, Typography, Row, Col, Input, Space, Tabs } from 'antd';
import { UserOutlined, FileTextOutlined, AccountBookOutlined, WalletOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import MyInfoPage from './components/MyInfoPage'; // 导入 MyInfoPage

const { Header, Sider, Content } = Layout;
const { Title, Text } = Typography;
const { Search } = Input;
const { TabPane } = Tabs;

const PersonalCenter = () => {
    const navigate = useNavigate();
    const phoneNumber = localStorage.getItem('phoneNumber'); // 从 localStorage 获取手机号
    const [selectedContentKey, setSelectedContentKey] = useState('my_orders_short_term'); // 新增状态来管理右侧内容

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('phoneNumber');
        navigate('/login-register'); // 返回登录/注册页面
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
                            <Button type="primary">开发票</Button>
                        </Space>
                    )}>
                        <Tabs defaultActiveKey="1" style={{ marginBottom: 24 }}
                            items={[
                                { label: '全部', key: '1' },
                                { label: '处理中', key: '2' },
                                { label: '等待付款', key: '3' },
                                { label: '预订成功', key: '4' },
                                { label: '租赁中', key: '5' },
                                { label: '已完成', key: '6' },
                                { label: '已取消', key: '7' },
                            ]}
                        />
                        <div style={{ textAlign: 'center', padding: '50px' }}>
                            <img src="https://gw.alipayobjects.com/zos/antfincdn/ZH9JzZhHw$/empty.svg" alt="Empty" style={{ width: '100px' }} />
                            <Text style={{ display: 'block', marginTop: '16px' }}>还没租过车？速速体验吧</Text>
                        </div>
                        <Button type="primary" style={{ marginTop: '24px' }}>去租车</Button>
                        <div style={{ marginTop: '40px', borderTop: '1px solid #f0f0f0', paddingTop: '20px' }}>
                            <Title level={5}>订单状态说明</Title>
                            <Row gutter={16}>
                                <Col span={6}>
                                    <Text strong>预订成功</Text><br/>
                                    <Text type="secondary">订单已确认，尚未提车</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>等待付款</Text><br/>
                                    <Text type="secondary">订单尚未付款，请在1小时内完成支付</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>处理中</Text><br/>
                                    <Text type="secondary">您的订单已提交，正在处理中</Text>
                                </Col>
                                <Col span={6}>
                                    <Text strong>租赁中</Text><br/>
                                    <Text type="secondary">车辆正在租赁过程中</Text>
                                </Col>
                            </Row>
                        </div>
                    </Card>
                );
            case 'my_account_info':
                return <MyInfoPage />;
            // 可以添加更多 case 来渲染其他内容
            default:
                return null;
        }
    };

    return (
        <Layout style={{ minHeight: '100vh', backgroundColor: '#f0f2f5' }}>
            <Header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#fff', padding: '0 50px', borderBottom: '1px solid #f0f0f0' }}>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                    <Title level={3} style={{ margin: 0, marginRight: '20px', whiteSpace: 'nowrap' }}>ZuLMe</Title>
                    {/* 移除右側的導航連結 */}
                    
                    <Menu
                        mode="horizontal"
                        defaultSelectedKeys={['home']}
                        style={{ borderBottom: 'none', lineHeight: '64px' }}
                        items={[
                            {
                                key: 'home',
                                label: '首页',
                                onClick: () => navigate('/dashboard'),
                            },
                        ]}
                    />
                    
                </div>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                    {/* <Text style={{ marginRight: '16px', whiteSpace: 'nowrap' }}>400-616-6666</Text> */}
                    <Button type="link" style={{ marginRight: '8px' }}>English</Button>
                    {phoneNumber && (
                        <Space>
                            <Text style={{ marginRight: '8px', whiteSpace: 'nowrap' }}>{phoneNumber ? phoneNumber.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2') : ''}</Text>
                            <Button type="primary" onClick={() => navigate('/personal-center')}>我的</Button>
                            <Button onClick={handleLogout}>退出</Button>
                        </Space>
                    )}
                </div>
            </Header>
            <Content style={{ padding: '24px 50px' }}>
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
                                            label: '短租自驾(0)',
                                        },
                                        {
                                            key: 'my_orders_shunfeng',
                                            icon: <FileTextOutlined />,
                                            label: '顺风车(0)',
                                        },
                                        {
                                            key: 'my_orders_international',
                                            icon: <FileTextOutlined />,
                                            label: '国际租车(0)',
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
                                            key: 'my_assets_stored_value',
                                            icon: <WalletOutlined />,
                                            label: '储值卡(0元)',
                                        },
                                        {
                                            key: 'my_assets_balance',
                                            icon: <WalletOutlined />,
                                            label: '账户余额(0元)',
                                        },
                                        {
                                            key: 'my_assets_bank_cards',
                                            icon: <WalletOutlined />,
                                            label: '银行卡(0张)',
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
        </Layout>
    );
};

export default PersonalCenter; 