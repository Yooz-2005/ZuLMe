import React, { useState, useEffect } from 'react';
import { Select, Space, Tag, Avatar, Spin } from 'antd';
import { CarOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import vehicleService from '../../services/vehicleService';

const { Option } = Select;

const StyledSelect = styled(Select)`
  .ant-select-selector {
    border-radius: 8px;
    border: 2px solid #e2e8f0;
    transition: all 0.3s ease;
    
    &:hover {
      border-color: #667eea;
    }
  }
  
  &.ant-select-focused .ant-select-selector {
    border-color: #667eea;
    box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
  }
`;

const BrandOption = styled.div`
  display: flex;
  align-items: center;
  padding: 8px 0;
  
  .brand-info {
    margin-left: 12px;
    
    .brand-name {
      font-weight: 600;
      color: #1e293b;
      margin-bottom: 2px;
    }
    
    .brand-english {
      font-size: 12px;
      color: #64748b;
    }
  }
`;

const HotBrands = styled.div`
  margin-bottom: 16px;
  
  .hot-title {
    font-size: 14px;
    font-weight: 600;
    color: #1e293b;
    margin-bottom: 8px;
  }
  
  .hot-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }
`;

const HotTag = styled(Tag)`
  cursor: pointer;
  border-radius: 16px;
  padding: 4px 12px;
  border: 2px solid #e2e8f0;
  background: white;
  color: #64748b;
  transition: all 0.3s ease;
  
  &:hover {
    border-color: #667eea;
    color: #667eea;
    transform: translateY(-1px);
  }
  
  &.selected {
    border-color: #667eea;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
  }
`;

const BrandSelector = ({ 
  value, 
  onChange, 
  placeholder = "请选择车辆品牌",
  showHotBrands = true,
  allowClear = true,
  size = "large",
  ...props 
}) => {
  const [brands, setBrands] = useState([]);
  const [hotBrands, setHotBrands] = useState([]);
  const [loading, setLoading] = useState(false);
  const [searchValue, setSearchValue] = useState('');

  // 获取品牌列表
  const fetchBrands = async (keyword = '') => {
    setLoading(true);
    try {
      const response = await vehicleService.getBrandList({
        page: 1,
        page_size: 100,
        keyword,
        status: 1
      });
      
      if (response.code === 200) {
        setBrands(response.data.vehicle_brands || []);
      }
    } catch (error) {
      console.error('获取品牌列表失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 获取热门品牌
  const fetchHotBrands = async () => {
    try {
      const response = await vehicleService.getBrandList({
        page: 1,
        page_size: 20,
        is_hot: 1,
        status: 1
      });
      
      if (response.code === 200) {
        setHotBrands(response.data.vehicle_brands || []);
      }
    } catch (error) {
      console.error('获取热门品牌失败:', error);
    }
  };

  useEffect(() => {
    fetchBrands();
    if (showHotBrands) {
      fetchHotBrands();
    }
  }, [showHotBrands]);

  // 搜索品牌
  const handleSearch = (value) => {
    setSearchValue(value);
    if (value) {
      fetchBrands(value);
    } else {
      fetchBrands();
    }
  };

  // 选择热门品牌
  const handleHotBrandClick = (brandId) => {
    onChange?.(brandId);
  };

  // 渲染品牌选项
  const renderBrandOption = (brand) => (
    <BrandOption>
      <Avatar 
        size={32}
        src={brand.logo}
        icon={<CarOutlined />}
        style={{
          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
        }}
      />
      <div className="brand-info">
        <div className="brand-name">{brand.name}</div>
        {brand.english_name && (
          <div className="brand-english">{brand.english_name}</div>
        )}
      </div>
    </BrandOption>
  );

  return (
    <div>
      {showHotBrands && hotBrands.length > 0 && (
        <HotBrands>
          <div className="hot-title">🔥 热门品牌</div>
          <div className="hot-tags">
            {hotBrands.map(brand => (
              <HotTag
                key={brand.id}
                className={value === brand.id ? 'selected' : ''}
                onClick={() => handleHotBrandClick(brand.id)}
              >
                {brand.name}
              </HotTag>
            ))}
          </div>
        </HotBrands>
      )}
      
      <StyledSelect
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        allowClear={allowClear}
        size={size}
        showSearch
        filterOption={false}
        onSearch={handleSearch}
        loading={loading}
        notFoundContent={loading ? <Spin size="small" /> : '暂无品牌'}
        {...props}
      >
        {brands.map(brand => (
          <Option key={brand.id} value={brand.id}>
            {renderBrandOption(brand)}
          </Option>
        ))}
      </StyledSelect>
    </div>
  );
};

export default BrandSelector;
