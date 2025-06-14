import React from 'react';
import { Card, Tag, Button, Row, Col, Typography, Space, Rate } from 'antd';
import {
  SettingOutlined,
  CalendarOutlined,
  EnvironmentOutlined,
  DashboardOutlined,
  HeartOutlined,
  EyeOutlined
} from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import { VEHICLE_STATUS_LABELS } from '../../utils/constants';
import { getFirstImage, handleImageError } from '../../utils/imageUtils';

const { Text, Title } = Typography;

const StyledCard = styled(Card)`
  border-radius: 16px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  overflow: hidden;
  transition: all 0.3s ease;

  .ant-card-cover {
    position: relative;
    overflow: hidden;

    img {
      height: ${props => props.viewMode === 'list' ? '160px' : '220px'};
      object-fit: cover;
      transition: transform 0.3s ease;
    }

    &:hover img {
      transform: scale(1.05);
    }
  }

  .ant-card-body {
    padding: ${props => props.viewMode === 'list' ? '16px' : '20px'};
  }

  &:hover {
    box-shadow: 0 8px 32px rgba(0,0,0,0.15);
    transform: translateY(-4px);
    border-color: #667eea;
  }
`;

const ListCard = styled(Card)`
  border-radius: 16px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  overflow: hidden;
  transition: all 0.3s ease;

  .ant-card-body {
    padding: 24px;
  }

  &:hover {
    box-shadow: 0 8px 32px rgba(0,0,0,0.15);
    transform: translateY(-2px);
    border-color: #667eea;
  }
`;

const PriceText = styled.div`
  font-size: ${props => props.viewMode === 'list' ? '24px' : '20px'};
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 8px 0;
`;

const FeatureItem = styled.div`
  display: flex;
  align-items: center;
  color: #64748b;
  font-size: 14px;
  margin-bottom: 8px;

  .anticon {
    margin-right: 8px;
    color: #667eea;
    font-size: 16px;
  }
`;

const StatusTag = styled(Tag)`
  border-radius: 12px;
  padding: 4px 12px;
  font-weight: 600;
  border: none;

  &.status-available {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: white;
  }

  &.status-rented {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    color: white;
  }

  &.status-maintenance {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
    color: white;
  }
`;

const ActionButton = styled(Button)`
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.3s ease;

  &.ant-btn-primary {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border: none;

    &:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
    }
  }

  &:not(.ant-btn-primary) {
    border: 1px solid #e2e8f0;
    color: #64748b;

    &:hover {
      border-color: #667eea;
      color: #667eea;
    }
  }
`;

const VehicleCard = ({ vehicle, showActions = true, viewMode = 'grid' }) => {
  const navigate = useNavigate();

  const handleViewDetail = () => {
    navigate(`/vehicle/${vehicle.id}`);
  };

  const handleFavorite = (e) => {
    e.stopPropagation();
    // TODO: 实现收藏功能
    console.log('收藏车辆:', vehicle.id);
  };

  if (viewMode === 'list') {
    return (
      <ListCard
        hoverable
        onClick={handleViewDetail}
        style={{ cursor: 'pointer' }}
      >
        <Row gutter={24} align="middle">
          <Col span={8}>
            <div style={{
              borderRadius: 12,
              overflow: 'hidden',
              height: 160
            }}>
              <img
                alt={vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '车辆图片'}
                src={getFirstImage(vehicle.images, vehicle.brand)}
                onError={(e) => handleImageError(e, vehicle.brand)}
                style={{
                  width: '100%',
                  height: '100%',
                  objectFit: 'cover',
                  transition: 'transform 0.3s ease'
                }}
              />
            </div>
          </Col>
          <Col span={12}>
            <div>
              <Row justify="space-between" align="top" style={{ marginBottom: 12 }}>
                <Col>
                  <Title level={3} style={{ margin: 0, color: '#1e293b' }}>
                    {vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '未知车型'}
                  </Title>
                </Col>
                <Col>
                  <StatusTag className={`status-${vehicle.status}`}>
                    {VEHICLE_STATUS_LABELS[vehicle.status] || vehicle.status}
                  </StatusTag>
                </Col>
              </Row>

              <Row gutter={[16, 8]} style={{ marginBottom: 16 }}>
                <Col span={12}>
                  <FeatureItem>
                    <CalendarOutlined />
                    <span>{vehicle.year}年</span>
                  </FeatureItem>
                </Col>
                <Col span={12}>
                  <FeatureItem>
                    <DashboardOutlined />
                    <span>{vehicle.mileage ? `${vehicle.mileage}km` : '新车'}</span>
                  </FeatureItem>
                </Col>
                <Col span={12}>
                  <FeatureItem>
                    <SettingOutlined />
                    <span>{vehicle.color || '未知颜色'}</span>
                  </FeatureItem>
                </Col>
                <Col span={12}>
                  <FeatureItem>
                    <EnvironmentOutlined />
                    <span>{vehicle.location || '暂无位置'}</span>
                  </FeatureItem>
                </Col>
              </Row>

              <PriceText viewMode={viewMode}>
                ¥{vehicle.price || 0}/天
              </PriceText>
            </div>
          </Col>
          <Col span={4}>
            <Space direction="vertical" style={{ width: '100%' }}>
              <ActionButton
                type="primary"
                block
                icon={<EyeOutlined />}
                onClick={(e) => {
                  e.stopPropagation();
                  handleViewDetail();
                }}
              >
                查看详情
              </ActionButton>
              <ActionButton
                block
                icon={<HeartOutlined />}
                onClick={handleFavorite}
              >
                收藏
              </ActionButton>
              <Rate disabled defaultValue={4.5} style={{ fontSize: 14 }} />
            </Space>
          </Col>
        </Row>
      </ListCard>
    );
  }

  return (
    <StyledCard
      hoverable
      viewMode={viewMode}
      onClick={handleViewDetail}
      style={{ cursor: 'pointer' }}
      cover={
        <div style={{ position: 'relative' }}>
          <img
            alt={vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '车辆图片'}
            src={getFirstImage(vehicle.images, vehicle.brand)}
            onError={(e) => handleImageError(e, vehicle.brand)}
          />
          <div style={{
            position: 'absolute',
            top: 12,
            right: 12,
            background: 'rgba(255, 255, 255, 0.9)',
            borderRadius: 8,
            padding: '4px 8px'
          }}>
            <StatusTag className={`status-${vehicle.status}`}>
              {VEHICLE_STATUS_LABELS[vehicle.status] || vehicle.status}
            </StatusTag>
          </div>
          <Button
            type="text"
            icon={<HeartOutlined />}
            onClick={handleFavorite}
            style={{
              position: 'absolute',
              top: 12,
              left: 12,
              background: 'rgba(255, 255, 255, 0.9)',
              borderRadius: 8,
              color: '#64748b'
            }}
          />
        </div>
      }
      actions={showActions ? [
        <ActionButton
          type="primary"
          icon={<EyeOutlined />}
          onClick={(e) => {
            e.stopPropagation();
            handleViewDetail();
          }}
        >
          查看详情
        </ActionButton>,
        <ActionButton>
          立即预订
        </ActionButton>
      ] : undefined}
    >
      <div>
        <Title level={4} style={{ margin: '0 0 8px 0', color: '#1e293b' }}>
          {vehicle.brand && vehicle.style ? `${vehicle.brand} ${vehicle.style}` : '未知车型'}
        </Title>

        <Row gutter={[8, 8]} style={{ marginBottom: 12 }}>
          <Col span={12}>
            <FeatureItem>
              <CalendarOutlined />
              <span>{vehicle.year}年</span>
            </FeatureItem>
          </Col>
          <Col span={12}>
            <FeatureItem>
              <DashboardOutlined />
              <span>{vehicle.mileage ? `${vehicle.mileage}km` : '新车'}</span>
            </FeatureItem>
          </Col>
        </Row>

        <PriceText viewMode={viewMode}>
          ¥{vehicle.price || 0}/天
        </PriceText>

        {vehicle.location && (
          <FeatureItem style={{ marginTop: 8 }}>
            <EnvironmentOutlined />
            <span>{vehicle.location}</span>
          </FeatureItem>
        )}

        <div style={{ marginTop: 12 }}>
          <Rate disabled defaultValue={4.5} style={{ fontSize: 12 }} />
          <Text type="secondary" style={{ marginLeft: 8, fontSize: 12 }}>
            (4.5分)
          </Text>
        </div>
      </div>
    </StyledCard>
  );
};

export default VehicleCard;
