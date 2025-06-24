import React, { useEffect, useRef, useState } from 'react';
import { Card, Button, Space, message, Input, List, Typography, Tag } from 'antd';
import { EnvironmentOutlined, SearchOutlined, AimOutlined, CarOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import '../../types/amap.d.ts';

const { Text } = Typography;

// 样式组件
const MapContainer = styled.div`
  width: 100%;
  height: 500px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e8e8e8;
`;

const ControlPanel = styled(Card)`
  margin-bottom: 16px;
  
  .ant-card-body {
    padding: 16px;
  }
`;

const SearchInput = styled(Input)`
  .ant-input {
    border-radius: 6px;
  }
`;

const AmapComponent = ({ 
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
  const [searchKeyword, setSearchKeyword] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [markers, setMarkers] = useState([]);

  // 初始化地图
  useEffect(() => {
    // 检查高德地图API是否加载
    const initMap = () => {
      console.log('开始初始化地图...');
      console.log('window.AMap:', window.AMap);

      if (!window.AMap) {
        console.error('高德地图API未加载');
        message.error('高德地图API未加载，请检查网络连接');
        return;
      }

      if (!mapRef.current) {
        console.error('地图容器未找到');
        return;
      }

      try {
        console.log('创建地图实例，中心点:', center, '缩放级别:', zoom);

        // 创建地图实例
        const mapInstance = new window.AMap.Map(mapRef.current, {
          center: center,
          zoom: zoom,
          mapStyle: 'amap://styles/normal', // 地图样式
          features: ['bg', 'road', 'building', 'point'], // 显示要素
          viewMode: '2D' // 2D模式
        });

        console.log('地图实例创建成功:', mapInstance);

        // 等待地图加载完成后添加控件
        mapInstance.on('complete', () => {
          try {
            // 添加地图控件
            if (window.AMap.Scale) {
              mapInstance.addControl(new window.AMap.Scale()); // 比例尺
            }
            if (window.AMap.ToolBar) {
              mapInstance.addControl(new window.AMap.ToolBar()); // 工具条
            }
            if (window.AMap.ControlBar) {
              mapInstance.addControl(new window.AMap.ControlBar()); // 控制条
            }
          } catch (error) {
            console.warn('添加地图控件失败:', error);
          }
        });

        // 地图点击事件
        mapInstance.on('click', (e) => {
          const { lng, lat } = e.lnglat;
          console.log('地图点击坐标:', lng, lat);

          // 逆地理编码获取地址
          try {
            const geocoder = new window.AMap.Geocoder();
            geocoder.getAddress([lng, lat], (status, result) => {
              if (status === 'complete' && result.regeocode) {
                const address = result.regeocode.formattedAddress;
                setCurrentLocation({
                  longitude: lng,
                  latitude: lat,
                  address: address
                });

                // 添加标记
                addMarker([lng, lat], address, 'user');

                // 通知父组件
                if (onLocationSelect) {
                  onLocationSelect({
                    longitude: lng,
                    latitude: lat,
                    address: address
                  });
                }

                message.success(`已选择位置: ${address}`);
              }
            });
          } catch (error) {
            console.error('逆地理编码失败:', error);
            message.error('获取地址信息失败');
          }
        });

        setMap(mapInstance);
      } catch (error) {
        console.error('地图初始化失败:', error);
        message.error('地图初始化失败，请刷新页面重试');
      }
    };

    // 如果高德地图API还未加载，等待加载
    if (window.AMap) {
      initMap();
    } else {
      // 等待API加载
      const checkAMap = setInterval(() => {
        if (window.AMap) {
          clearInterval(checkAMap);
          initMap();
        }
      }, 100);

      // 10秒后停止检查
      setTimeout(() => {
        clearInterval(checkAMap);
        if (!window.AMap) {
          message.error('高德地图API加载超时，请检查网络连接');
        }
      }, 10000);
    }

    // 清理函数
    return () => {
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

  // 添加商家标记
  useEffect(() => {
    if (map && showMerchants && merchants.length > 0 && window.AMap) {
      // 延迟添加标记，确保地图完全初始化
      const timer = setTimeout(() => {
        addMerchantMarkers();
      }, 100);

      return () => clearTimeout(timer);
    }
  }, [map, showMerchants, merchants]);

  // 添加标记点
  const addMarker = (position, title, type = 'default') => {
    if (!map || !window.AMap) {
      console.warn('地图未初始化或高德地图API未加载');
      return;
    }

    try {
      // 清除之前的用户标记
      if (type === 'user') {
        markers.forEach(marker => {
          try {
            if (marker.getExtData && marker.getExtData().type === 'user') {
              map.remove(marker);
            }
          } catch (error) {
            console.warn('移除用户标记失败:', error);
          }
        });
      }

      // 创建标记
      const marker = new window.AMap.Marker({
        position: position,
        title: title,
        icon: getMarkerIcon(type),
        anchor: 'bottom-center'
      });

      // 设置扩展数据
      marker.setExtData({ type: type, title: title });

      // 添加到地图
      if (map && typeof map.add === 'function') {
        map.add(marker);
      } else {
        console.error('地图对象无效或add方法不存在');
        return;
      }

      // 创建信息窗体
      const infoWindow = new window.AMap.InfoWindow({
        content: `<div style="padding: 10px;">
          <h4>${title}</h4>
          <p>经度: ${position[0].toFixed(6)}</p>
          <p>纬度: ${position[1].toFixed(6)}</p>
        </div>`,
        anchor: 'bottom-center',
        offset: new window.AMap.Pixel(0, -30)
      });

      // 点击标记显示信息窗体
      marker.on('click', () => {
        infoWindow.open(map, position);
      });

      // 更新标记列表
      setMarkers(prev => [...prev.filter(m => m.getExtData && m.getExtData().type !== type || type !== 'user'), marker]);

      return marker;
    } catch (error) {
      console.error('添加标记失败:', error);
      return null;
    }
  };

  // 获取标记图标
  const getMarkerIcon = (type) => {
    const iconMap = {
      user: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_r.png',
      merchant: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_b.png',
      default: 'https://webapi.amap.com/theme/v1.3/markers/n/mark_g.png'
    };
    return iconMap[type] || iconMap.default;
  };

  // 添加商家标记
  const addMerchantMarkers = () => {
    if (!map || !window.AMap) {
      console.warn('地图未初始化，无法添加商家标记');
      return;
    }

    merchants.forEach(merchant => {
      if (merchant.longitude && merchant.latitude) {
        try {
          addMarker(
            [merchant.longitude, merchant.latitude],
            merchant.name || `商家${merchant.id}`,
            'merchant'
          );
        } catch (error) {
          console.error('添加商家标记失败:', error);
        }
      }
    });
  };

  // 获取当前位置
  const getCurrentLocation = () => {
    if (!map) return;

    const geolocation = new window.AMap.Geolocation({
      enableHighAccuracy: true,
      timeout: 10000,
      maximumAge: 0,
      convert: true,
      showButton: true,
      buttonPosition: 'LB',
      showMarker: true,
      showCircle: true,
      panToLocation: true,
      zoomToAccuracy: true
    });

    geolocation.getCurrentPosition((status, result) => {
      if (status === 'complete') {
        const { lng, lat } = result.position;
        const address = result.formattedAddress;
        
        setCurrentLocation({
          longitude: lng,
          latitude: lat,
          address: address
        });

        // 移动地图中心到当前位置
        map.setCenter([lng, lat]);
        map.setZoom(15);

        message.success('定位成功');
      } else {
        message.error('定位失败: ' + result.message);
      }
    });
  };

  // 地址搜索
  const searchLocation = () => {
    if (!map || !searchKeyword.trim()) {
      message.warning('请输入搜索关键词');
      return;
    }

    const placeSearch = new window.AMap.PlaceSearch({
      pageSize: 10,
      pageIndex: 1,
      city: '全国'
    });

    placeSearch.search(searchKeyword, (status, result) => {
      if (status === 'complete' && result.poiList && result.poiList.pois.length > 0) {
        const pois = result.poiList.pois;
        setSearchResults(pois);
        
        // 清除之前的搜索标记
        markers.forEach(marker => {
          try {
            if (marker.getExtData().type === 'search') {
              map.remove(marker);
            }
          } catch (error) {
            console.warn('移除搜索标记失败:', error);
          }
        });

        // 添加搜索结果标记
        pois.forEach((poi, index) => {
          if (index < 5) { // 只显示前5个结果
            const position = [poi.location.lng, poi.location.lat];
            addMarker(position, poi.name, 'search');
          }
        });

        // 调整地图视野
        if (pois.length > 0) {
          const firstPoi = pois[0];
          map.setCenter([firstPoi.location.lng, firstPoi.location.lat]);
          map.setZoom(14);
        }

        message.success(`找到 ${pois.length} 个相关位置`);
      } else {
        message.error('未找到相关位置');
        setSearchResults([]);
      }
    });
  };

  // 选择搜索结果
  const selectSearchResult = (poi) => {
    const position = [poi.location.lng, poi.location.lat];
    map.setCenter(position);
    map.setZoom(16);

    setCurrentLocation({
      longitude: poi.location.lng,
      latitude: poi.location.lat,
      address: poi.name
    });

    // 添加用户选择标记
    addMarker(position, poi.name, 'user');

    if (onLocationSelect) {
      onLocationSelect({
        longitude: poi.location.lng,
        latitude: poi.location.lat,
        address: poi.name
      });
    }

    message.success(`已选择: ${poi.name}`);
  };

  // 路径规划
  const planRoute = (startPoint, endPoint) => {
    if (!map) return;

    // 创建驾车路线规划实例
    const driving = new window.AMap.Driving({
      map: map,
      showTraffic: true,
      hideMarkers: false
    });

    // 根据起终点经纬度规划驾车导航路线
    driving.search(startPoint, endPoint, (status, result) => {
      if (status === 'complete') {
        message.success('路线规划成功');
        console.log('路线信息:', result);
      } else {
        message.error('路线规划失败');
      }
    });
  };

  // 添加圆形覆盖物
  const addCircle = (center, radius, options = {}) => {
    if (!map) return;

    const circle = new window.AMap.Circle({
      center: center,
      radius: radius,
      strokeColor: options.strokeColor || '#FF33FF',
      strokeWeight: options.strokeWeight || 6,
      strokeOpacity: options.strokeOpacity || 0.2,
      fillOpacity: options.fillOpacity || 0.4,
      fillColor: options.fillColor || '#1791fc',
      zIndex: 50
    });

    map.add(circle);
    return circle;
  };

  return (
    <div>
      <ControlPanel size="small">
        <Space direction="vertical" style={{ width: '100%' }}>
          <Space.Compact style={{ width: '100%' }}>
            <SearchInput
              placeholder="搜索地点、商圈、地标等"
              value={searchKeyword}
              onChange={(e) => setSearchKeyword(e.target.value)}
              onPressEnter={searchLocation}
              prefix={<SearchOutlined />}
            />
            <Button type="primary" onClick={searchLocation}>
              搜索
            </Button>
            <Button icon={<AimOutlined />} onClick={getCurrentLocation}>
              定位
            </Button>
          </Space.Compact>

          {currentLocation && (
            <div>
              <Text strong>当前选择: </Text>
              <Tag color="blue">{currentLocation.address}</Tag>
            </div>
          )}
        </Space>
      </ControlPanel>

      <MapContainer 
        ref={mapRef} 
        style={{ height: `${height}px` }}
      />

      {searchResults.length > 0 && (
        <Card 
          title="搜索结果" 
          size="small" 
          style={{ marginTop: 16, maxHeight: 200, overflow: 'auto' }}
        >
          <List
            size="small"
            dataSource={searchResults.slice(0, 5)}
            renderItem={(poi) => (
              <List.Item
                actions={[
                  <Button 
                    type="link" 
                    size="small" 
                    onClick={() => selectSearchResult(poi)}
                  >
                    选择
                  </Button>
                ]}
              >
                <List.Item.Meta
                  avatar={<EnvironmentOutlined />}
                  title={poi.name}
                  description={poi.address}
                />
              </List.Item>
            )}
          />
        </Card>
      )}
    </div>
  );
};

export default AmapComponent;
