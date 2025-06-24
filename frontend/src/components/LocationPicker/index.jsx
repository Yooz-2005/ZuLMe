import React, { useState, useEffect } from 'react';
import { Input, Button, message, Spin, Card, Typography, Space, Tag, AutoComplete } from 'antd';
import { EnvironmentOutlined, SearchOutlined, AimOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import api from '../../services/api';

const { Text } = Typography;

// 样式组件
const LocationCard = styled(Card)`
  .ant-card-body {
    padding: 20px;
  }
`;

const AddressInput = styled(Input)`
  .ant-input {
    border-radius: 8px;
    border: 2px solid #e8e8e8;
    transition: all 0.3s ease;
    
    &:focus {
      border-color: #667eea;
      box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
    }
  }
`;

const DistanceTag = styled(Tag)`
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 14px;
  font-weight: 600;
  border: none;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
`;

const LocationPicker = ({ 
  onLocationChange, 
  onDistanceCalculated, 
  merchantId,
  placeholder = "请输入详细地址",
  showDistance = true 
}) => {
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const [coordinates, setCoordinates] = useState(null);
  const [distance, setDistance] = useState(null);
  const [suggestions, setSuggestions] = useState([]);

  // 地址建议列表（可以根据实际需求扩展）
  const addressSuggestions = [
    '北京市朝阳区三里屯',
    '上海市浦东新区陆家嘴',
    '深圳市南山区科技园',
    '广州市天河区珠江新城',
    '杭州市西湖区文三路',
    '成都市锦江区春熙路',
    '武汉市武昌区光谷',
    '南京市鼓楼区新街口'
  ];

  // 获取地址经纬度
  const getCoordinates = async (inputAddress) => {
    if (!inputAddress.trim()) {
      message.warning('请输入地址');
      return;
    }

    setLoading(true);
    try {
      const response = await api.post('/geocode/coordinates', {
        address: inputAddress
      });

      if (response.code === 200) {
        const coords = {
          longitude: response.longitude,
          latitude: response.latitude
        };
        setCoordinates(coords);
        
        // 通知父组件位置变化
        if (onLocationChange) {
          onLocationChange({
            address: inputAddress,
            coordinates: coords
          });
        }

        message.success('地址解析成功');
        
        // 如果需要计算距离且提供了商家ID
        if (showDistance && merchantId) {
          await calculateDistance(inputAddress);
        }
      } else {
        message.error(response.message || '地址解析失败');
      }
    } catch (error) {
      console.error('获取坐标失败:', error);
      message.error('地址解析失败，请检查网络连接');
    } finally {
      setLoading(false);
    }
  };

  // 计算到商家的距离
  const calculateDistance = async (inputAddress) => {
    if (!merchantId) {
      console.warn('未提供商家ID，无法计算距离');
      return;
    }

    try {
      const response = await api.post('/user/calculateDistance', {
        location: inputAddress,
        merchant_id: merchantId
      });

      if (response.code === 200) {
        setDistance(response.data.distance);
        
        // 通知父组件距离计算结果
        if (onDistanceCalculated) {
          onDistanceCalculated({
            distance: response.data.distance,
            distanceMeters: response.data.distance_meters
          });
        }
      } else {
        console.error('距离计算失败:', response.message);
      }
    } catch (error) {
      console.error('计算距离失败:', error);
    }
  };

  // 处理地址输入变化
  const handleAddressChange = (value) => {
    setAddress(value);
    
    // 简单的地址建议过滤
    if (value) {
      const filtered = addressSuggestions.filter(addr => 
        addr.toLowerCase().includes(value.toLowerCase())
      );
      setSuggestions(filtered.map(addr => ({ value: addr })));
    } else {
      setSuggestions([]);
    }
  };

  // 处理地址选择
  const handleAddressSelect = (value) => {
    setAddress(value);
    getCoordinates(value);
  };

  // 获取当前位置（如果浏览器支持）
  const getCurrentLocation = () => {
    if (!navigator.geolocation) {
      message.error('您的浏览器不支持地理定位');
      return;
    }

    setLoading(true);
    navigator.geolocation.getCurrentPosition(
      (position) => {
        const coords = {
          longitude: position.coords.longitude,
          latitude: position.coords.latitude
        };
        setCoordinates(coords);
        setAddress('当前位置');
        
        if (onLocationChange) {
          onLocationChange({
            address: '当前位置',
            coordinates: coords
          });
        }

        message.success('获取当前位置成功');
        setLoading(false);

        // 如果需要计算距离
        if (showDistance && merchantId) {
          // 注意：这里需要逆地理编码将坐标转换为地址
          // 暂时使用坐标字符串
          const coordsString = `${coords.longitude},${coords.latitude}`;
          calculateDistance(coordsString);
        }
      },
      (error) => {
        console.error('获取位置失败:', error);
        message.error('获取当前位置失败');
        setLoading(false);
      },
      {
        enableHighAccuracy: true,
        timeout: 10000,
        maximumAge: 300000
      }
    );
  };

  return (
    <LocationCard>
      <Space direction="vertical" style={{ width: '100%' }} size="middle">
        <div>
          <Text strong style={{ marginBottom: 8, display: 'block' }}>
            <EnvironmentOutlined style={{ marginRight: 8, color: '#667eea' }} />
            选择地址
          </Text>
          
          <Space.Compact style={{ width: '100%' }}>
            <AutoComplete
              style={{ flex: 1 }}
              value={address}
              options={suggestions}
              onSelect={handleAddressSelect}
              onChange={handleAddressChange}
              placeholder={placeholder}
            >
              <AddressInput
                prefix={<SearchOutlined style={{ color: '#999' }} />}
                size="large"
              />
            </AutoComplete>
            
            <Button
              type="primary"
              size="large"
              loading={loading}
              onClick={() => getCoordinates(address)}
              style={{
                background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                border: 'none'
              }}
            >
              解析
            </Button>
            
            <Button
              size="large"
              icon={<AimOutlined />}
              onClick={getCurrentLocation}
              loading={loading}
              title="获取当前位置"
            />
          </Space.Compact>
        </div>

        {/* 显示坐标信息 */}
        {coordinates && (
          <div>
            <Text type="secondary" style={{ fontSize: 12 }}>
              经纬度: {coordinates.longitude.toFixed(6)}, {coordinates.latitude.toFixed(6)}
            </Text>
          </div>
        )}

        {/* 显示距离信息 */}
        {showDistance && distance && (
          <div>
            <Text strong>距离最近网点: </Text>
            <DistanceTag>{distance}</DistanceTag>
          </div>
        )}

        {loading && (
          <div style={{ textAlign: 'center', padding: '20px 0' }}>
            <Spin />
            <Text type="secondary" style={{ marginLeft: 8 }}>
              正在处理...
            </Text>
          </div>
        )}
      </Space>
    </LocationCard>
  );
};

export default LocationPicker;
