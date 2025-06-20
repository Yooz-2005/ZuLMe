import React, { useState, useEffect } from 'react';
import {
  Layout,
  Table,
  Button,
  Space,
  Modal,
  Form,
  Input,
  Switch,
  InputNumber,
  Upload,
  message,
  Popconfirm,
  Tag,
  Avatar,
  Row,
  Col,
  Card,
  Statistic
} from 'antd';
import {
  PlusOutlined,
  EditOutlined,
  DeleteOutlined,
  UploadOutlined,
  CarOutlined,
  GlobalOutlined,
  FireOutlined,
  EyeOutlined
} from '@ant-design/icons';
import styled from 'styled-components';
import vehicleService from '../../services/vehicleService';

const { Header, Content } = Layout;

const StyledLayout = styled(Layout)`
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
`;

const ContentWrapper = styled.div`
  padding: 24px;
  max-width: 1400px;
  margin: 0 auto;
`;

const StatsCard = styled(Card)`
  border-radius: 16px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  margin-bottom: 24px;
  
  .ant-card-body {
    padding: 20px;
  }
`;

const BrandTable = styled(Table)`
  .ant-table {
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 4px 16px rgba(0,0,0,0.08);
  }
  
  .ant-table-thead > tr > th {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    font-weight: 600;
    border: none;
  }
`;

const BrandManagement = () => {
  const [brands, setBrands] = useState([]);
  const [loading, setLoading] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);
  const [editingBrand, setEditingBrand] = useState(null);
  const [form] = Form.useForm();
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });
  const [stats, setStats] = useState({
    total: 0,
    hot: 0,
    active: 0
  });

  // 获取品牌列表
  const fetchBrands = async (page = 1, pageSize = 10) => {
    setLoading(true);
    try {
      const response = await vehicleService.getBrandList({
        page,
        page_size: pageSize
      });
      
      if (response.code === 200) {
        setBrands(response.data.vehicle_brands || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.total || 0
        });
        
        // 更新统计信息
        const brandList = response.data.vehicle_brands || [];
        setStats({
          total: response.data.total || 0,
          hot: brandList.filter(b => b.is_hot === 1).length,
          active: brandList.filter(b => b.status === 1).length
        });
      }
    } catch (error) {
      message.error('获取品牌列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBrands();
  }, []);

  // 处理表格分页
  const handleTableChange = (pagination) => {
    fetchBrands(pagination.current, pagination.pageSize);
  };

  // 打开创建/编辑模态框
  const openModal = (brand = null) => {
    setEditingBrand(brand);
    setModalVisible(true);
    if (brand) {
      form.setFieldsValue({
        ...brand,
        status: brand.status === 1,
        is_hot: brand.is_hot === 1
      });
    } else {
      form.resetFields();
    }
  };

  // 关闭模态框
  const closeModal = () => {
    setModalVisible(false);
    setEditingBrand(null);
    form.resetFields();
  };

  // 保存品牌
  const saveBrand = async (values) => {
    try {
      const data = {
        ...values,
        status: values.status ? 1 : 0,
        is_hot: values.is_hot ? 1 : 0
      };

      if (editingBrand) {
        await vehicleService.updateBrand({ ...data, id: editingBrand.id });
        message.success('更新品牌成功');
      } else {
        await vehicleService.createBrand(data);
        message.success('创建品牌成功');
      }
      
      closeModal();
      fetchBrands(pagination.current, pagination.pageSize);
    } catch (error) {
      message.error(editingBrand ? '更新品牌失败' : '创建品牌失败');
    }
  };

  // 删除品牌
  const deleteBrand = async (id) => {
    try {
      await vehicleService.deleteBrand({ id });
      message.success('删除品牌成功');
      fetchBrands(pagination.current, pagination.pageSize);
    } catch (error) {
      message.error('删除品牌失败');
    }
  };

  const columns = [
    {
      title: '品牌',
      dataIndex: 'name',
      key: 'name',
      render: (text, record) => (
        <Space>
          <Avatar 
            size={40}
            src={record.logo}
            icon={<CarOutlined />}
            style={{
              background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
            }}
          />
          <div>
            <div style={{ fontWeight: 600 }}>{record.name}</div>
            {record.english_name && (
              <div style={{ fontSize: 12, color: '#64748b' }}>
                {record.english_name}
              </div>
            )}
          </div>
        </Space>
      )
    },
    {
      title: '国家',
      dataIndex: 'country',
      key: 'country',
      render: (text) => (
        <Space>
          <GlobalOutlined style={{ color: '#667eea' }} />
          {text || '未知'}
        </Space>
      )
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (status) => (
        <Tag color={status === 1 ? 'green' : 'red'}>
          {status === 1 ? '启用' : '禁用'}
        </Tag>
      )
    },
    {
      title: '热门',
      dataIndex: 'is_hot',
      key: 'is_hot',
      render: (isHot) => (
        isHot === 1 ? (
          <Tag color="orange" icon={<FireOutlined />}>热门</Tag>
        ) : (
          <Tag>普通</Tag>
        )
      )
    },
    {
      title: '排序',
      dataIndex: 'sort',
      key: 'sort',
      sorter: true
    },
    {
      title: '操作',
      key: 'action',
      render: (_, record) => (
        <Space>
          <Button
            type="primary"
            size="small"
            icon={<EditOutlined />}
            onClick={() => openModal(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除这个品牌吗？"
            onConfirm={() => deleteBrand(record.id)}
            okText="确定"
            cancelText="取消"
          >
            <Button
              danger
              size="small"
              icon={<DeleteOutlined />}
            >
              删除
            </Button>
          </Popconfirm>
        </Space>
      )
    }
  ];

  return (
    <StyledLayout>
      <Content>
        <ContentWrapper>
          {/* 统计卡片 */}
          <Row gutter={24} style={{ marginBottom: 24 }}>
            <Col span={8}>
              <StatsCard>
                <Statistic
                  title="总品牌数"
                  value={stats.total}
                  prefix={<CarOutlined style={{ color: '#667eea' }} />}
                />
              </StatsCard>
            </Col>
            <Col span={8}>
              <StatsCard>
                <Statistic
                  title="热门品牌"
                  value={stats.hot}
                  prefix={<FireOutlined style={{ color: '#f59e0b' }} />}
                />
              </StatsCard>
            </Col>
            <Col span={8}>
              <StatsCard>
                <Statistic
                  title="启用品牌"
                  value={stats.active}
                  prefix={<EyeOutlined style={{ color: '#10b981' }} />}
                />
              </StatsCard>
            </Col>
          </Row>

          {/* 操作栏 */}
          <div style={{ marginBottom: 16 }}>
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={() => openModal()}
              size="large"
            >
              添加品牌
            </Button>
          </div>

          {/* 品牌表格 */}
          <BrandTable
            columns={columns}
            dataSource={brands}
            rowKey="id"
            loading={loading}
            pagination={{
              ...pagination,
              showSizeChanger: true,
              showQuickJumper: true,
              showTotal: (total, range) => 
                `第 ${range[0]}-${range[1]} 条，共 ${total} 条`
            }}
            onChange={handleTableChange}
          />

          {/* 创建/编辑模态框 */}
          <Modal
            title={editingBrand ? '编辑品牌' : '添加品牌'}
            open={modalVisible}
            onCancel={closeModal}
            footer={null}
            width={600}
          >
            <Form
              form={form}
              layout="vertical"
              onFinish={saveBrand}
            >
              <Row gutter={16}>
                <Col span={12}>
                  <Form.Item
                    name="name"
                    label="品牌名称"
                    rules={[{ required: true, message: '请输入品牌名称' }]}
                  >
                    <Input placeholder="请输入品牌名称" />
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Form.Item
                    name="english_name"
                    label="英文名称"
                  >
                    <Input placeholder="请输入英文名称" />
                  </Form.Item>
                </Col>
              </Row>
              
              <Row gutter={16}>
                <Col span={12}>
                  <Form.Item
                    name="country"
                    label="品牌国家"
                  >
                    <Input placeholder="请输入品牌国家" />
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Form.Item
                    name="sort"
                    label="排序"
                  >
                    <InputNumber 
                      placeholder="排序值" 
                      style={{ width: '100%' }}
                      min={0}
                    />
                  </Form.Item>
                </Col>
              </Row>

              <Form.Item
                name="logo"
                label="品牌Logo"
              >
                <Input placeholder="请输入Logo URL" />
              </Form.Item>

              <Form.Item
                name="description"
                label="品牌描述"
              >
                <Input.TextArea 
                  rows={3} 
                  placeholder="请输入品牌描述" 
                />
              </Form.Item>

              <Row gutter={16}>
                <Col span={12}>
                  <Form.Item
                    name="status"
                    label="启用状态"
                    valuePropName="checked"
                  >
                    <Switch checkedChildren="启用" unCheckedChildren="禁用" />
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Form.Item
                    name="is_hot"
                    label="热门品牌"
                    valuePropName="checked"
                  >
                    <Switch checkedChildren="是" unCheckedChildren="否" />
                  </Form.Item>
                </Col>
              </Row>

              <Form.Item>
                <Space>
                  <Button type="primary" htmlType="submit">
                    {editingBrand ? '更新' : '创建'}
                  </Button>
                  <Button onClick={closeModal}>
                    取消
                  </Button>
                </Space>
              </Form.Item>
            </Form>
          </Modal>
        </ContentWrapper>
      </Content>
    </StyledLayout>
  );
};

export default BrandManagement;
