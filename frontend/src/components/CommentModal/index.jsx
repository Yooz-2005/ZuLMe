import React, { useState, useEffect } from 'react';
import { Modal, Form, Rate, Input, Upload, Button, message, Row, Col, Divider } from 'antd';
import { PlusOutlined, StarOutlined } from '@ant-design/icons';
import commentService from '../../services/commentService';
import minioService from '../../services/minioService';
import { processCommentImageUrl } from '../../utils/imageUtils';
import './index.css';

const { TextArea } = Input;

const CommentModal = ({
  visible,
  onCancel,
  onSuccess,
  orderInfo,
  editMode = false,
  initialValues = null,
  commentId = null
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [fileList, setFileList] = useState([]);

  // 处理编辑模式的初始化
  useEffect(() => {
    if (editMode && initialValues) {
      form.setFieldsValue(initialValues);

      // 处理图片初始化
      if (initialValues.images && initialValues.images.length > 0) {
        const imageFileList = initialValues.images.map((url, index) => {
          const processedUrl = processCommentImageUrl(url);
          return {
            uid: `existing-${index}`,
            name: `image-${index}`,
            status: 'done',
            url: processedUrl,
            thumbUrl: processedUrl
          };
        });
        setFileList(imageFileList);
      }
    } else {
      // 重置表单和文件列表
      form.resetFields();
      setFileList([]);
    }
  }, [editMode, initialValues, form, visible]);

  // 处理文件上传
  const handleUploadChange = ({ fileList: newFileList }) => {
    setFileList(newFileList);
  };

  // 自定义上传函数
  const customUpload = async ({ file, onSuccess, onError, onProgress }) => {
    console.log('开始上传文件:', file.name, file.type, file.size);

    try {
      // 生成唯一的文件名
      const timestamp = Date.now();
      const randomStr = Math.random().toString(36).substring(2, 11);
      const fileExtension = file.name.split('.').pop();
      const objectName = `comments/${timestamp}_${randomStr}.${fileExtension}`;

      console.log('生成的对象名:', objectName);

      // 获取预签名URL
      console.log('正在获取预签名URL...');
      const presignedResponse = await minioService.getPresignedUrl({
        bucket: 'zulme-06', // 使用正确的bucket名称
        objectName: objectName,
        expires: 3600 // 1小时有效期
      });

      console.log('预签名URL响应:', presignedResponse);

      if (!presignedResponse.success) {
        throw new Error(presignedResponse.message || '获取上传链接失败');
      }

      // 使用预签名URL上传文件
      console.log('正在上传文件到:', presignedResponse.data.url);
      const uploadResponse = await fetch(presignedResponse.data.url, {
        method: 'PUT',
        body: file,
        headers: {
          'Content-Type': file.type,
        },
      });

      console.log('上传响应状态:', uploadResponse.status, uploadResponse.statusText);

      if (!uploadResponse.ok) {
        const errorText = await uploadResponse.text();
        console.error('上传失败响应:', errorText);
        throw new Error(`文件上传失败: ${uploadResponse.status} ${uploadResponse.statusText}`);
      }

      // 构建文件访问URL
      const fileUrl = `http://14.103.149.192:9000/zulme-06/${objectName}`;

      console.log('文件上传成功，访问URL:', fileUrl);

      // 标记上传成功，并保存文件URL
      onSuccess({
        url: fileUrl,
        objectName: objectName
      });

      message.success('图片上传成功');
    } catch (error) {
      console.error('上传失败详细信息:', error);
      onError(error);
      message.error('图片上传失败: ' + error.message);
    }
  };

  // 提交评论
  const handleSubmit = async (values) => {
    setLoading(true);
    try {
      // 处理图片：获取已上传的图片URL
      const images = fileList
        .filter(file => file.status === 'done' && file.response && file.response.url)
        .map(file => file.response.url);

      console.log('准备提交的图片数据:', images); // 调试日志

      const commentData = {
        order_id: orderInfo.id,
        rating: values.rating,
        content: values.content,
        service_rating: values.service_rating,
        vehicle_rating: values.vehicle_rating,
        clean_rating: values.clean_rating,
        images: images,
        is_anonymous: values.is_anonymous || false
      };

      console.log('准备提交的评论数据:', commentData); // 调试日志

      let response;
      if (editMode && commentId) {
        // 编辑模式
        response = await commentService.updateComment(commentId, commentData);
      } else {
        // 创建模式
        response = await commentService.createComment(commentData);
      }

      if (response && response.code === 200) {
        message.success(editMode ? '评价更新成功！' : '评价提交成功！');
        form.resetFields();
        setFileList([]);
        onSuccess && onSuccess();
        onCancel();
      } else {
        message.error(response?.message || (editMode ? '评价更新失败' : '评价提交失败'));
      }
    } catch (error) {
      console.error('提交评价失败:', error);
      message.error('评价提交失败，请重试');
    } finally {
      setLoading(false);
    }
  };

  // 上传前的检查
  const beforeUpload = (file) => {
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
      message.error('只能上传 JPG/PNG 格式的图片!');
      return Upload.LIST_IGNORE;
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('图片大小不能超过 2MB!');
      return Upload.LIST_IGNORE;
    }
    // 允许文件进入上传队列，但使用 customRequest 处理
    return true;
  };

  const uploadButton = (
    <div>
      <PlusOutlined />
      <div style={{ marginTop: 8 }}>上传图片</div>
    </div>
  );

  return (
    <Modal
      title={
        <div style={{ textAlign: 'center' }}>
          <StarOutlined style={{ color: '#faad14', marginRight: 8 }} />
          {editMode ? '编辑评价' : '评价订单'}
        </div>
      }
      open={visible}
      onCancel={onCancel}
      footer={null}
      width={600}
      className="comment-modal"
    >
      {orderInfo && (
        <div className="order-info-section">
          <h4>订单信息</h4>
          <p><strong>订单号:</strong> {orderInfo.order_sn}</p>
          <p><strong>租赁时间:</strong> {orderInfo.pickup_time} 至 {orderInfo.return_time}</p>
          <p><strong>租赁天数:</strong> {orderInfo.rental_days} 天</p>
          <Divider />
        </div>
      )}

      <Form
        form={form}
        layout="vertical"
        onFinish={handleSubmit}
        initialValues={{
          rating: 5,
          service_rating: 5,
          vehicle_rating: 5,
          clean_rating: 5
        }}
      >
        <Form.Item
          name="rating"
          label="总体评分"
          rules={[{ required: true, message: '请给出总体评分' }]}
        >
          <Rate allowHalf />
        </Form.Item>

        <Row gutter={16}>
          <Col span={8}>
            <Form.Item
              name="service_rating"
              label="服务评分"
              rules={[{ required: true, message: '请给出服务评分' }]}
            >
              <Rate allowHalf />
            </Form.Item>
          </Col>
          <Col span={8}>
            <Form.Item
              name="vehicle_rating"
              label="车辆评分"
              rules={[{ required: true, message: '请给出车辆评分' }]}
            >
              <Rate allowHalf />
            </Form.Item>
          </Col>
          <Col span={8}>
            <Form.Item
              name="clean_rating"
              label="清洁评分"
              rules={[{ required: true, message: '请给出清洁评分' }]}
            >
              <Rate allowHalf />
            </Form.Item>
          </Col>
        </Row>

        <Form.Item
          name="content"
          label="评价内容"
          rules={[
            { required: true, message: '请输入评价内容' },
            { min: 10, message: '评价内容至少10个字符' }
          ]}
        >
          <TextArea
            rows={4}
            placeholder="请分享您的租车体验，帮助其他用户做出更好的选择..."
            maxLength={500}
            showCount
          />
        </Form.Item>

        <Form.Item
          name="images"
          label="上传图片（可选）"
        >
          <Upload
            listType="picture-card"
            fileList={fileList}
            onChange={handleUploadChange}
            beforeUpload={beforeUpload}
            maxCount={5}
            customRequest={customUpload}
          >
            {fileList.length >= 5 ? null : uploadButton}
          </Upload>
        </Form.Item>

        <Form.Item style={{ textAlign: 'center', marginTop: 24 }}>
          <Button onClick={onCancel} style={{ marginRight: 16 }}>
            取消
          </Button>
          <Button type="primary" htmlType="submit" loading={loading}>
            {editMode ? '更新评价' : '提交评价'}
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CommentModal;
