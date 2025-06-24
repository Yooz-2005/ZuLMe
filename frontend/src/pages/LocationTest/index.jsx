import React, { useState } from 'react';
import { Layout, Typography, Card, Row, Col, Space, Divider, message } from 'antd';
import { EnvironmentOutlined, CarOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import LocationPicker from '../../components/LocationPicker';
import locationService from '../../services/locationService';

const { Header, Content } = Layout;
const { Title, Text, Paragraph } = Typography;

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
  padding: 32px 50px;
  max-width: 1200px;
  margin: 0 auto;
`;

const DemoCard = styled(Card)`
  margin-bottom: 24px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
`;

const ResultCard = styled(Card)`
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%);
  border: 1px solid #e2e8f0;
  border-radius: 12px;
`;

const LocationTest = () => {
  const navigate = useNavigate();
  const [locationData, setLocationData] = useState(null);
  const [distanceData, setDistanceData] = useState(null);

  // 处理位置变化
  const handleLocationChange = (data) => {
    console.log('位置变化:', data);
    setLocationData(data);
    message.success(`地址解析成功: ${data.address}`);
  };

  // 处理距离计算结果
  const handleDistanceCalculated = (data) => {
    console.log('距离计算结果:', data);
    setDistanceData(data);
    message.success(`距离计算成功: ${data.distance}`);
  };

  // 测试不同的地址
  const testAddresses = [
    '北京市朝阳区三里屯',
    '上海市浦东新区陆家嘴',
    '深圳市南山区科技园',
    '广州市天河区珠江新城'
  ];

  // 批量测试地址解析
  const batchTestAddresses = async () => {
    message.info('开始批量测试地址解析...');
    
    for (const address of testAddresses) {
      try {
        const result = await locationService.getCoordinatesByAddress(address);
        if (result.success) {
          console.log(`${address} -> 经度: ${result.data.longitude}, 纬度: ${result.data.latitude}`);
        } else {
          console.error(`${address} 解析失败: ${result.message}`);
        }
      } catch (error) {
        console.error(`${address} 解析异常:`, error);
      }
    }
    
    message.success('批量测试完成，请查看控制台输出');
  };

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={3} style={{ margin: 0, color: '#fff', cursor: 'pointer' }} onClick={() => navigate('/')}>
              ZuLMe - 地理位置测试
            </Title>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          <Row gutter={24}>
            <Col span={16}>
              <DemoCard title="地址输入与距离计算" extra={<EnvironmentOutlined />}>
                <Space direction="vertical" style={{ width: '100%' }} size="large">
                  <div>
                    <Paragraph>
                      这个组件演示了如何使用高德地图API进行地址解析和距离计算：
                    </Paragraph>
                    <ul>
                      <li>输入详细地址，自动解析为经纬度坐标</li>
                      <li>计算用户位置到指定商家的距离</li>
                      <li>支持获取浏览器当前位置</li>
                      <li>提供地址输入建议</li>
                    </ul>
                  </div>

                  <LocationPicker
                    onLocationChange={handleLocationChange}
                    onDistanceCalculated={handleDistanceCalculated}
                    merchantId={1} // 测试商家ID
                    placeholder="请输入您的详细地址"
                    showDistance={true}
                  />
                </Space>
              </DemoCard>

              <DemoCard title="API测试功能">
                <Space wrap>
                  <button 
                    onClick={batchTestAddresses}
                    style={{
                      padding: '8px 16px',
                      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                      color: 'white',
                      border: 'none',
                      borderRadius: '6px',
                      cursor: 'pointer'
                    }}
                  >
                    批量测试地址解析
                  </button>
                </Space>
              </DemoCard>
            </Col>

            <Col span={8}>
              <ResultCard title="解析结果" extra={<CarOutlined />}>
                <Space direction="vertical" style={{ width: '100%' }} size="middle">
                  {locationData ? (
                    <div>
                      <Title level={5}>📍 位置信息</Title>
                      <Text strong>地址: </Text>
                      <Text>{locationData.address}</Text>
                      <br />
                      <Text strong>经度: </Text>
                      <Text code>{locationData.coordinates?.longitude?.toFixed(6)}</Text>
                      <br />
                      <Text strong>纬度: </Text>
                      <Text code>{locationData.coordinates?.latitude?.toFixed(6)}</Text>
                    </div>
                  ) : (
                    <Text type="secondary">请输入地址进行解析</Text>
                  )}

                  {distanceData && (
                    <>
                      <Divider />
                      <div>
                        <Title level={5}>📏 距离信息</Title>
                        <Text strong>到最近网点: </Text>
                        <Text style={{ 
                          color: '#667eea', 
                          fontWeight: 'bold',
                          fontSize: '16px'
                        }}>
                          {distanceData.distance}
                        </Text>
                      </div>
                    </>
                  )}

                  <Divider />
                  <div>
                    <Title level={5}>🔧 技术说明</Title>
                    <ul style={{ fontSize: '12px', color: '#666' }}>
                      <li>使用高德地图Web服务API</li>
                      <li>Redis Geo计算距离</li>
                      <li>支持浏览器地理定位</li>
                      <li>实时坐标验证</li>
                    </ul>
                  </div>
                </Space>
              </ResultCard>
            </Col>
          </Row>

          <DemoCard title="高德地图SDK集成架构">
            <Row gutter={16}>
              <Col span={6}>
                <Card size="small" style={{ textAlign: 'center' }}>
                  <Title level={5}>前端输入</Title>
                  <Text type="secondary">用户地址输入</Text>
                </Card>
              </Col>
              <Col span={6}>
                <Card size="small" style={{ textAlign: 'center' }}>
                  <Title level={5}>API网关</Title>
                  <Text type="secondary">请求转发</Text>
                </Card>
              </Col>
              <Col span={6}>
                <Card size="small" style={{ textAlign: 'center' }}>
                  <Title level={5}>高德API</Title>
                  <Text type="secondary">地址→坐标</Text>
                </Card>
              </Col>
              <Col span={6}>
                <Card size="small" style={{ textAlign: 'center' }}>
                  <Title level={5}>Redis Geo</Title>
                  <Text type="secondary">距离计算</Text>
                </Card>
              </Col>
            </Row>
          </DemoCard>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default LocationTest;
