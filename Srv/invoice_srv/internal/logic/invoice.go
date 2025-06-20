package logic

import (
	"Common/global"
	"Common/utils"
	"fmt"
	invoice "invoice_srv/proto_invoice"
	"models/model_mysql"
	"time"
)

// GenerateInvoice 生成发票
func GenerateInvoice(in *invoice.GenerateInvoiceRequest) (*invoice.GenerateInvoiceResponse, error) {
	// 1. 根据 OrderID 查找订单信息
	order := &model_mysql.Orders{}
	if err := order.GetByID(uint(in.OrderId)); err != nil {
		return &invoice.GenerateInvoiceResponse{Code: 404, Message: fmt.Errorf("查找订单失败: %v", err).Error()}, nil
	}

	// 2. 查找对应的车辆信息，用于获取车辆名称 (项目名称)
	vehicle := &model_mysql.Vehicle{}
	err := vehicle.GetByID(order.VehicleId)
	if err != nil {
		return &invoice.GenerateInvoiceResponse{Code: 404, Message: fmt.Errorf("查找车辆信息失败: %v", err).Error()}, nil
	}

	// 3. 查询购买方信息 (UserProfile)
	userProfile := &model_mysql.UserProfile{}
	if err := userProfile.GetByUserID(int64(order.UserId)); err != nil {
		// 如果找不到用户档案，可以根据业务需求选择报错或使用默认值
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: fmt.Errorf("查找购买方信息失败: %v", err).Error()}, nil
	}

	// 4. 查询销售方信息 (Merchant) - 使用 merchant_id 作为发票创建者/销售方
	merchant := &model_mysql.Merchant{}
	if err := merchant.GetByID(in.MerchantId); err != nil {
		// 如果找不到商家信息，可以根据业务需求选择报错或使用默认值
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: fmt.Errorf("查找销售方商家信息失败: %v", err).Error()}, nil
	}

	// 5. 构造 Invoice 对象
	newInvoice := &model_mysql.Invoice{
		OrderId:     int32(order.ID),
		MerchantId:  int32(in.MerchantId),                        // <-- 新增：赋值 MerchantId
		InvoiceNo:   fmt.Sprintf("INV%d", time.Now().UnixNano()), // 简单生成发票号码
		OrderSn:     order.OrderSn,
		Amount:      order.TotalAmount,
		IssuedAt:    time.Now(),
		Title:       in.InvoiceTitle,
		TaxNumber:   in.TaxNumber, // 使用请求中传入的税号作为发票的税号（销售方税号）
		InvoiceType: 1,            // Default to electronic invoice
		Status:      1,            // Initial status "待开"
		VehicleId:   int32(order.VehicleId),
		VehicleName: vehicle.Brand, // 使用车辆品牌作为项目名称
		RentalDays:  order.RentalDays,
		DailyRate:   order.DailyRate,
		PickupTime:  order.PickupTime,
		ReturnTime:  order.ReturnTime,
	}

	// 6. 保存发票信息到数据库
	if err := newInvoice.CreateInvoice(); err != nil {
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: fmt.Errorf("保存发票信息失败: %v", err).Error()}, nil
	}

	// 7. 生成 PDF
	// 传递购买方、销售方和开票人信息给 PDF 生成函数
	pdfPath, err := utils.GenerateInvoicePDF(
		newInvoice,
		order,
		userProfile.RealName, // 购买方名称
		userProfile.IdNumber, // 购买方税号 (使用证件号码)
		merchant.Name,        // 销售方名称 (从 merchant 表获取)
		newInvoice.TaxNumber, // 销售方税号 (从请求中传入的发票税号)
		merchant.Name,        // 开票人名称 (从 merchant 表获取，可根据需求修改)
	)

	if err != nil {
		// Attempt to rollback if PDF generation fails
		_ = global.DB.Delete(newInvoice).Error // Simple rollback, ignoring error for simplicity
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: "生成 PDF 失败: " + err.Error()}, nil
	}

	// 8. 更新发票状态和PDF URL
	newInvoice.Status = 2 // 2 代表"已开"
	newInvoice.PdfUrl = pdfPath
	if err := newInvoice.UpdateStatus(newInvoice.Status); err != nil {
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: "更新发票状态失败: " + err.Error()}, nil
	}
	if err := newInvoice.UpdatePDFUrl(newInvoice.PdfUrl); err != nil {
		return &invoice.GenerateInvoiceResponse{Code: 500, Message: "更新发票PDF URL失败: " + err.Error()}, nil
	}

	return &invoice.GenerateInvoiceResponse{
		Code:      200,
		Message:   "发票生成成功",
		InvoiceId: newInvoice.Id,
		InvoiceNo: newInvoice.InvoiceNo,
		PdfUrl:    pdfPath,
	}, nil
}
