import React from 'react';
import { List, Avatar, Rate, Typography, Divider } from 'antd';
import { UserOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Text } = Typography;

const CommentItem = styled.div`
  padding: 16px 0;
  border-bottom: 1px solid #f0f0f0;

  .comment-header {
    display: flex;
    align-items: center;
    margin-bottom: 8px;
  }

  .comment-author {
    color: #1e293b;
    font-weight: 500;
    margin-right: 16px;
  }

  .comment-rating {
    display: flex;
    align-items: center;

    .ant-rate {
      font-size: 14px;
      margin-right: 8px;
    }
  }

  .comment-content {
    color: #475569;
    margin: 8px 0;
  }

  .comment-time {
    color: #94a3b8;
    font-size: 12px;
  }

  .comment-images {
    margin-top: 8px;
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }
`;

// 适配后端评论数据结构，增强健壮性
const adaptComments = (comments) => {
  if (!Array.isArray(comments)) return [];
  return comments.map(comment => {
    let images = [];
    if (typeof comment.images === 'string' && comment.images) {
      images = comment.images.split(',').map(url => url.trim()).filter(Boolean);
    } else if (Array.isArray(comment.images)) {
      images = comment.images;
    }

    // 处理用户名显示逻辑
    let displayName = '匿名用户';
    if (comment.is_anonymous) {
      displayName = '匿名用户';
    } else if (comment.user_name) {
      displayName = comment.user_name;
    } else if (comment.user_id) {
      displayName = `用户${comment.user_id}`;
    }

    return {
      ...comment,
      user_avatar: comment.user_avatar || '',
      user_name: displayName,
      rating: comment.rating || comment.vehicle_rating || comment.service_rating || 5,
      images,
      created_at: comment.created_at || comment.create_time || '',
      id: comment.id || comment._id || Math.random().toString(36).slice(2),
    };
  });
};

const CommentList = ({ comments = [], loading = false }) => {
  const adapted = adaptComments(comments);
  console.log('原始comments:', comments);
  console.log('适配后:', adapted);
  return (
    <div>
      <Divider orientation="left">用户评论</Divider>
      <List
        loading={loading}
        dataSource={adapted}
        locale={{
          emptyText: '暂无评论'
        }}
        renderItem={comment => (
          <CommentItem key={comment.id}>
            <div className="comment-header">
              <Avatar
                size={32}
                icon={<UserOutlined />}
                src={comment.user_avatar}
                style={{ marginRight: 12 }}
              />
              <Text className="comment-author">
                {comment.user_name || '匿名用户'}
              </Text>
              <div className="comment-rating">
                <Rate disabled defaultValue={comment.rating} />
                <Text type="secondary">{comment.rating}.0</Text>
              </div>
            </div>
            <div className="comment-content">
              <Text>{comment.content}</Text>
            </div>
            {comment.images && comment.images.length > 0 && (
              <div className="comment-images">
                {comment.images.map((image, index) => (
                  <img
                    key={index}
                    src={image}
                    alt={`评论图片 ${index + 1}`}
                    style={{
                      width: 100,
                      height: 100,
                      objectFit: 'cover',
                      borderRadius: 4
                    }}
                  />
                ))}
              </div>
            )}
            <div className="comment-time">
              <Text type="secondary">
                {comment.created_at ? new Date(comment.created_at).toLocaleString() : ''}
              </Text>
            </div>
          </CommentItem>
        )}
      />
    </div>
  );
};

export default CommentList; 