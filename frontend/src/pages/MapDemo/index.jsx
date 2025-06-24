import React, { useState, useEffect } from 'react';
import { Layout, Typography, Card, Row, Col, Space, Button, message, Tabs, Divider, Collapse } from 'antd';
import { EnvironmentOutlined, CarOutlined, ShopOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import AmapComponent from '../../components/AmapComponent';
import SimpleMapComponent from '../../components/SimpleMapComponent';
import locationService from '../../services/locationService';
import vehicleService from '../../services/vehicleService';

const { Header, Content } = Layout;
const { Title, Text, Paragraph } = Typography;
const { TabPane } = Tabs;

// 样式组件
const StyledLayout = styled(Layout)`
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
`;

const StyledHeader = styled(Header)`
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 0 50px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.15);
`;

const ContentWrapper = styled.div`
  padding: 24px 50px;
  max-width: 1400px;
  margin: 0 auto;
`;

const InfoCard = styled(Card)`
  margin-bottom: 16px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);

  .ant-collapse {
    border: none;
    background: transparent;

    .ant-collapse-item {
      border: 1px solid #f0f0f0;
      border-radius: 8px;
      margin-bottom: 8px;

      &:last-child {
        margin-bottom: 0;
      }

      .ant-collapse-header {
        padding: 8px 12px;
        font-weight: 500;
        background: #fafafa;
        border-radius: 8px 8px 0 0;

        &:hover {
          background: #f0f0f0;
        }
      }

      .ant-collapse-content {
        border-top: 1px solid #f0f0f0;

        .ant-collapse-content-box {
          padding: 12px;
        }
      }
    }
  }

  /* 自定义滚动条样式 */
  .merchants-scroll {
    &::-webkit-scrollbar {
      width: 6px;
    }

    &::-webkit-scrollbar-track {
      background: #f1f1f1;
      border-radius: 3px;
    }

    &::-webkit-scrollbar-thumb {
      background: #c1c1c1;
      border-radius: 3px;

      &:hover {
        background: #a8a8a8;
      }
    }
  }
`;

const MapDemo = () => {
  const navigate = useNavigate();
  const [selectedLocation, setSelectedLocation] = useState(null);
  const [merchants, setMerchants] = useState([]);
  const [activeTab, setActiveTab] = useState('simple');
  const [loading, setLoading] = useState(true);

  // 获取真实的商家网点数据
  const fetchMerchants = async () => {
    try {
      setLoading(true);
      console.log('开始获取商家网点数据...');

      const response = await vehicleService.getLocationList({
        page: 1,
        pageSize: 100, // 获取所有网点
        status_filter: 1 // 只获取审核通过的商户
      });

      console.log('商家网点API响应:', response);

      if (response && response.data && response.data.merchants) {
        // 转换数据格式以适配地图组件
        const merchantsData = response.data.merchants.map(merchant => ({
          id: merchant.ID || merchant.id,
          name: merchant.Name || merchant.name || `网点${merchant.ID || merchant.id}`,
          longitude: parseFloat(merchant.Longitude || merchant.longitude || 0),
          latitude: parseFloat(merchant.Latitude || merchant.latitude || 0),
          address: merchant.Location || merchant.location || '地址未设置',
          phone: merchant.Phone || merchant.phone || '电话未设置',
          businessTime: merchant.BusinessTime || merchant.business_time || '营业时间未设置',
          email: merchant.Email || merchant.email || '',
          status: merchant.Status || merchant.status || 0,
          vehicles: 0 // 暂时设为0，后续可以通过其他API获取
        })).filter(merchant =>
          // 过滤掉没有坐标的商家
          merchant.longitude !== 0 && merchant.latitude !== 0
        );

        console.log('处理后的商家数据:', merchantsData);
        setMerchants(merchantsData);

        if (merchantsData.length === 0) {
          message.warning('暂无可用的网点数据或网点缺少坐标信息');
        } else {
          message.success(`成功加载 ${merchantsData.length} 个网点`);
        }
      } else {
        console.warn('API响应格式异常:', response);
        message.warning('获取网点数据格式异常');
        setMerchants([]);
      }
    } catch (error) {
      console.error('获取商家网点数据失败:', error);
      message.error('获取网点数据失败: ' + (error.message || '未知错误'));
      setMerchants([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMerchants();
  }, []);

  // 处理位置选择
  const handleLocationSelect = (location) => {
    console.log('选择的位置:', location);
    setSelectedLocation(location);
    message.success(`已选择位置: ${location.address}`);
  };

  // 计算到最近商家的距离
  const calculateNearestMerchant = async () => {
    if (!selectedLocation) {
      message.warning('请先选择一个位置');
      return;
    }

    try {
      // 计算到每个商家的距离
      const distances = [];
      for (const merchant of merchants) {
        const distance = locationService.calculateHaversineDistance(
          selectedLocation.latitude,
          selectedLocation.longitude,
          merchant.latitude,
          merchant.longitude
        );
        distances.push({
          merchant,
          distance,
          distanceStr: locationService.formatDistance(distance)
        });
      }

      // 按距离排序
      distances.sort((a, b) => a.distance - b.distance);
      
      const nearest = distances[0];
      message.success(`最近的网点是 ${nearest.merchant.name}，距离 ${nearest.distanceStr}`);
      
      console.log('距离排序结果:', distances);
    } catch (error) {
      console.error('计算距离失败:', error);
      message.error('计算距离失败');
    }
  };

  // 地图配置选项
  const mapConfigs = {
    basic: {
      center: [116.397428, 39.90923],
      zoom: 11,
      showMerchants: true,
      merchants: merchants
    },
    beijing: {
      center: [116.397428, 39.90923],
      zoom: 12,
      showMerchants: true,
      merchants: merchants.filter(m => m.address.includes('北京'))
    },
    shanghai: {
      center: [121.473701, 31.230416],
      zoom: 12,
      showMerchants: true,
      merchants: merchants.filter(m => m.address.includes('上海'))
    },
    shenzhen: {
      center: [114.057868, 22.543099],
      zoom: 12,
      showMerchants: true,
      merchants: merchants.filter(m => m.address.includes('深圳'))
    }
  };

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={3} style={{ margin: 0, color: '#fff', cursor: 'pointer' }} onClick={() => navigate('/')}>
              ZuLMe - 高德地图展示
            </Title>
          </Col>
          <Col>
            <Space>
              <Button type="primary" ghost onClick={() => navigate('/')}>
                返回首页
              </Button>
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          <Row gutter={24}>
            <Col span={18}>
              <InfoCard title="高德地图交互演示" extra={<EnvironmentOutlined />}>
                <Tabs activeKey={activeTab} onChange={setActiveTab}>
                  <TabPane tab="简化地图" key="simple">
                    <SimpleMapComponent
                      center={mapConfigs.basic.center}
                      zoom={mapConfigs.basic.zoom}
                      onLocationSelect={handleLocationSelect}
                      merchants={merchants}
                      height={600}
                    />
                  </TabPane>
                  <TabPane tab="全国视图" key="basic">
                    <AmapComponent
                      {...mapConfigs.basic}
                      onLocationSelect={handleLocationSelect}
                      merchants={merchants}
                      height={600}
                    />
                  </TabPane>
                </Tabs>
              </InfoCard>
            </Col>

            <Col span={6}>
              <InfoCard title="操作面板" extra={<UserOutlined />}>
                <Space direction="vertical" style={{ width: '100%' }} size="middle">
                  {selectedLocation ? (
                    <div>
                      <Title level={5}>📍 选择的位置</Title>
                      <Text strong>地址: </Text>
                      <Text>{selectedLocation.address}</Text>
                      <br />
                      <Text strong>经度: </Text>
                      <Text code>{selectedLocation.longitude?.toFixed(6)}</Text>
                      <br />
                      <Text strong>纬度: </Text>
                      <Text code>{selectedLocation.latitude?.toFixed(6)}</Text>
                      
                      <Divider />
                      <Space style={{ width: '100%' }}>
                        <Button
                          type="primary"
                          onClick={calculateNearestMerchant}
                          icon={<CarOutlined />}
                          style={{ flex: 1 }}
                        >
                          查找最近网点
                        </Button>
                        <Button
                          onClick={fetchMerchants}
                          loading={loading}
                          icon={<EnvironmentOutlined />}
                        >
                          刷新网点
                        </Button>
                      </Space>
                    </div>
                  ) : (
                    <div>
                      <Text type="secondary">请在地图上点击选择位置</Text>
                    </div>
                  )}

                  <Divider />
                  
                  <div>
                    <Collapse
                      size="small"
                      defaultActiveKey={['merchants']}
                      items={[
                        {
                          key: 'merchants',
                          label: (
                            <span>
                              🏪 网点信息
                              <span style={{
                                marginLeft: '8px',
                                padding: '2px 6px',
                                background: '#1890ff',
                                color: 'white',
                                borderRadius: '10px',
                                fontSize: '11px'
                              }}>
                                {merchants.length}
                              </span>
                            </span>
                          ),
                          children: loading ? (
                            <div style={{ textAlign: 'center', padding: '20px' }}>
                              <Text type="secondary">加载网点数据中...</Text>
                            </div>
                          ) : (
                            <div className="merchants-scroll" style={{ maxHeight: '280px', overflowY: 'auto', paddingRight: '4px' }}>
                              <Space direction="vertical" style={{ width: '100%' }} size="small">
                                {merchants.length > 0 ? (
                                  merchants.map(merchant => (
                                    <Card
                                      key={merchant.id}
                                      size="small"
                                      style={{
                                        marginBottom: '6px',
                                        border: '1px solid #e8e8e8',
                                        borderRadius: '6px',
                                        transition: 'all 0.2s'
                                      }}
                                      hoverable
                                    >
                                      <Space direction="vertical" size="small" style={{ width: '100%' }}>
                                        <Text strong style={{ fontSize: '13px', color: '#1890ff' }}>
                                          {merchant.name}
                                        </Text>
                                        <Text type="secondary" style={{ fontSize: '11px', lineHeight: '1.4' }}>
                                          📍 {merchant.address}
                                        </Text>
                                        {merchant.phone && merchant.phone !== '电话未设置' && (
                                          <Text style={{ fontSize: '11px', color: '#666' }}>
                                            📞 {merchant.phone}
                                          </Text>
                                        )}
                                        {merchant.businessTime && merchant.businessTime !== '营业时间未设置' && (
                                          <Text style={{ fontSize: '11px', color: '#666' }}>
                                            🕒 {merchant.businessTime}
                                          </Text>
                                        )}
                                        <div style={{
                                          display: 'flex',
                                          justifyContent: 'space-between',
                                          alignItems: 'center',
                                          marginTop: '4px',
                                          padding: '4px 0',
                                          borderTop: '1px solid #f0f0f0'
                                        }}>
                                          <Text style={{ fontSize: '10px', color: '#999' }}>
                                            🌍 {merchant.longitude.toFixed(4)}, {merchant.latitude.toFixed(4)}
                                          </Text>
                                          <Text style={{
                                            fontSize: '11px',
                                            background: '#f6ffed',
                                            color: '#52c41a',
                                            padding: '2px 6px',
                                            borderRadius: '4px'
                                          }}>
                                            <CarOutlined style={{ marginRight: 2 }} />
                                            {merchant.vehicles}辆
                                          </Text>
                                        </div>
                                      </Space>
                                    </Card>
                                  ))
                                ) : (
                                  <div style={{ textAlign: 'center', padding: '20px' }}>
                                    <Text type="secondary">暂无网点数据</Text>
                                    <br />
                                    <Button
                                      type="link"
                                      size="small"
                                      onClick={fetchMerchants}
                                      style={{ padding: 0 }}
                                    >
                                      重新加载
                                    </Button>
                                  </div>
                                )}
                              </Space>
                            </div>
                          )
                        }
                      ]}
                    />
                  </div>
                </Space>
              </InfoCard>

              <InfoCard title="功能说明" extra={<ShopOutlined />}>
                <Collapse
                  size="small"
                  ghost
                  items={[
                    {
                      key: 'features',
                      label: '🎯 交互功能',
                      children: (
                        <ul style={{
                          fontSize: '12px',
                          paddingLeft: '16px',
                          margin: 0,
                          lineHeight: '1.6'
                        }}>
                          <li style={{ marginBottom: '4px' }}>点击地图选择位置</li>
                          <li style={{ marginBottom: '4px' }}>搜索地点和地标</li>
                          <li style={{ marginBottom: '4px' }}>获取当前位置</li>
                          <li style={{ marginBottom: '4px' }}>查看商家网点</li>
                          <li>计算距离</li>
                        </ul>
                      )
                    },
                    {
                      key: 'controls',
                      label: '🗺️ 地图控件',
                      children: (
                        <ul style={{
                          fontSize: '12px',
                          paddingLeft: '16px',
                          margin: 0,
                          lineHeight: '1.6'
                        }}>
                          <li style={{ marginBottom: '4px' }}>缩放控制</li>
                          <li style={{ marginBottom: '4px' }}>比例尺</li>
                          <li style={{ marginBottom: '4px' }}>工具栏</li>
                          <li>定位按钮</li>
                        </ul>
                      )
                    },
                    {
                      key: 'markers',
                      label: '📍 标记说明',
                      children: (
                        <ul style={{
                          fontSize: '12px',
                          paddingLeft: '16px',
                          margin: 0,
                          lineHeight: '1.6'
                        }}>
                          <li style={{ marginBottom: '4px' }}>🔴 用户选择位置</li>
                          <li style={{ marginBottom: '4px' }}>🔵 商家网点</li>
                          <li>🟢 搜索结果</li>
                        </ul>
                      )
                    }
                  ]}
                />
              </InfoCard>
            </Col>
          </Row>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default MapDemo;
