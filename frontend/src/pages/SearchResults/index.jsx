import React, { useState, useEffect } from 'react';
import { 
  Layout, 
  Row, 
  Col, 
  Typography, 
  Spin, 
  Alert, 
  Pagination,
  Button,
  Space,
  Breadcrumb,
  Empty
} from 'antd';
import { HomeOutlined, SearchOutlined, CarOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate, useLocation } from 'react-router-dom';
import styled from 'styled-components';
import dayjs from 'dayjs';
import VehicleCard from '../../components/VehicleCard';
import SearchForm from '../../components/SearchForm';
import vehicleService from '../../services/vehicleService';
import { PAGINATION_CONFIG } from '../../utils/constants';

const { Header, Content } = Layout;
const { Title } = Typography;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
`;

const StyledHeader = styled(Header)`
  background: #000;
  padding: 0 50px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  position: fixed;
  width: 100%;
  z-index: 1;
  top: 0;
  left: 0;
`;

const ContentWrapper = styled.div`
  margin-top: 64px;
  padding: 24px 50px;
`;

const SearchSection = styled.div`
  background: #f0f2f5;
  padding: 30px 0;
  margin-bottom: 30px;
`;

const SearchResults = () => {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchParams, setSearchParams] = useState({});
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: PAGINATION_CONFIG.DEFAULT_PAGE_SIZE,
    total: 0
  });
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const navigate = useNavigate();
  const location = useLocation();

  // 检查用户登录状态
  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsLoggedIn(!!token);
  }, []);

  // 解析URL参数
  const parseUrlParams = () => {
    const urlParams = new URLSearchParams(location.search);
    const params = {};

    if (urlParams.get('location')) {
      params.location = urlParams.get('location');
    }

    if (urlParams.get('start_date') && urlParams.get('end_date')) {
      params.dates = [
        dayjs(urlParams.get('start_date')),
        dayjs(urlParams.get('end_date'))
      ];
    }

    if (urlParams.get('vehicle_type')) {
      params.carType = parseInt(urlParams.get('vehicle_type'));
    }

    if (urlParams.get('brand_id')) {
      params.brandId = parseInt(urlParams.get('brand_id'));
    }

    return params;
  };

  // 搜索车辆
  const searchVehicles = async (params = {}, page = 1, pageSize = PAGINATION_CONFIG.DEFAULT_PAGE_SIZE) => {
    setLoading(true);
    setError(null);
    
    try {
      const searchData = {
        ...params,
        page,
        pageSize
      };
      
      const response = await vehicleService.searchVehicles(searchData);
      
      if (response && response.code === 200 && response.data) {
        setVehicles(response.data.vehicles || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.total || 0
        });
      }
    } catch (err) {
      setError('搜索失败，请稍后重试');
      console.error('搜索失败:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const params = parseUrlParams();
    setSearchParams(params);
    searchVehicles(params);
  }, [location.search]);

  const handlePageChange = (page, pageSize) => {
    searchVehicles(searchParams, page, pageSize);
  };

  const handleSearch = async (newSearchParams) => {
    setSearchParams(newSearchParams);
    await searchVehicles(newSearchParams, 1, pagination.pageSize);

    // 更新URL参数
    const queryParams = new URLSearchParams();
    if (newSearchParams.location) {
      queryParams.append('location', newSearchParams.location);
    }
    if (newSearchParams.dates && newSearchParams.dates.length === 2) {
      queryParams.append('start_date', newSearchParams.dates[0].format('YYYY-MM-DD'));
      queryParams.append('end_date', newSearchParams.dates[1].format('YYYY-MM-DD'));
    }
    if (newSearchParams.carType) {
      queryParams.append('vehicle_type', newSearchParams.carType);
    }
    if (newSearchParams.brandId) {
      queryParams.append('brand_id', newSearchParams.brandId);
    }

    navigate(`/search?${queryParams.toString()}`, { replace: true });
  };

  const getSearchSummary = () => {
    const parts = [];
    if (searchParams.location) {
      parts.push(`地点: ${searchParams.location}`);
    }
    if (searchParams.dates && searchParams.dates.length === 2) {
      parts.push(`时间: ${searchParams.dates[0].format('MM-DD')} 至 ${searchParams.dates[1].format('MM-DD')}`);
    }
    if (searchParams.carType) {
      parts.push(`车型: ${searchParams.carType}`);
    }
    return parts.length > 0 ? parts.join(' | ') : '所有车辆';
  };

  return (
    <StyledLayout>
      <StyledHeader>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={3} style={{ margin: 0, cursor: 'pointer', color: '#fff' }} onClick={() => navigate('/')}>
              ZuLMe
            </Title>
          </Col>
          <Col>
            <Space>
              <Button onClick={() => navigate('/vehicles')}>所有车辆</Button>
              <Button onClick={() => navigate('/')}>返回首页</Button>
              {isLoggedIn && (
                <Button
                  icon={<UserOutlined />}
                  onClick={() => navigate('/personal-center')}
                >
                  我的
                </Button>
              )}
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          <Breadcrumb style={{ marginBottom: 16 }}>
            <Breadcrumb.Item href="/">
              <HomeOutlined />
            </Breadcrumb.Item>
            <Breadcrumb.Item>
              <SearchOutlined />
              <span>搜索结果</span>
            </Breadcrumb.Item>
          </Breadcrumb>

          <SearchSection>
            <SearchForm 
              initialValues={searchParams}
              onSearch={handleSearch}
              showTitle={false}
              layout="horizontal"
            />
          </SearchSection>

          <div style={{ marginBottom: 24 }}>
            <Title level={2}>
              搜索结果
              {!loading && (
                <span style={{ fontSize: '16px', fontWeight: 'normal', color: '#666' }}>
                  （{getSearchSummary()}，共 {pagination.total} 辆）
                </span>
              )}
            </Title>
          </div>

          {error && (
            <Alert
              message="错误"
              description={error}
              type="error"
              showIcon
              style={{ marginBottom: 24 }}
              action={
                <Button size="small" onClick={() => searchVehicles(searchParams)}>
                  重试
                </Button>
              }
            />
          )}

          <Spin spinning={loading}>
            {vehicles.length > 0 ? (
              <>
                <Row gutter={[24, 24]}>
                  {vehicles.map(vehicle => (
                    <Col xs={24} sm={12} md={8} lg={6} key={vehicle.id}>
                      <VehicleCard vehicle={vehicle} />
                    </Col>
                  ))}
                </Row>

                <div style={{ textAlign: 'center', marginTop: 40 }}>
                  <Pagination
                    current={pagination.current}
                    pageSize={pagination.pageSize}
                    total={pagination.total}
                    onChange={handlePageChange}
                    showSizeChanger
                    showQuickJumper
                    showTotal={(total, range) => 
                      `第 ${range[0]}-${range[1]} 条，共 ${total} 条`
                    }
                  />
                </div>
              </>
            ) : (
              !loading && (
                <Empty
                  image={Empty.PRESENTED_IMAGE_SIMPLE}
                  description={
                    <span>
                      没有找到符合条件的车辆
                      <br />
                      <Button type="link" onClick={() => navigate('/vehicles')}>
                        查看所有车辆
                      </Button>
                    </span>
                  }
                />
              )
            )}
          </Spin>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default SearchResults;
