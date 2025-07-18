/**
 * 图片处理工具函数
 */

// 默认车辆图片
export const DEFAULT_VEHICLE_IMAGE = '/images/default-car.jpg';

// 根据品牌获取默认图片
export const getDefaultImageByBrand = (brand) => {
  const brandLower = (brand || '').toLowerCase();

  // 根据品牌返回不同的默认图片 - 使用现有的图片文件
  if (brandLower.includes('奔驰') || brandLower.includes('mercedes') || brandLower.includes('amg')) {
    return '/images/my-car-b.jpg';  // 奔驰系列
  } else if (brandLower.includes('宝马') || brandLower.includes('bmw') || brandLower.includes('g11')) {
    return '/images/my-car-d.jpg';  // 宝马系列
  } else if (brandLower.includes('奥迪') || brandLower.includes('audi')) {
    return '/images/my-car-e.jpg';  // 奥迪系列
  } else if (brandLower.includes('兰博基尼') || brandLower.includes('lamborghini') || brandLower.includes('huracan')) {
    return '/images/my-car-a.jpg';  // 兰博基尼系列
  } else if (brandLower.includes('保时捷') || brandLower.includes('porsche') || brandLower.includes('718')) {
    return '/images/my-car-c.jpg';  // 保时捷系列
  } else if (brandLower.includes('power')) {
    return '/images/my-car-e.jpg';  // Power系列
  }

  return DEFAULT_VEHICLE_IMAGE;
};

/**
 * 解析图片字符串为图片数组
 * @param {string} imagesString - 逗号分隔的图片URL字符串
 * @param {string} brand - 车辆品牌（用于选择默认图片）
 * @returns {Array} 图片URL数组
 */
export const parseImages = (imagesString, brand = '') => {
  console.log('解析图片字符串:', imagesString, '品牌:', brand);

  if (!imagesString || typeof imagesString !== 'string') {
    const defaultImg = getDefaultImageByBrand(brand);
    console.log('图片字符串为空或无效，使用默认图片:', defaultImg);
    return [defaultImg];
  }

  // 分割字符串并过滤空值（支持中文逗号和英文逗号）
  const images = imagesString
    .split(/[,，]/)  // 支持中文逗号和英文逗号
    .map(url => url.trim())
    .filter(url => url.length > 0)
    .map(url => {
      // 如果已经是完整的URL（http或https开头），直接返回
      if (url.startsWith('http://') || url.startsWith('https://')) {
        return url;
      }

      // 如果是相对路径，拼接基础URL
      const baseUrl = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8888';
      if (url.startsWith('/')) {
        return `${baseUrl}${url}`;
      }

      // 其他情况，拼接基础URL和路径
      return `${baseUrl}/${url}`;
    });

  console.log('解析后的图片数组:', images);

  // 如果没有有效图片，返回默认图片
  if (images.length === 0) {
    const defaultImg = getDefaultImageByBrand(brand);
    console.log('没有有效图片，使用默认图片:', defaultImg);
    return [defaultImg];
  }

  return images;
};

/**
 * 获取车辆的第一张图片（用于列表展示）
 * @param {string} imagesString - 逗号分隔的图片URL字符串
 * @param {string} brand - 车辆品牌（用于选择默认图片）
 * @returns {string} 第一张图片URL
 */
export const getFirstImage = (imagesString, brand = '') => {
  const images = parseImages(imagesString, brand);
  return images[0];
};

/**
 * 获取车辆的所有图片（用于详情页展示）
 * @param {string} imagesString - 逗号分隔的图片URL字符串
 * @param {string} brand - 车辆品牌（用于选择默认图片）
 * @returns {Array} 所有图片URL数组
 */
export const getAllImages = (imagesString, brand = '') => {
  return parseImages(imagesString, brand);
};

/**
 * 验证图片URL是否有效
 * @param {string} url - 图片URL
 * @returns {Promise<boolean>} 是否有效
 */
export const validateImageUrl = (url) => {
  return new Promise((resolve) => {
    const img = new Image();
    img.onload = () => resolve(true);
    img.onerror = () => resolve(false);
    img.src = url;
  });
};

/**
 * 图片加载错误处理
 * @param {Event} event - 错误事件
 * @param {string} brand - 车辆品牌（用于选择默认图片）
 */
export const handleImageError = (event, brand = '') => {
  const failedUrl = event.target.src;
  const defaultImg = getDefaultImageByBrand(brand);

  console.log('图片加载失败:', failedUrl);
  console.log('品牌:', brand);
  console.log('替换为默认图片:', defaultImg);

  // 避免无限循环：如果默认图片也加载失败，就不再替换
  if (failedUrl !== defaultImg && failedUrl !== DEFAULT_VEHICLE_IMAGE) {
    event.target.src = defaultImg;
  } else {
    console.log('默认图片也加载失败，停止替换');
  }
};

/**
 * 预加载图片
 * @param {Array} imageUrls - 图片URL数组
 * @returns {Promise<Array>} 预加载结果
 */
export const preloadImages = (imageUrls) => {
  const promises = imageUrls.map(url => {
    return new Promise((resolve) => {
      const img = new Image();
      img.onload = () => resolve({ url, success: true });
      img.onerror = () => resolve({ url, success: false });
      img.src = url;
    });
  });

  return Promise.all(promises);
};

/**
 * 处理评论图片URL
 * @param {string} imageUrl - 原始图片URL
 * @returns {string} - 处理后的图片URL
 */
export const processCommentImageUrl = (imageUrl) => {
  if (!imageUrl) {
    return getCommentDefaultImage();
  }

  // 如果已经是完整的URL（http或https开头），直接返回
  if (imageUrl.startsWith('http://') || imageUrl.startsWith('https://')) {
    return imageUrl;
  }

  // 如果是临时图片标识符，返回默认图片
  if (imageUrl.startsWith('temp_image_') || imageUrl.startsWith('uploaded_')) {
    return getCommentDefaultImage();
  }

  // 如果是相对路径，拼接基础URL
  const baseUrl = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8888';
  if (imageUrl.startsWith('/')) {
    return `${baseUrl}${imageUrl}`;
  }

  // 其他情况，拼接基础URL和路径
  return `${baseUrl}/${imageUrl}`;
};

/**
 * 获取评论默认图片
 * @returns {string} - 默认图片的base64编码
 */
export const getCommentDefaultImage = () => {
  return "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iODAiIGhlaWdodD0iODAiIHZpZXdCb3g9IjAgMCA4MCA4MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3Qgd2lkdGg9IjgwIiBoZWlnaHQ9IjgwIiBmaWxsPSIjRjVGNUY1Ii8+CjxwYXRoIGQ9Ik0yNCAzMkM0Ni4wOTE0IDMyIDY0IDQ5LjkwODYgNjQgNzJINzJDNzIgNDUuNDkwMyA1MC41MDk3IDI0IDI0IDI0VjMyWiIgZmlsbD0iI0Q5RDlEOSIvPgo8cGF0aCBkPSJNMjQgNDBDNDEuNjczIDQwIDU2IDU0LjMyNyA1NiA3Mkg2NEM2NCA0OS45MDg2IDQ2LjA5MTQgMzIgMjQgMzJWNDBaIiBmaWxsPSIjQkZCRkJGIi8+CjxjaXJjbGUgY3g9IjI0IiBjeT0iNzIiIHI9IjgiIGZpbGw9IiNEOUQ5RDkiLz4KPC9zdmc+";
};

/**
 * 处理评论图片数组
 * @param {Array} images - 图片URL数组
 * @returns {Array} - 处理后的图片URL数组
 */
export const processCommentImages = (images) => {
  if (!Array.isArray(images)) {
    return [];
  }

  return images.map(processCommentImageUrl).filter(Boolean);
};
