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
  Breadcrumb
} from 'antd';
import { HomeOutlined, CarOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
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

const VehicleList = () => {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: PAGINATION_CONFIG.DEFAULT_PAGE_SIZE,
    total: 0
  });
  
  const navigate = useNavigate();

  // 获取车辆列表
  const fetchVehicles = async (page = 1, pageSize = PAGINATION_CONFIG.DEFAULT_PAGE_SIZE) => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await vehicleService.getVehicleList({
        page,
        page_size: pageSize
      });
      
      if (response && response.data) {
        setVehicles(response.data.vehicles || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.total || 0
        });
      }
    } catch (err) {
      setError('获取车辆列表失败，请稍后重试');
      console.error('获取车辆列表失败:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchVehicles();
  }, []);

  const handlePageChange = (page, pageSize) => {
    fetchVehicles(page, pageSize);
  };

  const handleSearch = async (searchParams) => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await vehicleService.searchVehicles({
        ...searchParams,
        page: 1,
        pageSize: pagination.pageSize
      });
      
      if (response && response.data) {
        setVehicles(response.data.vehicles || []);
        setPagination({
          current: 1,
          pageSize: pagination.pageSize,
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
              <Button onClick={() => navigate('/')}>返回首页</Button>
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
              <CarOutlined />
              <span>车辆列表</span>
            </Breadcrumb.Item>
          </Breadcrumb>

          <SearchSection>
            <SearchForm 
              onSearch={handleSearch}
              showTitle={false}
              layout="horizontal"
            />
          </SearchSection>

          <div style={{ marginBottom: 24 }}>
            <Title level={2}>
              所有车辆 
              {!loading && (
                <span style={{ fontSize: '16px', fontWeight: 'normal', color: '#666' }}>
                  （共 {pagination.total} 辆）
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
                <Button size="small" onClick={() => fetchVehicles()}>
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
                <div style={{ textAlign: 'center', padding: '60px 0' }}>
                  <CarOutlined style={{ fontSize: '64px', color: '#ccc' }} />
                  <Title level={4} style={{ color: '#999', marginTop: 16 }}>
                    暂无车辆数据
                  </Title>
                </div>
              )
            )}
          </Spin>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default VehicleList;
