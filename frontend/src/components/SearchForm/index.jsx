import React, { useState, useEffect } from 'react';
import { Input, DatePicker, Button, Row, Col, Select, Typography } from 'antd';
import { SearchOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import dayjs from 'dayjs';
import vehicleService from '../../services/vehicleService';
import BrandSelector from '../BrandSelector';

const { Title } = Typography;
const { RangePicker } = DatePicker;
const { Option } = Select;

const SearchContainer = styled.div`
  background: white;
  padding: 30px;
  border-radius: 8px;
  width: 80%;
  max-width: 1000px;
  margin: 0 auto;
  box-shadow: 0 4px 16px rgba(0,0,0,0.1);
`;

const SearchForm = ({ 
  initialValues = {}, 
  onSearch, 
  showTitle = true,
  layout = 'horizontal' 
}) => {
  const [searchParams, setSearchParams] = useState({
    location: '',
    dates: null,
    carType: undefined,
    brandId: undefined,
    ...initialValues
  });
  const [vehicleTypes, setVehicleTypes] = useState([]);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // 获取车辆类型列表
  useEffect(() => {
    const fetchVehicleTypes = async () => {
      try {
        const response = await vehicleService.getVehicleTypes();
        console.log('Vehicle types response:', response);

        // 确保 vehicleTypes 始终是数组
        if (response && response.code === 200 && response.data && Array.isArray(response.data.vehicle_types)) {
          setVehicleTypes(response.data.vehicle_types);
        } else if (response && response.code === 200 && response.data && Array.isArray(response.data)) {
          setVehicleTypes(response.data);
        } else {
          // 使用默认的车辆类型
          setVehicleTypes([
            { id: 'economy', name: '经济型' },
            { id: 'comfort', name: '舒适型' },
            { id: 'luxury', name: '豪华型' },
            { id: 'suv', name: 'SUV' }
          ]);
        }
      } catch (error) {
        console.error('获取车辆类型失败:', error);
        // 使用默认的车辆类型
        setVehicleTypes([
          { id: 'economy', name: '经济型' },
          { id: 'comfort', name: '舒适型' },
          { id: 'luxury', name: '豪华型' },
          { id: 'suv', name: 'SUV' }
        ]);
      }
    };

    fetchVehicleTypes();
  }, []);

  const handleSearch = async () => {
    setLoading(true);
    try {
      if (onSearch) {
        // 如果传入了onSearch回调，直接调用
        await onSearch(searchParams);
      } else {
        // 否则跳转到搜索结果页面
        const queryParams = new URLSearchParams();
        
        if (searchParams.location) {
          queryParams.append('location', searchParams.location);
        }
        if (searchParams.dates && searchParams.dates.length === 2) {
          queryParams.append('start_date', searchParams.dates[0].format('YYYY-MM-DD'));
          queryParams.append('end_date', searchParams.dates[1].format('YYYY-MM-DD'));
        }
        if (searchParams.carType) {
          queryParams.append('vehicle_type', searchParams.carType);
        }
        if (searchParams.brandId) {
          queryParams.append('brand_id', searchParams.brandId);
        }

        navigate(`/search?${queryParams.toString()}`);
      }
    } catch (error) {
      console.error('搜索失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const updateSearchParams = (key, value) => {
    setSearchParams(prev => ({
      ...prev,
      [key]: value
    }));
  };

  // 禁用当日之前的日期
  const disabledDate = (current) => {
    // 禁用今天之前的日期
    return current && current < dayjs().startOf('day');
  };

  if (layout === 'vertical') {
    return (
      <SearchContainer>
        {showTitle && (
          <Title level={2} style={{ textAlign: 'center', marginBottom: 24 }}>
            找到您的理想座驾
          </Title>
        )}
        <Row gutter={[16, 16]}>
          <Col span={24}>
            <Input
              placeholder="取车地点"
              prefix={<SearchOutlined />}
              value={searchParams.location}
              onChange={e => updateSearchParams('location', e.target.value)}
              size="large"
            />
          </Col>
          <Col span={24}>
            <RangePicker
              style={{ width: '100%' }}
              placeholder={['取车日期', '还车日期']}
              value={searchParams.dates}
              onChange={dates => updateSearchParams('dates', dates)}
              disabledDate={disabledDate}
              size="large"
            />
          </Col>
          <Col span={24}>
            <Select
              style={{ width: '100%' }}
              placeholder="选择车型"
              value={searchParams.carType}
              onChange={value => updateSearchParams('carType', value)}
              size="large"
              allowClear
            >
              {Array.isArray(vehicleTypes) && vehicleTypes.map(type => (
                <Option key={type.id} value={type.id}>
                  {type.name}
                </Option>
              ))}
            </Select>
          </Col>
          <Col span={24}>
            <BrandSelector
              value={searchParams.brandId}
              onChange={value => updateSearchParams('brandId', value)}
              placeholder="选择车辆品牌"
              size="large"
              showHotBrands={true}
            />
          </Col>
          <Col span={24}>
            <Button
              type="primary"
              block
              size="large"
              loading={loading}
              onClick={handleSearch}
            >
              搜索车辆
            </Button>
          </Col>
        </Row>
      </SearchContainer>
    );
  }

  return (
    <SearchContainer>
      {showTitle && (
        <Title level={2} style={{ textAlign: 'center', marginBottom: 24 }}>
          找到您的理想座驾
        </Title>
      )}
      <Row gutter={[16, 16]}>
        <Col span={24}>
          <Row gutter={16}>
            <Col span={6}>
              <Input
                placeholder="取车地点"
                prefix={<SearchOutlined />}
                value={searchParams.location}
                onChange={e => updateSearchParams('location', e.target.value)}
                size="large"
              />
            </Col>
            <Col span={6}>
              <RangePicker
                style={{ width: '100%' }}
                placeholder={['取车日期', '还车日期']}
                value={searchParams.dates}
                onChange={dates => updateSearchParams('dates', dates)}
                disabledDate={disabledDate}
                size="large"
              />
            </Col>
            <Col span={4}>
              <Select
                style={{ width: '100%' }}
                placeholder="车型"
                value={searchParams.carType}
                onChange={value => updateSearchParams('carType', value)}
                allowClear
                size="large"
              >
                {Array.isArray(vehicleTypes) && vehicleTypes.map(type => (
                  <Option key={type.id} value={type.id}>
                    {type.name}
                  </Option>
                ))}
              </Select>
            </Col>
            <Col span={5}>
              <BrandSelector
                value={searchParams.brandId}
                onChange={value => updateSearchParams('brandId', value)}
                placeholder="品牌"
                size="large"
                showHotBrands={false}
              />
            </Col>
            <Col span={3}>
              <Button
                type="primary"
                block
                loading={loading}
                onClick={handleSearch}
                size="large"
              >
                搜索
              </Button>
            </Col>
          </Row>
        </Col>
      </Row>
    </SearchContainer>
  );
};

export default SearchForm;
