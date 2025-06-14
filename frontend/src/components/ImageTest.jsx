import React from 'react';
import { Card, Row, Col, Typography, Carousel } from 'antd';
import { getAllImages, getFirstImage } from '../utils/imageUtils';

const { Title, Text } = Typography;

const ImageTest = () => {
  // 测试数据 - 模拟数据库中的图片字符串
  const testImageString = "https://p3.dcarimg.com/img/motor-mis-img/a6c93acbda04bcc21dd02478188606f3~512x0.webp，https://p3.dcarimg.com/img/motor-mis-img/DCP_a2f31397befca287bf2024c852c73a33~512x0.webp，https://p3.dcarimg.com/img/motor-mis-img/a5bb164776efda381e6baf6c5172b071~~512x0.webp";
  
  const allImages = getAllImages(testImageString);
  const firstImage = getFirstImage(testImageString);

  return (
    <div style={{ padding: '20px' }}>
      <Title level={2}>图片解析测试</Title>
      
      <Row gutter={[16, 16]}>
        <Col span={12}>
          <Card title="第一张图片（列表页显示）">
            <img 
              src={firstImage} 
              alt="第一张图片" 
              style={{ width: '100%', height: '200px', objectFit: 'cover' }}
              onError={(e) => {
                e.target.src = '/images/my-car-a.jpg';
              }}
            />
            <Text>URL: {firstImage}</Text>
          </Card>
        </Col>
        
        <Col span={12}>
          <Card title="所有图片（详情页轮播）">
            <Carousel autoplay dots={true}>
              {allImages.map((imageUrl, index) => (
                <div key={index}>
                  <img
                    src={imageUrl}
                    alt={`图片 ${index + 1}`}
                    style={{ width: '100%', height: '200px', objectFit: 'cover' }}
                    onError={(e) => {
                      e.target.src = '/images/my-car-a.jpg';
                    }}
                  />
                </div>
              ))}
            </Carousel>
            <div style={{ marginTop: '10px' }}>
              <Text>共 {allImages.length} 张图片</Text>
              <ul>
                {allImages.map((url, index) => (
                  <li key={index}>
                    <Text code>{url}</Text>
                  </li>
                ))}
              </ul>
            </div>
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default ImageTest;
