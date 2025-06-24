import React, { useState } from 'react';
import { Button, message, Modal } from 'antd';
import { FileTextOutlined, LoadingOutlined } from '@ant-design/icons';
import invoiceService from '../../services/invoiceService';

const InvoiceButton = ({ order, onInvoiceGenerated }) => {
  const [loading, setLoading] = useState(false);

  // 检查是否可以开发票
  const canGenerateInvoice = () => {
    // 只有已支付的订单才能开发票
    return order.status === 2; // 2 = 已支付
  };

  // 处理开发票
  const handleGenerateInvoice = async () => {
    if (!canGenerateInvoice()) {
      message.warning('只有已支付的订单才能开具发票');
      return;
    }

    Modal.confirm({
      title: '确认开具发票',
      content: `确定要为订单 ${order.order_sn || order.id} 开具发票吗？`,
      okText: '确认',
      cancelText: '取消',
      onOk: async () => {
        setLoading(true);
        try {
          const response = await invoiceService.applyInvoice(order.id);
          
          if (response && response.code === 200) {
            message.success('发票开具成功！');

            // 如果有PDF链接，可以打开下载
            // 注意：后端返回的数据结构可能是 response.data 或直接在 response 中
            const invoiceData = response.data || response;
            if (invoiceData && invoiceData.pdf_url) {
              const link = document.createElement('a');
              link.href = invoiceData.pdf_url;
              link.download = `发票_${invoiceData.invoice_no || order.order_sn}.pdf`;
              document.body.appendChild(link);
              link.click();
              document.body.removeChild(link);
            }

            // 通知父组件发票已生成
            if (onInvoiceGenerated) {
              onInvoiceGenerated(invoiceData);
            }
          } else {
            message.error(response?.message || '开具发票失败');
          }
        } catch (error) {
          console.error('开具发票失败:', error);
          message.error('开具发票失败，请稍后重试');
        } finally {
          setLoading(false);
        }
      }
    });
  };

  // 如果订单状态不允许开发票，显示禁用状态
  if (!canGenerateInvoice()) {
    return (
      <Button
        disabled
        size="small"
        icon={<FileTextOutlined />}
        title="只有已支付的订单才能开具发票"
      >
        开发票
      </Button>
    );
  }

  return (
    <Button
      type="primary"
      size="small"
      icon={loading ? <LoadingOutlined /> : <FileTextOutlined />}
      loading={loading}
      onClick={handleGenerateInvoice}
      style={{
        backgroundColor: '#52c41a',
        borderColor: '#52c41a'
      }}
    >
      开发票
    </Button>
  );
};

export default InvoiceButton;
