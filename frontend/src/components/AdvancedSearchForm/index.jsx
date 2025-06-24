import React, { useState, useEffect } from 'react';
import { 
  Input, 
  DatePicker, 
  Button, 
  Row, 
  Col, 
  Select, 
  Typography, 
  Slider, 
  InputNumber,
  Collapse,
  Space
} from 'antd';
import { SearchOutlined, FilterOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import dayjs from 'dayjs';
import vehicleService from '../../services/vehicleService';
import BrandSelector from '../BrandSelector';

const { Title } = Typography;
const { RangePicker } = DatePicker;
const { Option } = Select;
const { Panel } = Collapse;

const SearchContainer = styled.div`
  background: white;
  padding: 30px;
  border-radius: 8px;
  width: 80%;
  max-width: 1000px;
  margin: 0 auto;
  box-shadow: 0 4px 16px rgba(0,0,0,0.1);
`;

const PriceRangeContainer = styled.div`
  .ant-slider {
    margin: 10px 0;
  }
  .price-inputs {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;
  }
`;

const AdvancedSearchForm = ({ 
  initialValues = {}, 
  onSearch, 
  showTitle = true,
  layout = 'horizontal',
  showAdvanced = false
}) => {
  const [searchParams, setSearchParams] = useState({
    location: '',
    dates: null,
    carType: undefined,
    brandId: undefined,
    priceMin: 0,
    priceMax: 10000,
    yearMin: 2015,
    yearMax: new Date().getFullYear(),
    ...initialValues
  });
  const [vehicleTypes, setVehicleTypes] = useState([]);
  const [loading, setLoading] = useState(false);
  const [priceRange, setPriceRange] = useState([0, 10000]);
  const [yearRange, setYearRange] = useState([2015, new Date().getFullYear()]);
  const navigate = useNavigate();

  // 获取车辆类型列表
  useEffect(() => {
    const fetchVehicleTypes = async () => {
      try {
        const response = await vehicleService.getVehicleTypes();
        console.log('Vehicle types response:', response);

        if (response && response.code === 200 && response.data && Array.isArray(response.data.vehicle_types)) {
          setVehicleTypes(response.data.vehicle_types);
        } else if (response && response.code === 200 && response.data && Array.isArray(response.data)) {
          setVehicleTypes(response.data);
        } else {
          setVehicleTypes([
            { id: 1, name: '经济型' },
            { id: 2, name: '舒适型' },
            { id: 3, name: '豪华型' },
            { id: 4, name: 'SUV' }
          ]);
        }
      } catch (error) {
        console.error('获取车辆类型失败:', error);
        setVehicleTypes([
          { id: 1, name: '经济型' },
          { id: 2, name: '舒适型' },
          { id: 3, name: '豪华型' },
          { id: 4, name: 'SUV' }
        ]);
      }
    };

    fetchVehicleTypes();
  }, []);

  const handleSearch = async () => {
    setLoading(true);
    try {
      const searchData = {
        ...searchParams,
        priceMin: priceRange[0],
        priceMax: priceRange[1],
        yearMin: yearRange[0],
        yearMax: yearRange[1]
      };

      if (onSearch) {
        await onSearch(searchData);
      } else {
        const queryParams = new URLSearchParams();
        
        if (searchData.location) {
          queryParams.append('location', searchData.location);
        }
        if (searchData.dates && searchData.dates.length === 2) {
          queryParams.append('start_date', searchData.dates[0].format('YYYY-MM-DD'));
          queryParams.append('end_date', searchData.dates[1].format('YYYY-MM-DD'));
        }
        if (searchData.carType) {
          queryParams.append('vehicle_type', searchData.carType);
        }
        if (searchData.brandId) {
          queryParams.append('brand_id', searchData.brandId);
        }
        if (searchData.priceMin > 0) {
          queryParams.append('price_min', searchData.priceMin);
        }
        if (searchData.priceMax < 10000) {
          queryParams.append('price_max', searchData.priceMax);
        }
        if (searchData.yearMin > 2015) {
          queryParams.append('year_min', searchData.yearMin);
        }
        if (searchData.yearMax < new Date().getFullYear()) {
          queryParams.append('year_max', searchData.yearMax);
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

  const disabledDate = (current) => {
    return current && current < dayjs().startOf('day');
  };

  const handlePriceChange = (value) => {
    setPriceRange(value);
  };

  const handleYearChange = (value) => {
    setYearRange(value);
  };

  const renderBasicSearch = () => (
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
  );

  const renderAdvancedFilters = () => (
    <Collapse ghost>
      <Panel header={<><FilterOutlined /> 高级筛选</>} key="1">
        <Row gutter={[16, 16]}>
          <Col span={12}>
            <Typography.Text strong>价格范围 (元/天)</Typography.Text>
            <PriceRangeContainer>
              <Slider
                range
                min={0}
                max={10000}
                step={100}
                value={priceRange}
                onChange={handlePriceChange}
                marks={{
                  0: '0',
                  2500: '2500',
                  5000: '5000',
                  7500: '7500',
                  10000: '10000+'
                }}
              />
              <div className="price-inputs">
                <InputNumber
                  min={0}
                  max={10000}
                  value={priceRange[0]}
                  onChange={value => setPriceRange([value, priceRange[1]])}
                  formatter={value => `¥ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                  parser={value => value.replace(/¥\s?|(,*)/g, '')}
                />
                <span style={{ margin: '0 8px' }}>-</span>
                <InputNumber
                  min={0}
                  max={10000}
                  value={priceRange[1]}
                  onChange={value => setPriceRange([priceRange[0], value])}
                  formatter={value => `¥ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                  parser={value => value.replace(/¥\s?|(,*)/g, '')}
                />
              </div>
            </PriceRangeContainer>
          </Col>
          <Col span={12}>
            <Typography.Text strong>车辆年份</Typography.Text>
            <PriceRangeContainer>
              <Slider
                range
                min={2015}
                max={new Date().getFullYear()}
                value={yearRange}
                onChange={handleYearChange}
                marks={{
                  2015: '2015',
                  2018: '2018',
                  2021: '2021',
                  [new Date().getFullYear()]: new Date().getFullYear().toString()
                }}
              />
              <div className="price-inputs">
                <InputNumber
                  min={2015}
                  max={new Date().getFullYear()}
                  value={yearRange[0]}
                  onChange={value => setYearRange([value, yearRange[1]])}
                />
                <span style={{ margin: '0 8px' }}>-</span>
                <InputNumber
                  min={2015}
                  max={new Date().getFullYear()}
                  value={yearRange[1]}
                  onChange={value => setYearRange([yearRange[0], value])}
                />
              </div>
            </PriceRangeContainer>
          </Col>
        </Row>
      </Panel>
    </Collapse>
  );

  return (
    <SearchContainer>
      {showTitle && (
        <Title level={2} style={{ textAlign: 'center', marginBottom: 24 }}>
          找到您的理想座驾
        </Title>
      )}
      <Space direction="vertical" style={{ width: '100%' }} size="large">
        {renderBasicSearch()}
        {showAdvanced && renderAdvancedFilters()}
      </Space>
    </SearchContainer>
  );
};

export default AdvancedSearchForm;
