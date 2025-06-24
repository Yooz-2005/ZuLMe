import React, { useEffect, useRef, useState } from 'react';
import { Card, Button, Space, message, Typography } from 'antd';
import { EnvironmentOutlined, AimOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Text } = Typography;

// 样式组件
const MapContainer = styled.div`
  width: 100%;
  height: 500px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e8e8e8;
  background: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const ControlPanel = styled(Card)`
  margin-bottom: 16px;
  
  .ant-card-body {
    padding: 16px;
  }
`;

const SimpleMapComponent = ({ 
  center = [116.397428, 39.90923], // 默认中心点（北京天安门）
  zoom = 13,
  onLocationSelect,
  showMerchants = true,
  merchants = [],
  height = 500
}) => {
  const mapRef = useRef(null);
  const [map, setMap] = useState(null);
  const [currentLocation, setCurrentLocation] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [markers, setMarkers] = useState([]);

  // 添加商家标记
  const addMerchantMarkers = (mapInstance) => {
    if (!mapInstance || !merchants || merchants.length === 0) return;
    if (!window.AMap) {
      console.error('高德地图API未加载');
      return;
    }

    console.log('添加商家标记:', merchants);

    const newMarkers = [];
    merchants.forEach(merchant => {
      if (merchant.longitude && merchant.latitude) {
        try {
          const marker = new window.AMap.Marker({
            position: [merchant.longitude, merchant.latitude],
            title: merchant.name || `商家${merchant.id}`,
            icon: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_b.png',
            anchor: 'bottom-center'
          });

          // 创建信息窗体
          const infoWindow = new window.AMap.InfoWindow({
            content: `<div style="padding: 10px;">
              <h4>${merchant.name}</h4>
              <p>📍 ${merchant.address}</p>
              ${merchant.phone && merchant.phone !== '电话未设置' ? `<p>📞 ${merchant.phone}</p>` : ''}
              ${merchant.businessTime && merchant.businessTime !== '营业时间未设置' ? `<p>🕒 ${merchant.businessTime}</p>` : ''}
              <p>🌍 ${merchant.longitude.toFixed(6)}, ${merchant.latitude.toFixed(6)}</p>
            </div>`,
            anchor: 'bottom-center',
            offset: new window.AMap.Pixel(0, -30)
          });

          // 点击标记显示信息窗体
          marker.on('click', () => {
            infoWindow.open(mapInstance, [merchant.longitude, merchant.latitude]);
          });

          // 确保地图实例有效且有add方法
          if (mapInstance && typeof mapInstance.add === 'function') {
            mapInstance.add(marker);
            newMarkers.push(marker);
          } else {
            console.error('地图实例无效或add方法不存在');
          }
        } catch (error) {
          console.error('添加标记失败:', error);
        }
      }
    });

    // 批量更新标记状态
    setMarkers(prev => [...prev, ...newMarkers]);

    if (newMarkers.length > 0) {
      console.log(`成功添加 ${newMarkers.length} 个商家标记`);
    }
  };

  // 初始化地图
  useEffect(() => {
    const initMap = () => {
      console.log('开始初始化简化地图...');
      
      // 检查高德地图API
      if (typeof window.AMap === 'undefined') {
        console.error('高德地图API未加载');
        message.error('高德地图API未加载，请刷新页面重试');
        setIsLoading(false);
        return;
      }

      if (!mapRef.current) {
        console.error('地图容器未找到');
        setIsLoading(false);
        return;
      }

      try {
        console.log('创建地图实例...');
        
        // 创建最简单的地图实例
        const mapInstance = new window.AMap.Map(mapRef.current, {
          center: center,
          zoom: zoom,
          viewMode: '2D'
        });

        console.log('地图实例创建成功');

        // 地图加载完成事件
        mapInstance.on('complete', () => {
          console.log('地图加载完成');
          setIsLoading(false);
          message.success('地图加载成功');
          // 地图加载完成后添加商家标记
          addMerchantMarkers(mapInstance);
        });

        // 地图点击事件
        mapInstance.on('click', (e) => {
          console.log('地图点击事件:', e);
          const { lng, lat } = e.lnglat;
          
          setCurrentLocation({
            longitude: lng,
            latitude: lat,
            address: `经度: ${lng.toFixed(6)}, 纬度: ${lat.toFixed(6)}`
          });

          if (onLocationSelect) {
            onLocationSelect({
              longitude: lng,
              latitude: lat,
              address: `经度: ${lng.toFixed(6)}, 纬度: ${lat.toFixed(6)}`
            });
          }

          message.success(`已选择位置: ${lng.toFixed(6)}, ${lat.toFixed(6)}`);
        });

        setMap(mapInstance);

      } catch (error) {
        console.error('地图初始化失败:', error);
        message.error('地图初始化失败: ' + error.message);
        setIsLoading(false);
      }
    };

    // 延迟初始化，确保DOM已渲染
    const timer = setTimeout(() => {
      initMap();
    }, 100);

    return () => {
      clearTimeout(timer);
      // 清理标记
      markers.forEach(marker => {
        try {
          if (map) {
            map.remove(marker);
          }
        } catch (error) {
          console.warn('清理标记失败:', error);
        }
      });
      // 销毁地图
      if (map) {
        try {
          map.destroy();
        } catch (error) {
          console.warn('地图销毁失败:', error);
        }
      }
    };
  }, [center, zoom, onLocationSelect]);

  // 监听merchants数据变化，更新标记
  useEffect(() => {
    if (map && merchants && merchants.length > 0) {
      // 清除之前的商家标记
      markers.forEach(marker => {
        try {
          map.remove(marker);
        } catch (error) {
          console.warn('移除标记失败:', error);
        }
      });
      setMarkers([]);

      // 添加新的商家标记
      addMerchantMarkers(map);
    }
  }, [merchants, map]);

  // 获取当前位置
  const getCurrentLocation = () => {
    if (!map) {
      message.warning('地图未初始化');
      return;
    }

    if (!window.AMap.Geolocation) {
      message.error('定位功能不可用');
      return;
    }

    try {
      const geolocation = new window.AMap.Geolocation({
        enableHighAccuracy: true,
        timeout: 10000
      });

      geolocation.getCurrentPosition((status, result) => {
        if (status === 'complete') {
          const { lng, lat } = result.position;
          
          setCurrentLocation({
            longitude: lng,
            latitude: lat,
            address: result.formattedAddress || `经度: ${lng.toFixed(6)}, 纬度: ${lat.toFixed(6)}`
          });

          map.setCenter([lng, lat]);
          map.setZoom(15);

          message.success('定位成功');
        } else {
          message.error('定位失败: ' + (result.message || '未知错误'));
        }
      });
    } catch (error) {
      console.error('定位失败:', error);
      message.error('定位功能异常');
    }
  };

  return (
    <div>
      <ControlPanel size="small">
        <Space direction="vertical" style={{ width: '100%' }}>
          <Space>
            <Button 
              icon={<AimOutlined />} 
              onClick={getCurrentLocation}
              disabled={!map}
            >
              获取当前位置
            </Button>
            <Text type="secondary">
              {isLoading ? '地图加载中...' : '点击地图选择位置'}
            </Text>
          </Space>

          {currentLocation && (
            <div>
              <Text strong>当前选择: </Text>
              <Text>{currentLocation.address}</Text>
            </div>
          )}
        </Space>
      </ControlPanel>

      <MapContainer 
        ref={mapRef} 
        style={{ height: `${height}px` }}
      >
        {isLoading && (
          <div style={{ textAlign: 'center' }}>
            <EnvironmentOutlined style={{ fontSize: '48px', color: '#ccc', marginBottom: '16px' }} />
            <br />
            <Text type="secondary">地图加载中...</Text>
          </div>
        )}
      </MapContainer>
    </div>
  );
};

export default SimpleMapComponent;
