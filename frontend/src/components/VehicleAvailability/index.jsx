import React, { useState, useEffect } from 'react';
import {
  Card,
  DatePicker,
  Button,
  Space,
  Alert,
  Spin,
  Tag,
  Typography,
  Row,
  Col,
  Statistic,
  Calendar,
  Badge,
  Modal,
  List,
  Avatar
} from 'antd';
import {
  CheckCircleOutlined,
  ClockCircleOutlined,
  CarOutlined,
  ToolOutlined,
  StopOutlined,
  CalendarOutlined
} from '@ant-design/icons';
import styled from 'styled-components';
import moment from 'moment';
import vehicleService from '../../services/vehicleService';

const { RangePicker } = DatePicker;
const { Title, Text } = Typography;

const StyledCard = styled(Card)`
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  margin-bottom: 24px;
  
  .ant-card-head {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 16px 16px 0 0;
    
    .ant-card-head-title {
      color: white;
      font-weight: 600;
    }
  }
`;

const StatusTag = styled(Tag)`
  border-radius: 16px;
  padding: 4px 12px;
  font-weight: 500;
`;

const CalendarWrapper = styled.div`
  .ant-picker-calendar {
    border-radius: 16px;
    overflow: hidden;
  }
  
  .ant-picker-calendar-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    padding: 16px;
  }
`;

const VehicleAvailability = ({ vehicleId, onDateSelect, showCalendar = true }) => {
  const [loading, setLoading] = useState(false);
  const [availability, setAvailability] = useState(null);
  const [selectedDates, setSelectedDates] = useState(null);
  const [inventoryData, setInventoryData] = useState([]);
  const [stats, setStats] = useState({});
  const [calendarVisible, setCalendarVisible] = useState(false);

  // 状态配置
  const statusConfig = {
    1: { text: '可租用', color: 'green', icon: <CheckCircleOutlined /> },
    2: { text: '已预订', color: 'orange', icon: <ClockCircleOutlined /> },
    3: { text: '租用中', color: 'blue', icon: <CarOutlined /> },
    4: { text: '维护中', color: 'purple', icon: <ToolOutlined /> },
    5: { text: '不可用', color: 'red', icon: <StopOutlined /> }
  };

  // 检查车辆可用性
  const checkAvailability = async (dates) => {
    if (!dates || dates.length !== 2 || !vehicleId) return;
    
    setLoading(true);
    try {
      const response = await vehicleService.checkAvailability({
        vehicle_id: vehicleId,
        start_date: dates[0].format('YYYY-MM-DD'),
        end_date: dates[1].format('YYYY-MM-DD')
      });
      
      if (response.code === 200) {
        setAvailability(response.data);
        if (onDateSelect) {
          onDateSelect(dates, response.data.is_available);
        }
      }
    } catch (error) {
      console.error('检查可用性失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 获取库存统计
  const fetchInventoryStats = async () => {
    try {
      const response = await vehicleService.getInventoryStats({
        merchant_id: 1 // 这里应该从用户信息中获取
      });
      
      if (response.code === 200) {
        setStats(response.data);
      }
    } catch (error) {
      console.error('获取库存统计失败:', error);
    }
  };

  // 获取车辆库存日历数据
  const fetchInventoryCalendar = async () => {
    if (!vehicleId) return;
    
    try {
      const response = await vehicleService.getVehicleInventory({
        vehicle_id: vehicleId,
        start_date: moment().format('YYYY-MM-DD'),
        end_date: moment().add(3, 'months').format('YYYY-MM-DD')
      });
      
      if (response.code === 200) {
        setInventoryData(response.data.inventories || []);
      }
    } catch (error) {
      console.error('获取库存日历失败:', error);
    }
  };

  useEffect(() => {
    if (vehicleId) {
      fetchInventoryCalendar();
    }
    fetchInventoryStats();
  }, [vehicleId]);

  // 处理日期选择
  const handleDateChange = (dates) => {
    setSelectedDates(dates);
    if (dates && dates.length === 2) {
      checkAvailability(dates);
    } else {
      setAvailability(null);
    }
  };

  // 日历单元格渲染
  const dateCellRender = (value) => {
    const dateStr = value.format('YYYY-MM-DD');
    const dayInventory = inventoryData.filter(item => {
      const startDate = moment(item.start_date);
      const endDate = moment(item.end_date);
      return value.isBetween(startDate, endDate, 'day', '[]');
    });

    if (dayInventory.length === 0) return null;

    const status = dayInventory[0].status;
    const config = statusConfig[status];

    return (
      <Badge
        status={config.color}
        text={config.text}
        style={{ fontSize: '12px' }}
      />
    );
  };

  // 渲染可用性结果
  const renderAvailabilityResult = () => {
    if (!availability) return null;

    return (
      <Alert
        message={
          availability.is_available 
            ? "车辆在选定日期可用" 
            : "车辆在选定日期不可用"
        }
        description={
          availability.is_available
            ? "您可以继续预订此车辆"
            : "请选择其他日期或车辆"
        }
        type={availability.is_available ? "success" : "warning"}
        showIcon
        style={{ marginTop: 16 }}
      />
    );
  };

  return (
    <div>
      {/* 库存统计卡片 */}
      {Object.keys(stats).length > 0 && (
        <StyledCard title="库存概览">
          <Row gutter={16}>
            <Col span={6}>
              <Statistic
                title="总车辆"
                value={stats.total}
                prefix={<CarOutlined />}
              />
            </Col>
            <Col span={6}>
              <Statistic
                title="可用"
                value={stats.available}
                valueStyle={{ color: '#52c41a' }}
                prefix={<CheckCircleOutlined />}
              />
            </Col>
            <Col span={6}>
              <Statistic
                title="已预订"
                value={stats.reserved}
                valueStyle={{ color: '#fa8c16' }}
                prefix={<ClockCircleOutlined />}
              />
            </Col>
            <Col span={6}>
              <Statistic
                title="租用中"
                value={stats.rented}
                valueStyle={{ color: '#1890ff' }}
                prefix={<CarOutlined />}
              />
            </Col>
          </Row>
        </StyledCard>
      )}

      {/* 日期选择和可用性检查 */}
      <StyledCard title="检查车辆可用性">
        <Space direction="vertical" style={{ width: '100%' }}>
          <div>
            <Text strong>选择租用日期：</Text>
            <RangePicker
              style={{ marginLeft: 16, width: 300 }}
              value={selectedDates}
              onChange={handleDateChange}
              disabledDate={(current) => current && current < moment().startOf('day')}
              placeholder={['开始日期', '结束日期']}
            />
            <Button
              type="link"
              icon={<CalendarOutlined />}
              onClick={() => setCalendarVisible(true)}
            >
              查看日历
            </Button>
          </div>
          
          {loading && <Spin tip="检查中..." />}
          {renderAvailabilityResult()}
        </Space>
      </StyledCard>

      {/* 库存状态说明 */}
      <StyledCard title="状态说明">
        <Space wrap>
          {Object.entries(statusConfig).map(([status, config]) => (
            <StatusTag key={status} color={config.color} icon={config.icon}>
              {config.text}
            </StatusTag>
          ))}
        </Space>
      </StyledCard>

      {/* 库存日历模态框 */}
      <Modal
        title="车辆库存日历"
        open={calendarVisible}
        onCancel={() => setCalendarVisible(false)}
        footer={null}
        width={800}
      >
        <CalendarWrapper>
          <Calendar
            dateCellRender={dateCellRender}
            headerRender={({ value, type, onChange, onTypeChange }) => (
              <div style={{ padding: 8 }}>
                <Row justify="space-between" align="middle">
                  <Col>
                    <Title level={4} style={{ margin: 0, color: 'white' }}>
                      {value.format('YYYY年MM月')}
                    </Title>
                  </Col>
                  <Col>
                    <Space>
                      <Button 
                        size="small" 
                        onClick={() => onChange(value.clone().subtract(1, 'month'))}
                      >
                        上月
                      </Button>
                      <Button 
                        size="small" 
                        onClick={() => onChange(moment())}
                      >
                        今天
                      </Button>
                      <Button 
                        size="small" 
                        onClick={() => onChange(value.clone().add(1, 'month'))}
                      >
                        下月
                      </Button>
                    </Space>
                  </Col>
                </Row>
              </div>
            )}
          />
        </CalendarWrapper>
        
        {/* 当前月份的预订列表 */}
        <Card title="当前预订" style={{ marginTop: 16 }}>
          <List
            dataSource={inventoryData.filter(item => 
              moment(item.start_date).month() === moment().month()
            )}
            renderItem={item => {
              const config = statusConfig[item.status];
              return (
                <List.Item>
                  <List.Item.Meta
                    avatar={<Avatar icon={config.icon} style={{ backgroundColor: config.color }} />}
                    title={
                      <Space>
                        <StatusTag color={config.color}>{config.text}</StatusTag>
                        <Text>{moment(item.start_date).format('MM-DD')} 至 {moment(item.end_date).format('MM-DD')}</Text>
                      </Space>
                    }
                    description={item.notes || '无备注'}
                  />
                </List.Item>
              );
            }}
            locale={{ emptyText: '本月暂无预订记录' }}
          />
        </Card>
      </Modal>
    </div>
  );
};

export default VehicleAvailability;
