# MinIO 文件上傳功能使用說明

## 功能概述

本系統提供了兩種 MinIO 文件上傳方式：
1. **預簽名 URL 上傳** - 獲取預簽名 URL，客戶端直接上傳到 MinIO
2. **服務器端上傳** - 通過 API 服務器上傳文件到 MinIO

## API 接口

### 1. 獲取預簽名 URL

**接口地址：** `GET /minio/presigned-url`

**請求參數：**
```json
{
  "bucket": "my-bucket",
  "object_name": "test.jpg",
  "expires": 3600,
  "content_type": "image/jpeg"
}
```

**響應示例：**
```json
{
  "success": true,
  "data": {
    "success": true,
    "url": "https://minio-server:9000/my-bucket/test.jpg?...",
    "method": "PUT",
    "expires_at": "2024-01-01T12:00:00Z",
    "message": "获取预签名URL成功"
  }
}
```

### 2. 直接上傳文件

**接口地址：** `POST /minio/upload`

**請求方式：** `multipart/form-data`

**請求參數：**
- `file`: 要上傳的文件
- `bucket`: 存儲桶名稱（必填）
- `object_name`: 對象名稱（必填）
- `content_type`: 文件類型（可選，會自動檢測）

**響應示例：**
```json
{
  "success": true,
  "data": {
    "message": "文件上傳成功",
    "bucket": "my-bucket",
    "object_name": "test.jpg",
    "file_size": 1024,
    "file_url": "http://14.103.140.237:9000/my-bucket/test.jpg"
  }
}
```

## 使用方式

### 方式一：預簽名 URL 上傳（推薦）

1. **獲取預簽名 URL**
```javascript
const response = await fetch('/minio/presigned-url', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    bucket: 'my-bucket',
    object_name: 'test.jpg',
    expires: 3600,
    content_type: 'image/jpeg'
  })
});

const { url } = await response.json();
```

2. **使用預簽名 URL 上傳文件**
```javascript
const file = document.getElementById('fileInput').files[0];
await fetch(url, {
  method: 'PUT',
  body: file,
  headers: {
    'Content-Type': file.type
  }
});
```

### 方式二：服務器端上傳

```javascript
const formData = new FormData();
formData.append('file', file);
formData.append('bucket', 'my-bucket');
formData.append('object_name', 'test.jpg');

const response = await fetch('/minio/upload', {
  method: 'POST',
  body: formData
});

const result = await response.json();
console.log('上傳成功：', result.data.file_url);
```

## 配置說明

### MinIO 配置

在 `Common/appconfig/config.yaml` 中配置 MinIO：

```yaml
minio:
  endpoint: "14.103.140.237:9000"
  access_key_id: "your-access-key"
  secret_access_key: "your-secret-key"
  use_ssl: false
  bucket_name: "default-bucket"
```

### 初始化流程

1. 系統啟動時會自動初始化 MinIO 客戶端
2. 客戶端實例存儲在 `global.Minio` 中
3. 所有 MinIO 操作都通過 `Common/utils/minio.go` 中的工具方法進行

## 錯誤處理

### 常見錯誤

1. **MinIO 客戶端未初始化**
   - 錯誤信息：`MinIO 客戶端未初始化`
   - 解決方案：檢查 MinIO 配置和初始化代碼

2. **Bucket 不存在**
   - 錯誤信息：`檢查 bucket 是否存在失敗`
   - 解決方案：系統會自動創建不存在的 bucket

3. **文件上傳失敗**
   - 錯誤信息：`上傳文件失敗`
   - 解決方案：檢查網絡連接和 MinIO 服務狀態

### 錯誤響應格式

```json
{
  "success": false,
  "message": "錯誤描述"
}
```

## 安全考慮

1. **文件大小限制**：建議在客戶端和服務器端都設置文件大小限制
2. **文件類型驗證**：根據業務需求驗證允許的文件類型
3. **訪問權限**：確保 MinIO 的訪問權限配置正確
4. **預簽名 URL 過期時間**：根據安全需求設置合適的過期時間

## 性能優化

1. **使用預簽名 URL**：減少服務器負載，提高上傳性能
2. **分片上傳**：對於大文件，可以實現分片上傳
3. **並發上傳**：支持多個文件同時上傳
4. **CDN 加速**：可以配置 CDN 來加速文件訪問

## 部署注意事項

1. 確保 MinIO 服務正常運行
2. 檢查網絡連接和防火牆設置
3. 配置正確的 CORS 策略
4. 監控 MinIO 服務的資源使用情況 