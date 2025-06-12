import React from 'react';
import { Card, Tag, Button, Row, Col, Typography } from 'antd';
import { CarOutlined, UserOutlined, SettingOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import { VEHICLE_STATUS_LABELS, VEHICLE_TYPE_LABELS } from '../../utils/constants';

const { Meta } = Card;
const { Text, Title } = Typography;

const StyledCard = styled(Card)`
  .ant-card-cover img {
    height: 200px;
    object-fit: cover;
  }
  
  &:hover {
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
    transform: translateY(-2px);
    transition: all 0.3s ease;
  }
`;

const PriceText = styled(Title)`
  color: #ff4d4f !important;
  margin: 8px 0 !important;
`;

const FeatureRow = styled(Row)`
  margin: 8px 0;
  
  .feature-item {
    display: flex;
    align-items: center;
    color: #666;
    font-size: 12px;
    
    .anticon {
      margin-right: 4px;
      color: #1890ff;
    }
  }
`;

const VehicleCard = ({ vehicle, showActions = true }) => {
  const navigate = useNavigate();

  const handleViewDetail = () => {
    navigate(`/vehicle/${vehicle.id}`);
  };

  const getStatusColor = (status) => {
    switch (status) {
      case 'available':
        return 'green';
      case 'rented':
        return 'red';
      case 'maintenance':
        return 'orange';
      default:
        return 'default';
    }
  };

  return (
    <StyledCard
      hoverable
      cover={
        <img
          alt={vehicle.name || '车辆图片'}
          src={vehicle.images || '/images/my-car-a.jpg'}
          onError={(e) => {
            e.target.src = '/images/my-car-a.jpg';
          }}
        />
      }
      actions={showActions ? [
        <Button type="primary" onClick={handleViewDetail}>
          查看详情
        </Button>,
        <Button>
          立即预订
        </Button>
      ] : undefined}
    >
      <Meta
        title={
          <div>
            <Row justify="space-between" align="middle">
              <Col>
                <Title level={5} style={{ margin: 0 }}>
                  {vehicle.name || `${vehicle.brand} ${vehicle.style}` || '未知车型'}
                </Title>
              </Col>
              <Col>
                <Tag color={getStatusColor(vehicle.status)}>
                  {VEHICLE_STATUS_LABELS[vehicle.status] || vehicle.status}
                </Tag>
              </Col>
            </Row>
          </div>
        }
        description={
          <div>
            <Text type="secondary">{vehicle.brand} {vehicle.style}</Text>
            
            <FeatureRow gutter={16}>
              <Col span={8}>
                <div className="feature-item">
                  <UserOutlined />
                  <span>{vehicle.seats || 5}座</span>
                </div>
              </Col>
              <Col span={8}>
                <div className="feature-item">
                  <SettingOutlined />
                  <span>{vehicle.transmission || '自动'}</span>
                </div>
              </Col>
              <Col span={8}>
                <div className="feature-item">
                  <CarOutlined />
                  <span>{VEHICLE_TYPE_LABELS[vehicle.vehicle_type] || vehicle.vehicle_type}</span>
                </div>
              </Col>
            </FeatureRow>

            <PriceText level={4}>
              ¥{vehicle.price || 0}/天
            </PriceText>
            
            {vehicle.location && (
              <Text type="secondary" style={{ fontSize: '12px' }}>
                📍 {vehicle.location}
              </Text>
            )}
          </div>
        }
      />
    </StyledCard>
  );
};

export default VehicleCard;
