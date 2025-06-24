import React from 'react';
import { List, Avatar, Rate, Typography, Divider, Image } from 'antd';
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

    // 处理图片数据
    if (typeof comment.images === 'string' && comment.images) {
      // 如果是字符串，尝试解析JSON或按逗号分割
      try {
        const parsed = JSON.parse(comment.images);
        if (Array.isArray(parsed)) {
          images = parsed;
        } else {
          images = comment.images.split(',').map(url => url.trim()).filter(Boolean);
        }
      } catch (e) {
        images = comment.images.split(',').map(url => url.trim()).filter(Boolean);
      }
    } else if (Array.isArray(comment.images)) {
      images = comment.images;
    }

    // 确保图片URL是完整的
    images = images.map(imageUrl => {
      if (!imageUrl) return '';

      // 如果已经是完整URL，直接返回
      if (imageUrl.startsWith('http://') || imageUrl.startsWith('https://')) {
        return imageUrl;
      }

      // 如果是相对路径，拼接基础URL
      const baseUrl = 'http://14.103.149.192:9000/zulme-06';
      if (imageUrl.startsWith('/')) {
        return `${baseUrl}${imageUrl}`;
      }

      // 其他情况，直接拼接
      return imageUrl.includes('zulme-06') ? `http://14.103.149.192:9000/${imageUrl}` : `${baseUrl}/${imageUrl}`;
    }).filter(Boolean);

    // 处理用户名显示逻辑
    let displayName = '匿名用户';
    if (comment.is_anonymous) {
      displayName = '匿名用户';
    } else if (comment.user_name) {
      displayName = comment.user_name;
    } else if (comment.user_id) {
      displayName = `用户${comment.user_id}`;
    }

    console.log('处理评论图片:', {
      原始images: comment.images,
      处理后images: images,
      commentId: comment.id
    });

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
  console.log('CommentList接收到的comments:', comments);
  console.log('适配后的comments:', adapted);

  // 如果没有评论数据，显示空状态
  if (!loading && (!adapted || adapted.length === 0)) {
    return (
      <div>
        <Divider orientation="left">用户评论</Divider>
        <div style={{ textAlign: 'center', padding: '40px 0' }}>
          <Text type="secondary">暂无评论</Text>
        </div>
      </div>
    );
  }

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
                <Image.PreviewGroup>
                  {comment.images.map((image, index) => (
                    <Image
                      key={index}
                      src={image}
                      alt={`评论图片 ${index + 1}`}
                      width={100}
                      height={100}
                      style={{
                        objectFit: 'cover',
                        borderRadius: 4,
                        marginRight: 8,
                        marginBottom: 8
                      }}
                      fallback="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgdmlld0JveD0iMCAwIDEwMCAxMDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIiBmaWxsPSIjRjVGNUY1Ii8+CjxwYXRoIGQ9Ik0zMCA0MEM1Ny42MTQyIDQwIDgwIDYyLjM4NTggODAgOTBIOTBDOTAgNTYuODYyOSA2My4xMzcxIDMwIDMwIDMwVjQwWiIgZmlsbD0iI0Q5RDlEOSIvPgo8cGF0aCBkPSJNMzAgNTBDNTIuMDkxNCA1MCA3MCA2Ny45MDg2IDcwIDkwSDgwQzgwIDYyLjM4NTggNTcuNjE0MiA0MCAzMCA0MFY1MFoiIGZpbGw9IiNCRkJGQkYiLz4KPGNpcmNsZSBjeD0iMzAiIGN5PSI5MCIgcj0iMTAiIGZpbGw9IiNEOUQ5RDkiLz4KPC9zdmc+"
                      preview={{
                        src: image
                      }}
                    />
                  ))}
                </Image.PreviewGroup>

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