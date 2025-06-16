import React, { useState, useEffect } from 'react';
import { 
  Modal, 
  Form, 
  DatePicker, 
  Input, 
  Button, 
  Space, 
  Typography, 
  Row, 
  Col, 
  Card,
  message,
  Spin,
  Alert
} from 'antd';
import { CalendarOutlined, ClockCircleOutlined, CarOutlined } from '@ant-design/icons';
import styled from 'styled-components';
import dayjs from 'dayjs';
import vehicleService from '../../services/vehicleService';

const { Title, Text } = Typography;
const { RangePicker } = DatePicker;
const { TextArea } = Input;

const StyledModal = styled(Modal)`
  .ant-modal-content {
    border-radius: 12px;
    overflow: hidden;
  }
  
  .ant-modal-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-bottom: none;
    padding: 24px;
    
    .ant-modal-title {
      color: white;
      font-size: 20px;
      font-weight: 600;
    }
  }
  
  .ant-modal-body {
    padding: 24px;
  }
`;

const VehicleInfoCard = styled(Card)`
  margin-bottom: 24px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  
  .ant-card-body {
    padding: 16px;
  }
`;

const PriceDisplay = styled.div`
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  padding: 16px;
  border-radius: 8px;
  margin: 16px 0;
  text-align: center;
  
  .price-label {
    color: #64748b;
    font-size: 14px;
    margin-bottom: 4px;
  }
  
  .price-value {
    color: #1e293b;
    font-size: 24px;
    font-weight: 700;
  }
  
  .price-unit {
    color: #64748b;
    font-size: 16px;
    margin-left: 4px;
  }
`;

const ReservationForm = ({ 
  visible, 
  onCancel, 
  onSuccess, 
  vehicle,
  initialDates = null 
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [checkingAvailability, setCheckingAvailability] = useState(false);
  const [availability, setAvailability] = useState(null);
  const [selectedDates, setSelectedDates] = useState(initialDates);
  const [totalDays, setTotalDays] = useState(0);
  const [totalPrice, setTotalPrice] = useState(0);

  useEffect(() => {
    if (visible && initialDates) {
      setSelectedDates(initialDates);
      form.setFieldsValue({
        dates: initialDates
      });
      calculatePrice(initialDates);
    }
  }, [visible, initialDates, form]);

  // 禁用过去的日期
  const disabledDate = (current) => {
    return current && current < dayjs().startOf('day');
  };

  // 计算价格
  const calculatePrice = (dates) => {
    if (dates && dates.length === 2 && vehicle?.price) {
      const days = dates[1].diff(dates[0], 'day');
      setTotalDays(days);
      setTotalPrice(days * vehicle.price);
    } else {
      setTotalDays(0);
      setTotalPrice(0);
    }
  };

  // 处理日期变化
  const handleDateChange = async (dates) => {
    setSelectedDates(dates);
    setAvailability(null);
    
    if (dates && dates.length === 2) {
      calculatePrice(dates);
      await checkAvailability(dates);
    } else {
      setTotalDays(0);
      setTotalPrice(0);
    }
  };

  // 检查车辆可用性
  const checkAvailability = async (dates) => {
    if (!dates || dates.length !== 2 || !vehicle?.id) return;
    
    setCheckingAvailability(true);
    try {
      const response = await vehicleService.checkAvailability({
        vehicle_id: vehicle.id,
        start_date: dates[0].format('YYYY-MM-DD'),
        end_date: dates[1].format('YYYY-MM-DD')
      });
      
      if (response.code === 200) {
        setAvailability(response.data);
      } else {
        message.error(response.message || '检查可用性失败');
      }
    } catch (error) {
      console.error('检查可用性失败:', error);
      message.error('检查可用性失败，请稍后重试');
    } finally {
      setCheckingAvailability(false);
    }
  };

  // 提交预订
  const handleSubmit = async (values) => {
    if (!availability?.is_available) {
      message.error('车辆在选定时间段不可用');
      return;
    }

    setLoading(true);
    try {
      const response = await vehicleService.createReservation({
        vehicle_id: vehicle.id,
        start_date: values.dates[0].format('YYYY-MM-DD'),
        end_date: values.dates[1].format('YYYY-MM-DD'),
        notes: values.notes || ''
      });

      if (response.code === 200) {
        message.success('预订创建成功！');
        form.resetFields();
        setSelectedDates(null);
        setAvailability(null);
        setTotalDays(0);
        setTotalPrice(0);
        
        if (onSuccess) {
          onSuccess(response.data);
        }
      } else {
        message.error(response.message || '预订失败');
      }
    } catch (error) {
      console.error('创建预订失败:', error);
      message.error('创建预订失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 渲染可用性状态
  const renderAvailabilityStatus = () => {
    if (checkingAvailability) {
      return (
        <div style={{ textAlign: 'center', padding: '16px' }}>
          <Spin tip="检查车辆可用性..." />
        </div>
      );
    }

    if (availability) {
      return (
        <Alert
          type={availability.is_available ? 'success' : 'error'}
          message={availability.is_available ? '车辆可用' : '车辆不可用'}
          description={availability.message}
          showIcon
          style={{ marginBottom: 16 }}
        />
      );
    }

    return null;
  };

  return (
    <StyledModal
      title={
        <Space>
          <CarOutlined />
          预订车辆
        </Space>
      }
      open={visible}
      onCancel={onCancel}
      footer={null}
      width={600}
      destroyOnClose
    >
      {vehicle && (
        <VehicleInfoCard>
          <Row align="middle">
            <Col span={16}>
              <Title level={4} style={{ margin: 0 }}>
                {vehicle.brand} {vehicle.style}
              </Title>
              <Text type="secondary">
                {vehicle.year}年 · {vehicle.color || '标准色'}
              </Text>
            </Col>
            <Col span={8} style={{ textAlign: 'right' }}>
              <Text strong style={{ fontSize: '18px', color: '#667eea' }}>
                ¥{vehicle.price}/天
              </Text>
            </Col>
          </Row>
        </VehicleInfoCard>
      )}

      <Form
        form={form}
        layout="vertical"
        onFinish={handleSubmit}
        initialValues={{
          dates: initialDates
        }}
      >
        <Form.Item
          name="dates"
          label="租用日期"
          rules={[
            { required: true, message: '请选择租用日期' }
          ]}
        >
          <RangePicker
            style={{ width: '100%' }}
            placeholder={['取车日期', '还车日期']}
            disabledDate={disabledDate}
            onChange={handleDateChange}
            size="large"
          />
        </Form.Item>

        {renderAvailabilityStatus()}

        {totalDays > 0 && (
          <PriceDisplay>
            <div className="price-label">租用 {totalDays} 天，总费用</div>
            <div className="price-value">
              ¥{totalPrice.toLocaleString()}
              <span className="price-unit">元</span>
            </div>
          </PriceDisplay>
        )}

        <Form.Item
          name="notes"
          label="备注信息"
        >
          <TextArea
            placeholder="请输入特殊需求或备注信息（可选）"
            rows={3}
            maxLength={200}
            showCount
          />
        </Form.Item>

        <Form.Item style={{ marginBottom: 0, marginTop: 24 }}>
          <Row gutter={12}>
            <Col span={12}>
              <Button 
                size="large" 
                block 
                onClick={onCancel}
              >
                取消
              </Button>
            </Col>
            <Col span={12}>
              <Button
                type="primary"
                size="large"
                block
                htmlType="submit"
                loading={loading}
                disabled={!availability?.is_available || checkingAvailability}
              >
                确认预订
              </Button>
            </Col>
          </Row>
        </Form.Item>
      </Form>
    </StyledModal>
  );
};

export default ReservationForm;
