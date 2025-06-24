import React, { useEffect, useState } from 'react';
import { Layout, Typography, Card, Button, Space, message, Alert } from 'antd';
import { CheckCircleOutlined, CloseCircleOutlined, EnvironmentOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';

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

const MapTest = () => {
  const navigate = useNavigate();
  const [apiStatus, setApiStatus] = useState({
    loaded: false,
    version: null,
    plugins: [],
    error: null
  });

  useEffect(() => {
    // 检查高德地图API状态
    const checkApiStatus = () => {
      try {
        if (typeof window.AMap !== 'undefined') {
          console.log('高德地图API已加载');
          console.log('AMap对象:', window.AMap);
          
          // 获取版本信息
          const version = window.AMap.version || 'Unknown';
          
          // 检查可用的插件
          const availablePlugins = [];
          const pluginsToCheck = [
            'Map', 'Marker', 'InfoWindow', 'Scale', 'ToolBar', 'ControlBar',
            'Geolocation', 'PlaceSearch', 'Autocomplete', 'Geocoder', 'Driving', 'Circle'
          ];
          
          pluginsToCheck.forEach(plugin => {
            if (window.AMap[plugin]) {
              availablePlugins.push(plugin);
            }
          });

          setApiStatus({
            loaded: true,
            version: version,
            plugins: availablePlugins,
            error: null
          });

          message.success('高德地图API检查完成');
        } else {
          console.error('高德地图API未加载');
          setApiStatus({
            loaded: false,
            version: null,
            plugins: [],
            error: 'API未加载'
          });
        }
      } catch (error) {
        console.error('检查API状态时出错:', error);
        setApiStatus({
          loaded: false,
          version: null,
          plugins: [],
          error: error.message
        });
      }
    };

    // 延迟检查，确保API有时间加载
    const timer = setTimeout(checkApiStatus, 1000);
    
    return () => clearTimeout(timer);
  }, []);

  const testBasicMap = () => {
    if (!apiStatus.loaded) {
      message.error('高德地图API未加载，无法创建地图');
      return;
    }

    try {
      // 创建一个临时的地图容器
      const tempDiv = document.createElement('div');
      tempDiv.style.width = '100px';
      tempDiv.style.height = '100px';
      tempDiv.style.position = 'absolute';
      tempDiv.style.top = '-1000px';
      document.body.appendChild(tempDiv);

      // 尝试创建地图实例
      const testMap = new window.AMap.Map(tempDiv, {
        center: [116.397428, 39.90923],
        zoom: 13
      });

      message.success('地图实例创建成功！');
      
      // 清理
      setTimeout(() => {
        testMap.destroy();
        document.body.removeChild(tempDiv);
      }, 1000);

    } catch (error) {
      console.error('创建地图实例失败:', error);
      message.error('创建地图实例失败: ' + error.message);
    }
  };

  return (
    <StyledLayout>
      <StyledHeader>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', height: '100%' }}>
          <Title level={3} style={{ margin: 0, color: '#fff', cursor: 'pointer' }} onClick={() => navigate('/')}>
            ZuLMe - 地图API测试
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
          <TestCard title="高德地图API状态检查" extra={<EnvironmentOutlined />}>
            <Space direction="vertical" style={{ width: '100%' }} size="middle">
              <div>
                <Text strong>API加载状态: </Text>
                {apiStatus.loaded ? (
                  <Text type="success">
                    <CheckCircleOutlined style={{ marginRight: 4 }} />
                    已加载
                  </Text>
                ) : (
                  <Text type="danger">
                    <CloseCircleOutlined style={{ marginRight: 4 }} />
                    未加载
                  </Text>
                )}
              </div>

              {apiStatus.version && (
                <div>
                  <Text strong>API版本: </Text>
                  <Text code>{apiStatus.version}</Text>
                </div>
              )}

              {apiStatus.error && (
                <Alert
                  message="错误信息"
                  description={apiStatus.error}
                  type="error"
                  showIcon
                />
              )}

              <div>
                <Text strong>可用插件 ({apiStatus.plugins.length}): </Text>
                <div style={{ marginTop: 8 }}>
                  {apiStatus.plugins.map(plugin => (
                    <Text key={plugin} code style={{ marginRight: 8, marginBottom: 4, display: 'inline-block' }}>
                      {plugin}
                    </Text>
                  ))}
                </div>
              </div>

              <div>
                <Button 
                  type="primary" 
                  onClick={testBasicMap}
                  disabled={!apiStatus.loaded}
                >
                  测试创建地图实例
                </Button>
              </div>
            </Space>
          </TestCard>

          <TestCard title="调试信息">
            <Space direction="vertical" style={{ width: '100%' }}>
              <div>
                <Text strong>window.AMap: </Text>
                <Text code>{typeof window.AMap}</Text>
              </div>
              <div>
                <Text strong>User Agent: </Text>
                <Text style={{ fontSize: '12px', wordBreak: 'break-all' }}>
                  {navigator.userAgent}
                </Text>
              </div>
              <div>
                <Text strong>当前URL: </Text>
                <Text code>{window.location.href}</Text>
              </div>
            </Space>
          </TestCard>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default MapTest;
