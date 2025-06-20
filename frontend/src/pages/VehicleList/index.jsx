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
  Card,
  Select,
  Input,
  Badge
} from 'antd';
import {
  HomeOutlined,
  CarOutlined,
  DashboardOutlined,
  SearchOutlined,
  FilterOutlined,
  AppstoreOutlined,
  BarsOutlined,
  SortAscendingOutlined,
  UserOutlined
} from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import VehicleCard from '../../components/VehicleCard';
import SearchForm from '../../components/SearchForm';
import vehicleService from '../../services/vehicleService';
import { PAGINATION_CONFIG } from '../../utils/constants';

const { Header, Content } = Layout;
const { Title, Text } = Typography;
const { Option } = Select;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
`;

const StyledHeader = styled(Header)`
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 0 50px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.15);
  position: fixed;
  width: 100%;
  z-index: 1000;
  top: 0;
  left: 0;
  backdrop-filter: blur(10px);
`;

const ContentWrapper = styled.div`
  margin-top: 64px;
  padding: 32px 50px;
  max-width: 1400px;
  margin-left: auto;
  margin-right: auto;
`;

const SearchSection = styled.div`
  background: white;
  padding: 32px;
  margin-bottom: 32px;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
  border: 1px solid #e2e8f0;
`;

const FilterSection = styled.div`
  background: white;
  padding: 24px 32px;
  margin-bottom: 24px;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  border: 1px solid #e2e8f0;
`;

const VehicleGrid = styled.div`
  background: white;
  padding: 32px;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.12);
  border: 1px solid #e2e8f0;
  margin-bottom: 32px;
`;

const StatsCard = styled(Card)`
  border-radius: 16px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);

  .ant-card-body {
    padding: 20px;
  }
`;

const ViewToggle = styled.div`
  display: flex;
  background: #f1f5f9;
  border-radius: 12px;
  padding: 4px;

  .toggle-btn {
    padding: 8px 16px;
    border-radius: 8px;
    border: none;
    background: transparent;
    color: #64748b;
    cursor: pointer;
    transition: all 0.3s ease;

    &.active {
      background: white;
      color: #667eea;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }

    &:hover:not(.active) {
      color: #667eea;
    }
  }
`;

const StyledNavButton = styled(Button)`
  background: transparent;
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #fff;
  font-size: 16px;
  font-weight: 500;
  padding: 8px 24px;
  height: 40px;
  border-radius: 6px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;

  &:before {
    content: '';
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.2), transparent);
    transition: left 0.5s;
  }

  &:hover {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
    border: 1px solid #667eea !important;
    color: #fff !important;
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);

    &:before {
      left: 100%;
    }
  }

  &:active {
    transform: translateY(0);
    box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
  }
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
  const [viewMode, setViewMode] = useState('grid'); // 'grid' or 'list'
  const [sortBy, setSortBy] = useState('default');
  const [filterStatus, setFilterStatus] = useState('all');

  const navigate = useNavigate();

  // 检查用户是否已登录
  const userPhone = localStorage.getItem('userPhone');
  const token = localStorage.getItem('token');
  const isLoggedIn = userPhone && token;

  // 获取车辆列表
  const fetchVehicles = async (page = 1, pageSize = PAGINATION_CONFIG.DEFAULT_PAGE_SIZE) => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await vehicleService.getVehicleList({
        page,
        page_size: pageSize
      });
      
      if (response && response.code === 200 && response.data) {
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
      
      if (response && response.code === 200 && response.data) {
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
              {isLoggedIn && (
                <>
                  <StyledNavButton
                    icon={<DashboardOutlined />}
                    onClick={() => navigate('/dashboard')}
                  >
                    返回Dashboard
                  </StyledNavButton>
                  <StyledNavButton
                    icon={<UserOutlined />}
                    onClick={() => navigate('/personal-center')}
                  >
                    我的
                  </StyledNavButton>
                </>
              )}
              <StyledNavButton onClick={() => navigate('/')}>
                <HomeOutlined style={{ marginRight: 8 }} />
                返回首页
              </StyledNavButton>
            </Space>
          </Col>
        </Row>
      </StyledHeader>

      <Content>
        <ContentWrapper>
          {/* 面包屑导航 */}
          <div style={{
            background: 'white',
            padding: '16px 24px',
            borderRadius: 16,
            marginBottom: 24,
            boxShadow: '0 2px 8px rgba(0,0,0,0.06)',
            border: '1px solid #e2e8f0'
          }}>
            <Breadcrumb
              separator=">"
              style={{
                fontSize: 16,
                fontWeight: 500
              }}
            >
              <Breadcrumb.Item href="/">
                <HomeOutlined style={{ color: '#667eea', marginRight: 4 }} />
                <span style={{ color: '#667eea' }}>首页</span>
              </Breadcrumb.Item>
              <Breadcrumb.Item>
                <CarOutlined style={{ color: '#64748b', marginRight: 4 }} />
                <span style={{ color: '#64748b' }}>车辆列表</span>
              </Breadcrumb.Item>
            </Breadcrumb>
          </div>

          {/* 搜索区域 */}
          <SearchSection>
            <div style={{ marginBottom: 16 }}>
              <Title level={3} style={{
                margin: 0,
                color: '#1e293b',
                fontWeight: 600
              }}>
                <SearchOutlined style={{ marginRight: 12, color: '#667eea' }} />
                搜索车辆
              </Title>
            </div>
            <SearchForm
              onSearch={handleSearch}
              showTitle={false}
              layout="horizontal"
            />
          </SearchSection>

          {/* 统计信息和筛选 */}
          <Row gutter={24} style={{ marginBottom: 24 }}>
            <Col span={6}>
              <StatsCard>
                <div style={{ textAlign: 'center' }}>
                  <div style={{
                    fontSize: 32,
                    fontWeight: 700,
                    background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                    WebkitBackgroundClip: 'text',
                    WebkitTextFillColor: 'transparent',
                    backgroundClip: 'text'
                  }}>
                    {!loading ? pagination.total : '--'}
                  </div>
                  <Text type="secondary">总车辆数</Text>
                </div>
              </StatsCard>
            </Col>
            <Col span={6}>
              <StatsCard>
                <div style={{ textAlign: 'center' }}>
                  <div style={{
                    fontSize: 32,
                    fontWeight: 700,
                    color: '#10b981'
                  }}>
                    {!loading ? vehicles.filter(v => v.status === 'available').length : '--'}
                  </div>
                  <Text type="secondary">可租车辆</Text>
                </div>
              </StatsCard>
            </Col>
            <Col span={6}>
              <StatsCard>
                <div style={{ textAlign: 'center' }}>
                  <div style={{
                    fontSize: 32,
                    fontWeight: 700,
                    color: '#ef4444'
                  }}>
                    {!loading ? vehicles.filter(v => v.status === 'rented').length : '--'}
                  </div>
                  <Text type="secondary">已租车辆</Text>
                </div>
              </StatsCard>
            </Col>
            <Col span={6}>
              <StatsCard>
                <div style={{ textAlign: 'center' }}>
                  <div style={{
                    fontSize: 32,
                    fontWeight: 700,
                    color: '#f59e0b'
                  }}>
                    {!loading ? vehicles.filter(v => v.status === 'maintenance').length : '--'}
                  </div>
                  <Text type="secondary">维护中</Text>
                </div>
              </StatsCard>
            </Col>
          </Row>

          {/* 筛选和视图切换 */}
          <FilterSection>
            <Row justify="space-between" align="middle">
              <Col>
                <Space size="large">
                  <div>
                    <Text style={{ marginRight: 8, fontWeight: 500 }}>状态筛选:</Text>
                    <Select
                      value={filterStatus}
                      onChange={setFilterStatus}
                      style={{ width: 120 }}
                    >
                      <Option value="all">全部</Option>
                      <Option value="available">可租</Option>
                      <Option value="rented">已租</Option>
                      <Option value="maintenance">维护中</Option>
                    </Select>
                  </div>
                  <div>
                    <Text style={{ marginRight: 8, fontWeight: 500 }}>排序:</Text>
                    <Select
                      value={sortBy}
                      onChange={setSortBy}
                      style={{ width: 140 }}
                      suffixIcon={<SortAscendingOutlined />}
                    >
                      <Option value="default">默认排序</Option>
                      <Option value="price_asc">价格从低到高</Option>
                      <Option value="price_desc">价格从高到低</Option>
                      <Option value="year_desc">年份从新到旧</Option>
                      <Option value="year_asc">年份从旧到新</Option>
                    </Select>
                  </div>
                </Space>
              </Col>
              <Col>
                <Space>
                  <Text style={{ fontWeight: 500 }}>视图:</Text>
                  <ViewToggle>
                    <button
                      className={`toggle-btn ${viewMode === 'grid' ? 'active' : ''}`}
                      onClick={() => setViewMode('grid')}
                    >
                      <AppstoreOutlined /> 网格
                    </button>
                    <button
                      className={`toggle-btn ${viewMode === 'list' ? 'active' : ''}`}
                      onClick={() => setViewMode('list')}
                    >
                      <BarsOutlined /> 列表
                    </button>
                  </ViewToggle>
                </Space>
              </Col>
            </Row>
          </FilterSection>

          {/* 错误提示 */}
          {error && (
            <Alert
              message="错误"
              description={error}
              type="error"
              showIcon
              style={{
                marginBottom: 24,
                borderRadius: 12,
                border: '1px solid #fecaca'
              }}
              action={
                <Button
                  size="small"
                  onClick={() => fetchVehicles()}
                  style={{ borderRadius: 8 }}
                >
                  重试
                </Button>
              }
            />
          )}

          {/* 车辆展示区域 */}
          <VehicleGrid>
            <div style={{ marginBottom: 24 }}>
              <Row justify="space-between" align="middle">
                <Col>
                  <Title level={2} style={{
                    margin: 0,
                    color: '#1e293b',
                    fontWeight: 600
                  }}>
                    <CarOutlined style={{ marginRight: 12, color: '#667eea' }} />
                    车辆列表
                    {!loading && (
                      <Badge
                        count={pagination.total}
                        style={{
                          backgroundColor: '#667eea',
                          marginLeft: 12
                        }}
                      />
                    )}
                  </Title>
                </Col>
                <Col>
                  {!loading && vehicles.length > 0 && (
                    <Text type="secondary">
                      显示 {((pagination.current - 1) * pagination.pageSize) + 1}-{Math.min(pagination.current * pagination.pageSize, pagination.total)} 条，共 {pagination.total} 条
                    </Text>
                  )}
                </Col>
              </Row>
            </div>

            <Spin spinning={loading}>
              {vehicles.length > 0 ? (
                <>
                  <Row gutter={[24, 24]}>
                    {vehicles
                      .filter(vehicle => filterStatus === 'all' || vehicle.status === filterStatus)
                      .sort((a, b) => {
                        switch (sortBy) {
                          case 'price_asc':
                            return (a.price || 0) - (b.price || 0);
                          case 'price_desc':
                            return (b.price || 0) - (a.price || 0);
                          case 'year_desc':
                            return (b.year || 0) - (a.year || 0);
                          case 'year_asc':
                            return (a.year || 0) - (b.year || 0);
                          default:
                            return 0;
                        }
                      })
                      .map(vehicle => (
                        <Col
                          xs={24}
                          sm={viewMode === 'list' ? 24 : 12}
                          md={viewMode === 'list' ? 24 : 8}
                          lg={viewMode === 'list' ? 24 : 6}
                          key={vehicle.id}
                        >
                          <VehicleCard vehicle={vehicle} viewMode={viewMode} />
                        </Col>
                      ))}
                  </Row>

                  {/* 分页 */}
                  <div style={{
                    textAlign: 'center',
                    marginTop: 48,
                    padding: '24px 0',
                    borderTop: '1px solid #e2e8f0'
                  }}>
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
                      style={{
                        '& .ant-pagination-item': {
                          borderRadius: '8px',
                        },
                        '& .ant-pagination-item-active': {
                          background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                          borderColor: '#667eea',
                        }
                      }}
                    />
                  </div>
                </>
              ) : (
                !loading && (
                  <div style={{
                    textAlign: 'center',
                    padding: '80px 0',
                    background: 'linear-gradient(135deg, #f8fafc 0%, #ffffff 100%)',
                    borderRadius: 16,
                    border: '2px dashed #e2e8f0'
                  }}>
                    <div style={{
                      width: 120,
                      height: 120,
                      borderRadius: '50%',
                      background: 'linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%)',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      margin: '0 auto 24px'
                    }}>
                      <CarOutlined style={{ fontSize: '48px', color: '#94a3b8' }} />
                    </div>
                    <Title level={3} style={{ color: '#64748b', marginBottom: 8 }}>
                      暂无车辆数据
                    </Title>
                    <Text type="secondary" style={{ fontSize: 16 }}>
                      当前没有符合条件的车辆，请尝试调整筛选条件
                    </Text>
                  </div>
                )
              )}
            </Spin>
          </VehicleGrid>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default VehicleList;
