import React, { useState } from 'react';
import { Layout, Typography, Card, Button, Space, message, Divider } from 'antd';
import { ApiOutlined, DatabaseOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import vehicleService from '../../services/vehicleService';

const { Header, Content } = Layout;
const { Title, Text } = Typography;

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
  max-width: 1200px;
  margin: 0 auto;
`;

const TestCard = styled(Card)`
  margin-bottom: 16px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
`;

const ApiTest = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [apiResponse, setApiResponse] = useState(null);

  const testMerchantListAPI = async () => {
    try {
      setLoading(true);
      console.log('测试商家列表API...');
      
      const response = await vehicleService.getLocationList({
        page: 1,
        pageSize: 100,
        status_filter: 1
      });

      console.log('API响应:', response);
      setApiResponse(response);
      
      if (response && response.data) {
        message.success('API调用成功');
      } else {
        message.warning('API响应格式异常');
      }
    } catch (error) {
      console.error('API调用失败:', error);
      message.error('API调用失败: ' + (error.message || '未知错误'));
      setApiResponse({ error: error.message || '未知错误' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <StyledLayout>
      <StyledHeader>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', height: '100%' }}>
          <Title level={3} style={{ margin: 0, color: '#fff', cursor: 'pointer' }} onClick={() => navigate('/')}>
            ZuLMe - API测试
          </Title>
          <Space>
            <Button type="primary" ghost onClick={() => navigate('/map-demo')}>
              地图演示
            </Button>
            <Button type="primary" ghost onClick={() => navigate('/')}>
              返回首页
            </Button>
          </Space>
        </div>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          <TestCard title="商家网点API测试" extra={<ApiOutlined />}>
            <Space direction="vertical" style={{ width: '100%' }} size="middle">
              <div>
                <Text strong>测试接口: </Text>
                <Text code>/admin/merchant/list</Text>
              </div>
              
              <div>
                <Text strong>请求参数: </Text>
                <pre style={{ background: '#f5f5f5', padding: '8px', borderRadius: '4px', fontSize: '12px' }}>
{JSON.stringify({
  page: 1,
  pageSize: 100,
  status_filter: 1
}, null, 2)}
                </pre>
              </div>

              <Button 
                type="primary" 
                onClick={testMerchantListAPI}
                loading={loading}
                icon={<DatabaseOutlined />}
              >
                测试商家列表API
              </Button>

              <Divider />

              {apiResponse && (
                <div>
                  <Text strong>API响应: </Text>
                  <pre style={{ 
                    background: '#f5f5f5', 
                    padding: '12px', 
                    borderRadius: '4px', 
                    fontSize: '12px',
                    maxHeight: '400px',
                    overflow: 'auto',
                    whiteSpace: 'pre-wrap'
                  }}>
                    {JSON.stringify(apiResponse, null, 2)}
                  </pre>
                </div>
              )}
            </Space>
          </TestCard>

          <TestCard title="数据分析">
            <Space direction="vertical" style={{ width: '100%' }}>
              {apiResponse && !apiResponse.error && (
                <>
                  <div>
                    <Text strong>响应状态: </Text>
                    <Text type={apiResponse.data ? 'success' : 'danger'}>
                      {apiResponse.data ? '✅ 成功' : '❌ 失败'}
                    </Text>
                  </div>
                  
                  {apiResponse.data && (
                    <>
                      <div>
                        <Text strong>商家数量: </Text>
                        <Text>{apiResponse.data.merchants ? apiResponse.data.merchants.length : 0}</Text>
                      </div>
                      
                      <div>
                        <Text strong>总数: </Text>
                        <Text>{apiResponse.data.total || 0}</Text>
                      </div>

                      {apiResponse.data.merchants && apiResponse.data.merchants.length > 0 && (
                        <div>
                          <Text strong>商家列表: </Text>
                          <div style={{ marginTop: '8px' }}>
                            {apiResponse.data.merchants.map((merchant, index) => (
                              <div key={index} style={{ 
                                background: '#fafafa', 
                                padding: '8px', 
                                margin: '4px 0', 
                                borderRadius: '4px',
                                fontSize: '12px'
                              }}>
                                <div><strong>ID:</strong> {merchant.ID || merchant.id}</div>
                                <div><strong>名称:</strong> {merchant.Name || merchant.name}</div>
                                <div><strong>地址:</strong> {merchant.Location || merchant.location}</div>
                                <div><strong>经度:</strong> {merchant.Longitude || merchant.longitude}</div>
                                <div><strong>纬度:</strong> {merchant.Latitude || merchant.latitude}</div>
                                <div><strong>状态:</strong> {merchant.Status || merchant.status}</div>
                              </div>
                            ))}
                          </div>
                        </div>
                      )}
                    </>
                  )}
                </>
              )}

              {apiResponse && apiResponse.error && (
                <div>
                  <Text strong>错误信息: </Text>
                  <Text type="danger">{apiResponse.error}</Text>
                </div>
              )}
            </Space>
          </TestCard>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default ApiTest;
