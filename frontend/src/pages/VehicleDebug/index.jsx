import React, { useState, useEffect } from 'react';
import { Card, Button, Typography, Space, Table, Image, Alert } from 'antd';
import vehicleService from '../../services/vehicleService';
import { parseImages, getFirstImage } from '../../utils/imageUtils';

const { Title, Text, Paragraph } = Typography;

const VehicleDebug = () => {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchVehicles = async () => {
    setLoading(true);
    try {
      const response = await vehicleService.getVehicleList({
        page: 1,
        page_size: 10
      });
      
      if (response && response.code === 200 && response.data) {
        setVehicles(response.data.vehicles || []);
      }
    } catch (err) {
      console.error('获取车辆列表失败:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchVehicles();
  }, []);

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: 60,
    },
    {
      title: '品牌型号',
      key: 'vehicle',
      width: 120,
      render: (_, record) => `${record.brand} ${record.style}`,
    },
    {
      title: '原始图片字符串',
      dataIndex: 'images',
      key: 'images',
      width: 300,
      render: (images) => (
        <div style={{ wordBreak: 'break-all', fontSize: '12px' }}>
          {images || '无'}
        </div>
      ),
    },
    {
      title: '解析后的图片',
      key: 'parsedImages',
      width: 200,
      render: (_, record) => {
        const parsedImages = parseImages(record.images);
        return (
          <div>
            {parsedImages.map((url, index) => (
              <div key={index} style={{ marginBottom: 4, fontSize: '12px' }}>
                <Text code>{index + 1}: {url}</Text>
              </div>
            ))}
          </div>
        );
      },
    },
    {
      title: '第一张图片',
      key: 'firstImage',
      width: 150,
      render: (_, record) => {
        const firstImage = getFirstImage(record.images);
        return (
          <div>
            <Image
              src={firstImage}
              alt="第一张图片"
              width={100}
              height={60}
              style={{ objectFit: 'cover' }}
              onError={(e) => {
                console.log(`车辆 ${record.id} 图片加载失败:`, e.target.src);
              }}
            />
            <div style={{ fontSize: '10px', marginTop: 4 }}>
              {firstImage}
            </div>
          </div>
        );
      },
    },
  ];

  return (
    <div style={{ padding: '20px' }}>
      <Title level={2}>车辆图片调试页面</Title>
      
      <Alert
        message="调试信息"
        description="这个页面用于调试车辆图片显示问题。请打开浏览器控制台查看详细的调试信息。"
        type="info"
        style={{ marginBottom: 20 }}
      />

      <Space style={{ marginBottom: 16 }}>
        <Button type="primary" onClick={fetchVehicles} loading={loading}>
          刷新数据
        </Button>
      </Space>

      <Table
        columns={columns}
        dataSource={vehicles}
        rowKey="id"
        loading={loading}
        scroll={{ x: 1000 }}
        pagination={false}
      />
    </div>
  );
};

export default VehicleDebug;
