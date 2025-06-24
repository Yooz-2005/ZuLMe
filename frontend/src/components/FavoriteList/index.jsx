import React, { useState, useEffect } from 'react';
import { Card, Button, message, Empty, Spin, Row, Col, Typography, Space, Tag } from 'antd';
import { HeartFilled, CarOutlined, EyeOutlined, DeleteOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import favoriteService from '../../services/favoriteService';
import { getAllImages } from '../../utils/imageUtils';

const { Text, Title } = Typography;

// 样式组件
const StyledCard = styled(Card)`
  .ant-card-body {
    padding: 0;
  }
`;

const VehicleItem = styled.div`
  display: flex;
  padding: 20px;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.3s ease;
  
  &:hover {
    background-color: #fafafa;
  }
  
  &:last-child {
    border-bottom: none;
  }
`;

const VehicleImageContainer = styled.div`
  width: 120px;
  height: 80px;
  border-radius: 8px;
  margin-right: 16px;
  background-color: #f5f5f5;
  border: 1px solid #e8e8e8;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
`;

const VehicleImage = styled.img`
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
`;

const ImagePlaceholder = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  color: #94a3b8;
  font-size: 12px;

  .anticon {
    font-size: 24px;
    margin-bottom: 4px;
  }
`;

const VehicleInfo = styled.div`
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`;

const VehicleActions = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-end;
  min-width: 120px;
`;

const ActionButton = styled(Button)`
  margin-bottom: 8px;
  
  &:last-child {
    margin-bottom: 0;
  }
`;

// 车辆图片组件
const VehicleImageComponent = ({ imageString, vehicleName }) => {
  const [imageError, setImageError] = useState(false);

  // 使用与车辆详情页相同的图片解析逻辑
  const getVehicleImage = (imageString) => {
    console.log('收藏列表图片字符串:', imageString);

    if (!imageString) {
      return '/images/default-car.jpg';
    }

    // 使用项目中的图片解析工具
    const images = getAllImages(imageString);
    console.log('解析后的图片数组:', images);

    if (images && images.length > 0) {
      return images[0]; // 返回第一张图片
    }

    return '/images/default-car.jpg';
  };

  const handleImageError = () => {
    console.log('图片加载失败:', getVehicleImage(imageString));
    setImageError(true);
  };

  if (imageError) {
    return (
      <VehicleImageContainer>
        <ImagePlaceholder>
          <CarOutlined />
          <span>暂无图片</span>
        </ImagePlaceholder>
      </VehicleImageContainer>
    );
  }

  const imageUrl = getVehicleImage(imageString);
  console.log('最终图片URL:', imageUrl);

  return (
    <VehicleImageContainer>
      <VehicleImage
        src={imageUrl}
        alt={vehicleName || '车辆图片'}
        onError={handleImageError}
      />
    </VehicleImageContainer>
  );
};

const FavoriteList = () => {
  const [favoriteList, setFavoriteList] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();



  // 获取收藏列表
  const fetchFavoriteList = async () => {
    try {
      setLoading(true);
      const list = await favoriteService.getFavoriteList();
      setFavoriteList(list);
    } catch (error) {
      message.error('获取收藏列表失败');
      console.error('获取收藏列表失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 取消收藏
  const handleRemoveFavorite = async (vehicleId, vehicleName) => {
    try {
      const response = await favoriteService.toggleCollect(vehicleId);
      if (response.code === 200) {
        message.success(`已取消收藏 ${vehicleName}`);
        // 重新获取收藏列表
        fetchFavoriteList();
      } else {
        message.error(response.message || '取消收藏失败');
      }
    } catch (error) {
      message.error('取消收藏失败');
      console.error('取消收藏失败:', error);
    }
  };

  // 查看车辆详情
  const handleViewDetail = (vehicleId) => {
    navigate(`/vehicle/${vehicleId}`);
  };

  // 立即预订
  const handleReservation = (vehicleId) => {
    navigate(`/vehicle/${vehicleId}?action=reserve`);
  };

  useEffect(() => {
    fetchFavoriteList();
  }, []);

  if (loading) {
    return (
      <div style={{ textAlign: 'center', padding: '50px' }}>
        <Spin size="large" />
        <div style={{ marginTop: 16 }}>
          <Text type="secondary">正在加载收藏列表...</Text>
        </div>
      </div>
    );
  }

  if (!favoriteList || favoriteList.length === 0) {
    return (
      <Empty
        image={Empty.PRESENTED_IMAGE_SIMPLE}
        description={
          <div>
            <Text type="secondary">您还没有收藏任何车辆</Text>
            <br />
            <Text type="secondary">去看看有哪些心仪的车辆吧</Text>
          </div>
        }
      >
        <Button type="primary" onClick={() => navigate('/dashboard')}>
          去首页看看
        </Button>
      </Empty>
    );
  }

  return (
    <StyledCard>
      <div style={{ padding: '20px 20px 0 20px' }}>
        <Row justify="space-between" align="middle">
          <Col>
            <Title level={4} style={{ margin: 0 }}>
              <HeartFilled style={{ color: '#ff4d4f', marginRight: 8 }} />
              我的收藏
            </Title>
          </Col>
          <Col>
            <Tag color="blue">{favoriteList.length} 辆车</Tag>
          </Col>
        </Row>
      </div>
      
      <div>
        {favoriteList.map((vehicle) => (
          <VehicleItem key={vehicle.VehicleId}>
            <VehicleImageComponent
              imageString={vehicle.Image}
              vehicleName={vehicle.VehicleName}
            />
            
            <VehicleInfo>
              <div>
                <Title level={5} style={{ margin: '0 0 8px 0' }}>
                  {vehicle.VehicleName}
                </Title>
                <Space>
                  <Tag icon={<CarOutlined />} color="blue">
                    豪华轿车
                  </Tag>
                </Space>
              </div>
              
              <div>
                <Text type="secondary" style={{ fontSize: 12 }}>
                  收藏时间：{new Date().toLocaleDateString()}
                </Text>
              </div>
            </VehicleInfo>
            
            <VehicleActions>
              <ActionButton
                type="primary"
                size="small"
                icon={<EyeOutlined />}
                onClick={() => handleViewDetail(vehicle.VehicleId)}
              >
                查看详情
              </ActionButton>
              
              <ActionButton
                type="default"
                size="small"
                onClick={() => handleReservation(vehicle.VehicleId)}
              >
                立即预订
              </ActionButton>
              
              <ActionButton
                type="text"
                size="small"
                danger
                icon={<DeleteOutlined />}
                onClick={() => handleRemoveFavorite(vehicle.VehicleId, vehicle.VehicleName)}
              >
                取消收藏
              </ActionButton>
            </VehicleActions>
          </VehicleItem>
        ))}
      </div>
    </StyledCard>
  );
};

export default FavoriteList;
