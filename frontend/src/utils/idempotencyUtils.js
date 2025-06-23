import { message, Modal } from 'antd';
import orderService from '../services/orderService';

/**
 * 幂等性检查工具
 * 确保用户在有未支付订单时不能创建新的预订或订单
 */

/**
 * 检查用户是否有未支付的订单
 * @returns {Promise<{hasUnpaidOrder: boolean, unpaidOrder?: Object}>}
 */
export const checkUserUnpaidOrder = async () => {
  try {
    const response = await orderService.checkUnpaidOrder();
    
    if (response && response.code === 200) {
      return {
        hasUnpaidOrder: response.data.has_unpaid_order,
        unpaidOrder: response.data.unpaid_order,
        message: response.data.message
      };
    } else {
      console.error('检查未支付订单失败:', response);
      return { hasUnpaidOrder: false };
    }
  } catch (error) {
    console.error('检查未支付订单异常:', error);
    // 网络错误时，为了安全起见，假设没有未支付订单
    return { hasUnpaidOrder: false };
  }
};

/**
 * 在执行预订或下单前进行幂等性检查
 * @param {Function} onProceed - 检查通过后执行的回调函数
 * @param {Function} onBlocked - 检查不通过时执行的回调函数（可选）
 * @returns {Promise<boolean>} 是否可以继续执行
 */
export const checkIdempotencyBeforeAction = async (onProceed, onBlocked) => {
  const checkResult = await checkUserUnpaidOrder();
  
  if (checkResult.hasUnpaidOrder) {
    const unpaidOrder = checkResult.unpaidOrder;
    
    // 显示未支付订单的提示弹窗
    Modal.confirm({
      title: '您有未完成的订单',
      content: (
        <div>
          <p>{checkResult.message}</p>
          <div style={{ marginTop: 16, padding: 12, backgroundColor: '#f6f6f6', borderRadius: 4 }}>
            <p><strong>订单号：</strong>{unpaidOrder.order_sn}</p>
            <p><strong>订单金额：</strong>¥{unpaidOrder.total_amount}</p>
            <p><strong>创建时间：</strong>{new Date(unpaidOrder.created_at).toLocaleString()}</p>
          </div>
          <p style={{ marginTop: 12, color: '#666' }}>
            请先完成当前订单的支付，然后再进行新的预订。
          </p>
        </div>
      ),
      okText: '去支付',
      cancelText: '取消',
      onOk: () => {
        // 跳转到支付页面
        if (unpaidOrder.payment_url) {
          window.open(unpaidOrder.payment_url, '_blank');
        } else {
          message.error('支付链接不可用，请联系客服');
        }
      },
      onCancel: () => {
        if (onBlocked) {
          onBlocked(checkResult);
        }
      }
    });
    
    return false; // 阻止继续执行
  } else {
    // 没有未支付订单，可以继续执行
    if (onProceed) {
      await onProceed();
    }
    return true;
  }
};

/**
 * 在预订前进行幂等性检查
 * @param {Function} reservationCallback - 预订回调函数
 * @returns {Promise<boolean>}
 */
export const checkBeforeReservation = async (reservationCallback) => {
  return await checkIdempotencyBeforeAction(
    reservationCallback,
    (checkResult) => {
      message.warning('请先完成当前订单的支付后再进行新的预订');
    }
  );
};

/**
 * 在下单前进行幂等性检查
 * @param {Function} orderCallback - 下单回调函数
 * @returns {Promise<boolean>}
 */
export const checkBeforeOrder = async (orderCallback) => {
  return await checkIdempotencyBeforeAction(
    orderCallback,
    (checkResult) => {
      message.warning('请先完成当前订单的支付后再创建新订单');
    }
  );
};

/**
 * 显示未支付订单提示
 * @param {Object} unpaidOrder - 未支付订单信息
 */
export const showUnpaidOrderNotification = (unpaidOrder) => {
  message.warning({
    content: `您有未完成支付的订单（${unpaidOrder.order_sn}），请先完成支付`,
    duration: 5,
    onClick: () => {
      if (unpaidOrder.payment_url) {
        window.open(unpaidOrder.payment_url, '_blank');
      }
    }
  });
};

/**
 * 检查并处理页面加载时的未支付订单
 * 可以在需要的页面组件中调用，提醒用户完成支付
 */
export const checkUnpaidOrderOnPageLoad = async () => {
  const checkResult = await checkUserUnpaidOrder();
  
  if (checkResult.hasUnpaidOrder) {
    // 延迟显示，避免与页面加载冲突
    setTimeout(() => {
      showUnpaidOrderNotification(checkResult.unpaidOrder);
    }, 2000);
  }
  
  return checkResult;
};
