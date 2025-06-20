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
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
    border: none;
  }

  .ant-modal-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-bottom: none;
    padding: 28px 32px;
    position: relative;

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%);
      pointer-events: none;
    }

    .ant-modal-title {
      color: white;
      font-size: 22px;
      font-weight: 700;
      text-shadow: 0 2px 4px rgba(0,0,0,0.2);
      display: flex;
      align-items: center;
      gap: 12px;

      .anticon {
        font-size: 24px;
        filter: drop-shadow(0 2px 4px rgba(0,0,0,0.2));
      }
    }
  }

  .ant-modal-close {
    color: white;
    font-size: 18px;
    top: 24px;
    right: 24px;

    &:hover {
      color: rgba(255,255,255,0.8);
    }
  }

  .ant-modal-body {
    padding: 32px;
    background: #fafbfc;
  }
`;

const VehicleInfoCard = styled(Card)`
  margin-bottom: 28px;
  border: none;
  border-radius: 12px;
  background: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;

  &:hover {
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
    transform: translateY(-2px);
  }

  .ant-card-body {
    padding: 24px;
  }
`;

const PriceDisplay = styled.div`
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 24px;
  border-radius: 12px;
  margin: 24px 0;
  text-align: center;
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.25);

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%);
    pointer-events: none;
  }

  .price-label {
    color: rgba(255, 255, 255, 0.9);
    font-size: 16px;
    margin-bottom: 8px;
    font-weight: 500;
    text-shadow: 0 1px 2px rgba(0,0,0,0.1);
  }

  .price-value {
    color: white;
    font-size: 32px;
    font-weight: 800;
    text-shadow: 0 2px 4px rgba(0,0,0,0.2);
    position: relative;
    z-index: 1;
  }

  .price-unit {
    color: rgba(255, 255, 255, 0.9);
    font-size: 18px;
    margin-left: 6px;
    font-weight: 500;
  }
`;

const StyledForm = styled(Form)`
  .ant-form-item-label > label {
    font-weight: 600;
    color: #374151;
    font-size: 16px;
  }

  .ant-picker {
    border-radius: 8px;
    border: 2px solid #e5e7eb;
    transition: all 0.3s ease;

    &:hover {
      border-color: #667eea;
    }

    &.ant-picker-focused {
      border-color: #667eea;
      box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
    }
  }

  .ant-input {
    border-radius: 8px;
    border: 2px solid #e5e7eb;
    transition: all 0.3s ease;

    &:hover {
      border-color: #667eea;
    }

    &:focus {
      border-color: #667eea;
      box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
    }
  }
`;

const StyledButton = styled(Button)`
  border-radius: 8px;
  font-weight: 600;
  height: 48px;
  font-size: 16px;
  transition: all 0.3s ease;

  &.ant-btn-primary {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border: none;
    box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 8px 24px rgba(102, 126, 234, 0.4);
    }

    &:disabled {
      background: #d1d5db;
      transform: none;
      box-shadow: none;
    }
  }

  &:not(.ant-btn-primary) {
    border: 2px solid #e5e7eb;
    color: #6b7280;

    &:hover {
      border-color: #667eea;
      color: #667eea;
      transform: translateY(-1px);
    }
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
        <div style={{
          textAlign: 'center',
          padding: '24px',
          background: 'white',
          borderRadius: '12px',
          marginBottom: '24px',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.08)'
        }}>
          <Spin tip="检查车辆可用性..." size="large" />
        </div>
      );
    }

    if (availability) {
      return (
        <Alert
          type={availability.is_available ? 'success' : 'error'}
          message={availability.is_available ? '✅ 车辆可用' : '❌ 车辆不可用'}
          description={availability.message}
          showIcon
          style={{
            marginBottom: 24,
            borderRadius: 12,
            border: 'none',
            boxShadow: '0 4px 12px rgba(0, 0, 0, 0.08)'
          }}
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
      width={680}
      destroyOnClose
      centered
    >
      {vehicle && (
        <VehicleInfoCard>
          <Row align="middle">
            <Col span={16}>
              <div style={{ display: 'flex', alignItems: 'center', gap: '12px', marginBottom: '8px' }}>
                <CarOutlined style={{ fontSize: '20px', color: '#667eea' }} />
                <Title level={4} style={{ margin: 0, color: '#1f2937' }}>
                  {vehicle.brand} {vehicle.style}
                </Title>
              </div>
              <Text type="secondary" style={{ fontSize: '14px' }}>
                {vehicle.year}年 · {vehicle.color || '标准色'}
              </Text>
            </Col>
            <Col span={8} style={{ textAlign: 'right' }}>
              <div style={{
                background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                color: 'white',
                padding: '8px 16px',
                borderRadius: '8px',
                display: 'inline-block'
              }}>
                <Text strong style={{ fontSize: '18px', color: 'white' }}>
                  ¥{vehicle.price}/天
                </Text>
              </div>
            </Col>
          </Row>
        </VehicleInfoCard>
      )}

      <StyledForm
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

        <Form.Item style={{ marginBottom: 0, marginTop: 32 }}>
          <Row gutter={16}>
            <Col span={12}>
              <StyledButton
                size="large"
                block
                onClick={onCancel}
              >
                取消
              </StyledButton>
            </Col>
            <Col span={12}>
              <StyledButton
                type="primary"
                size="large"
                block
                htmlType="submit"
                loading={loading}
                disabled={!availability?.is_available || checkingAvailability}
              >
                确认预订
              </StyledButton>
            </Col>
          </Row>
        </Form.Item>
      </StyledForm>
    </StyledModal>
  );
};

export default ReservationForm;
