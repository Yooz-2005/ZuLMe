import React, { useState, useEffect } from 'react';
import { 
  Card, 
  List, 
  Rate, 
  Tag, 
  Image, 
  Typography, 
  Empty, 
  Spin, 
  message, 
  Row, 
  Col,
  Divider,
  Button,
  Modal
} from 'antd';
import {
  StarOutlined,
  CarOutlined,
  CustomerServiceOutlined,
  EditOutlined,
  DeleteOutlined,
  CheckCircleOutlined
} from '@ant-design/icons';
import commentService from '../../services/commentService';
import CommentModal from '../CommentModal';
import './index.css';

const { Text, Paragraph } = Typography;

const UserCommentList = () => {
  const [comments, setComments] = useState([]);
  const [loading, setLoading] = useState(false);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingComment, setEditingComment] = useState(null);

  // 获取用户评价列表
  const fetchUserComments = async (page = 1, pageSize = 10) => {
    setLoading(true);
    try {
      const userPhone = localStorage.getItem('userPhone');
      if (!userPhone) {
        message.error('请先登录');
        return;
      }

      // 这里需要根据实际API调整，暂时使用模拟数据
      const response = await commentService.getUserComments({
        page,
        page_size: pageSize
      });

      if (response && response.code === 200) {
        const commentsData = response.data || [];
        console.log('获取到的评价数据:', commentsData); // 调试日志
        setComments(commentsData);
        setPagination({
          current: page,
          pageSize: pageSize,
          total: response.total || 0
        });
      } else {
        message.error(response?.message || '获取评价列表失败');
      }
    } catch (error) {
      console.error('获取评价列表失败:', error);
      message.error('获取评价列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUserComments();
  }, []);

  // 处理分页变化
  const handlePageChange = (page, pageSize) => {
    fetchUserComments(page, pageSize);
  };

  // 编辑评价
  const handleEditComment = (comment) => {
    setEditingComment(comment);
    setEditModalVisible(true);
  };

  // 删除评价
  const handleDeleteComment = (commentId) => {
    Modal.confirm({
      title: '确认删除',
      content: '确定要删除这条评价吗？删除后无法恢复。',
      okText: '确定',
      cancelText: '取消',
      onOk: async () => {
        try {
          const response = await commentService.deleteComment(commentId);
          if (response && response.code === 200) {
            message.success('删除成功');
            fetchUserComments(pagination.current, pagination.pageSize);
          } else {
            message.error(response?.message || '删除失败');
          }
        } catch (error) {
          console.error('删除评价失败:', error);
          message.error('删除失败');
        }
      }
    });
  };

  // 编辑成功回调
  const handleEditSuccess = () => {
    setEditModalVisible(false);
    setEditingComment(null);
    fetchUserComments(pagination.current, pagination.pageSize);
  };

  // 渲染评分项
  const renderRatingItem = (label, value, icon) => (
    <div style={{ display: 'flex', alignItems: 'center', marginBottom: 8 }}>
      {icon}
      <Text style={{ marginLeft: 8, marginRight: 16, minWidth: 80 }}>{label}:</Text>
      <Rate disabled value={value} style={{ fontSize: 14 }} />
      <Text style={{ marginLeft: 8, color: '#666' }}>({value}分)</Text>
    </div>
  );

  // 渲染评价项
  const renderCommentItem = (item) => (
    <List.Item key={item.id}>
      <Card 
        style={{ width: '100%' }}
        actions={[
          <Button 
            type="link" 
            icon={<EditOutlined />} 
            onClick={() => handleEditComment(item)}
          >
            编辑
          </Button>,
          <Button 
            type="link" 
            danger 
            icon={<DeleteOutlined />} 
            onClick={() => handleDeleteComment(item.id)}
          >
            删除
          </Button>
        ]}
      >
        <Row gutter={16}>
          <Col span={16}>
            <div style={{ marginBottom: 16 }}>
              <Text strong style={{ fontSize: 16 }}>订单号: {item.order_sn || item.order_id}</Text>
              <Tag color="blue" style={{ marginLeft: 8 }}>
                {item.vehicle_name || '车辆信息'}
              </Tag>
              {item.is_anonymous && <Tag color="orange">匿名评价</Tag>}
            </div>

            <div style={{ marginBottom: 16 }}>
              <div style={{ display: 'flex', alignItems: 'center', marginBottom: 12 }}>
                <StarOutlined style={{ color: '#faad14', marginRight: 8 }} />
                <Text strong>总体评分:</Text>
                <Rate disabled value={item.rating} style={{ marginLeft: 8, fontSize: 16 }} />
                <Text style={{ marginLeft: 8, color: '#666' }}>({item.rating}分)</Text>
              </div>

              <Row gutter={16}>
                <Col span={8}>
                  {renderRatingItem('服务评分', item.service_rating, <CustomerServiceOutlined style={{ color: '#52c41a' }} />)}
                </Col>
                <Col span={8}>
                  {renderRatingItem('车辆评分', item.vehicle_rating, <CarOutlined style={{ color: '#1890ff' }} />)}
                </Col>
                <Col span={8}>
                  {renderRatingItem('清洁评分', item.clean_rating, <CheckCircleOutlined style={{ color: '#722ed1' }} />)}
                </Col>
              </Row>
            </div>

            <div style={{ marginBottom: 16 }}>
              <Text strong>评价内容:</Text>
              <Paragraph style={{ marginTop: 8, marginBottom: 0 }}>
                {item.content}
              </Paragraph>
            </div>

            {item.images && item.images.length > 0 && (
              <div style={{ marginBottom: 16 }}>
                <Text strong>评价图片:</Text>
                <div style={{ marginTop: 8 }}>
                  <Image.PreviewGroup>
                    {item.images.map((imageUrl, index) => {
                      // 处理图片URL
                      let processedUrl = imageUrl;

                      // 如果是临时图片标识符，显示占位符
                      if (imageUrl.startsWith('temp_image_') || imageUrl.startsWith('uploaded_')) {
                        return (
                          <div
                            key={index}
                            style={{
                              display: 'inline-block',
                              width: 80,
                              height: 80,
                              marginRight: 8,
                              backgroundColor: '#f5f5f5',
                              border: '1px solid #d9d9d9',
                              borderRadius: 4,
                              textAlign: 'center',
                              lineHeight: '80px',
                              fontSize: '12px',
                              color: '#999'
                            }}
                          >
                            图片{index + 1}
                          </div>
                        );
                      }

                      // 如果不是完整URL，尝试拼接基础URL
                      if (!imageUrl.startsWith('http://') && !imageUrl.startsWith('https://')) {
                        const baseUrl = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8888';
                        processedUrl = imageUrl.startsWith('/') ? `${baseUrl}${imageUrl}` : `${baseUrl}/${imageUrl}`;
                      }

                      return (
                        <Image
                          key={index}
                          width={80}
                          height={80}
                          src={processedUrl}
                          style={{ marginRight: 8, objectFit: 'cover', borderRadius: 4 }}
                          fallback="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iODAiIGhlaWdodD0iODAiIHZpZXdCb3g9IjAgMCA4MCA4MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3Qgd2lkdGg9IjgwIiBoZWlnaHQ9IjgwIiBmaWxsPSIjRjVGNUY1Ii8+CjxwYXRoIGQ9Ik0yNCAzMkM0Ni4wOTE0IDMyIDY0IDQ5LjkwODYgNjQgNzJINzJDNzIgNDUuNDkwMyA1MC41MDk3IDI0IDI0IDI0VjMyWiIgZmlsbD0iI0Q5RDlEOSIvPgo8cGF0aCBkPSJNMjQgNDBDNDEuNjczIDQwIDU2IDU0LjMyNyA1NiA3Mkg2NEM2NCA0OS45MDg2IDQ2LjA5MTQgMzIgMjQgMzJWNDBaIiBmaWxsPSIjQkZCRkJGIi8+CjxjaXJjbGUgY3g9IjI0IiBjeT0iNzIiIHI9IjgiIGZpbGw9IiNEOUQ5RDkiLz4KPC9zdmc+"
                          onError={() => {
                            console.log('图片加载失败:', processedUrl);
                            console.log('原始URL:', imageUrl);
                          }}
                        />
                      );
                    })}
                  </Image.PreviewGroup>
                  <div style={{ marginTop: 4, fontSize: '12px', color: '#999' }}>
                    原始图片数据: {JSON.stringify(item.images)}
                  </div>
                </div>
              </div>
            )}

            {item.reply_content && (
              <div style={{ 
                backgroundColor: '#f6f6f6', 
                padding: 12, 
                borderRadius: 4,
                marginTop: 16 
              }}>
                <Text strong style={{ color: '#1890ff' }}>商家回复:</Text>
                <Paragraph style={{ marginTop: 8, marginBottom: 0 }}>
                  {item.reply_content}
                </Paragraph>
                {item.reply_time && (
                  <Text type="secondary" style={{ fontSize: 12 }}>
                    回复时间: {item.reply_time}
                  </Text>
                )}
              </div>
            )}
          </Col>

          <Col span={8}>
            <div style={{ textAlign: 'right' }}>
              <Text type="secondary">评价时间</Text>
              <br />
              <Text type="secondary" style={{ fontSize: 12 }}>
                {item.created_at}
              </Text>
            </div>
          </Col>
        </Row>
      </Card>
    </List.Item>
  );

  return (
    <div className="user-comment-list">
      <Spin spinning={loading}>
        {comments.length > 0 ? (
          <List
            dataSource={comments}
            renderItem={renderCommentItem}
            pagination={{
              ...pagination,
              onChange: handlePageChange,
              showSizeChanger: true,
              showQuickJumper: true,
              showTotal: (total, range) => 
                `第 ${range[0]}-${range[1]} 条/共 ${total} 条`,
            }}
          />
        ) : (
          <Empty 
            description="暂无评价记录"
            image={Empty.PRESENTED_IMAGE_SIMPLE}
          />
        )}
      </Spin>

      {/* 编辑评价弹窗 */}
      {editModalVisible && editingComment && (
        <CommentModal
          visible={editModalVisible}
          onCancel={() => {
            setEditModalVisible(false);
            setEditingComment(null);
          }}
          onSuccess={handleEditSuccess}
          orderInfo={{
            id: editingComment.order_id,
            order_sn: editingComment.order_sn || editingComment.order_id,
            // 这里可以添加更多订单信息
          }}
          editMode={true}
          commentId={editingComment.id}
          initialValues={{
            rating: editingComment.rating,
            content: editingComment.content,
            service_rating: editingComment.service_rating,
            vehicle_rating: editingComment.vehicle_rating,
            clean_rating: editingComment.clean_rating,
            is_anonymous: editingComment.is_anonymous,
            images: editingComment.images || []
          }}
        />
      )}
    </div>
  );
};

export default UserCommentList;
