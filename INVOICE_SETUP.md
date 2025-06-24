# 发票功能设置指南

## 功能概述

ZuLMe 租车系统现已支持电子发票功能，用户可以为已支付的订单申请开具发票。

## 主要功能

1. **发票申请条件**：
   - ✅ 订单必须已支付（status = 2）
   - ✅ 订单必须属于当前用户
   - ✅ 订单不能重复开票

2. **发票生成**：
   - 自动生成发票号码
   - 生成PDF格式的电子发票
   - 支持在线下载

3. **发票内容**：
   - 购买方信息（用户真实姓名、证件号码）
   - 销售方信息（商家名称、税号）
   - 订单详情（车辆信息、租期、金额等）
   - 发票号码、开票日期等

## PDF生成工具安装

### Windows 系统

1. **下载 wkhtmltopdf**：
   - 访问：https://wkhtmltopdf.org/downloads.html
   - 下载 Windows 版本（推荐 64-bit）

2. **安装步骤**：
   ```bash
   # 下载后运行安装程序
   # 默认安装路径：C:\Program Files\wkhtmltopdf
   ```

3. **添加到系统PATH**：
   - 右键"此电脑" → "属性" → "高级系统设置"
   - 点击"环境变量"
   - 在"系统变量"中找到"Path"，点击"编辑"
   - 添加：`C:\Program Files\wkhtmltopdf\bin`
   - 重启命令行/IDE

4. **验证安装**：
   ```bash
   wkhtmltopdf --version
   ```

### 替代方案

如果无法安装 wkhtmltopdf，系统会自动生成HTML格式的发票文件作为替代。

## API 接口

### 用户申请发票
```http
POST /invoice/apply
Authorization: Bearer <user_token>
Content-Type: application/json

{
  "order_id": 123
}
```

### 响应格式
```json
{
  "code": 200,
  "message": "发票生成成功",
  "data": {
    "invoice_id": 1,
    "invoice_no": "INV1703123456789",
    "pdf_url": "http://localhost:8888/invoices/INV1703123456789.pdf"
  }
}
```

## 前端集成

发票按钮已集成到订单列表中，支持：
- 自动检查订单状态
- 一键申请发票
- 自动下载PDF文件

## 文件存储

- PDF文件存储在：`./invoices/` 目录
- 通过静态文件服务提供下载：`http://localhost:8888/invoices/`

## 故障排除

1. **PDF生成失败**：
   - 检查 wkhtmltopdf 是否正确安装
   - 确认系统PATH配置
   - 查看错误日志

2. **文件下载失败**：
   - 确认 invoices 目录存在
   - 检查文件权限
   - 验证静态文件服务配置

3. **发票重复申请**：
   - 系统会自动检查重复申请
   - 每个订单只能开具一次发票

## 注意事项

- 发票功能需要用户完善个人资料（真实姓名、证件号码）
- 商家信息需要在系统中正确配置
- PDF生成需要一定时间，请耐心等待
