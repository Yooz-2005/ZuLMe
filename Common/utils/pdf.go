package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"models/model_mysql"
	"os"
	"os/exec"
	"path/filepath"
)

// GenerateInvoicePDF 生成发票PDF (使用 wkhtmltopdf)
func GenerateInvoicePDF(invoice *model_mysql.Invoice, order *model_mysql.Orders, buyerName, buyerTaxID, sellerName, sellerTaxID, issuerName string) (string, error) {
	// 定义 HTML 模板
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="zh-CN">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>电子发票</title>
	    <style>
	        body { font-family: 'SimFang', 'Microsoft YaHei', sans-serif; margin: 0; padding: 20px; font-size: 10pt; }
	        .invoice-container { width: 210mm; margin: 0 auto; border: 1px solid #000; padding: 10px; box-sizing: border-box; }
	        .header { text-align: center; margin-bottom: 10px; position: relative; }
	        .header h1 { font-size: 20pt; margin: 0; padding: 5px 0; }
	        .invoice-type { font-size: 14pt; font-weight: bold; position: absolute; left: 50%; top: 50%; transform: translate(-50%, -50%); border: 2px solid red; border-radius: 50%; padding: 5px 10px; color: red; }
	        .header-info { display: flex; justify-content: space-between; font-size: 9pt; margin-bottom: 10px; }
	        .header-info div { flex: 1; text-align: left; }
	        .header-info div:last-child { text-align: right; }

	        .section-box { border: 1px solid #000; margin-bottom: 10px; }
	        .section-title { background-color: #f2f2f2; padding: 5px; font-weight: bold; border-bottom: 1px solid #000; }
	        .section-content { padding: 5px; display: flex; }
	        .section-half { flex: 1; padding: 0 5px; box-sizing: border-box; }
	        .section-half + .section-half { border-left: 1px solid #000; }

	        table { width: 100%; border-collapse: collapse; margin-bottom: 10px; }
	        table th, table td { border: 1px solid #000; padding: 4px; text-align: left; vertical-align: top; }
	        table th { background-color: #f2f2f2; font-weight: bold; text-align: center; }
	        .text-right { text-align: right; }
	        .text-center { text-align: center; }
	        .bold { font-weight: bold; }

	        .footer-notes { border: 1px solid #000; padding: 5px; font-size: 9pt; min-height: 50px; }
	        .footer-issuer { text-align: right; margin-top: 10px; font-size: 9pt; }
	    </style>
	</head>
	<body>
	    <div class="invoice-container">
	        <div class="header">
	            <!-- QR Code Placeholder (You would generate this dynamically or provide an image URL) -->
	            
	            <h1>电子发票 (普通发票)</h1>
	            <!-- Red Seal Placeholder (You would provide an image URL) -->
	            
	            <div class="header-info">
	                <div>发票号码: {{.InvoiceNo}}</div>
	                <div>开票日期: {{.IssuedAt}}</div>
	            </div>
	        </div>

	        <div class="section-box">
	            <div class="section-content">
	                <div class="section-half">
	                    <p class="bold">购买方信息</p>
	                    <p>名称: {{.BuyerName}}</p>
	                    <p>统一社会信用代码/纳税人识别号: {{.BuyerTaxID}}</p>
	                </div>
	                <div class="section-half">
	                    <p class="bold">销售方信息</p>
	                    <p>名称: {{.SellerName}}</p>
	                    <p>统一社会信用代码/纳税人识别号: {{.SellerTaxID}}</p>
	                </div>
	            </div>
	        </div>

	        <table>
	            <thead>
	                <tr>
	                    <th>项目名称</th>
	                    <th>规格型号</th>
	                    <th>单位</th>
	                    <th>数量</th>									
	                    <th>单价</th>
	                    <th>金额</th>
	                    <th>税率/征收率</th>
	                    <th>税额</th>
	                </tr>
	            </thead>
	            <tbody>
	                <tr>
	                    <td>租车服务</td>
	                    <td>{{.VehicleName}}</td>
	                    <td>天</td>
	                    <td class="text-right">{{.RentalDays}}</td>
	                    <td class="text-right">{{printf "%.2f" .DailyRate}}</td>
	                    <td class="text-right">{{printf "%.2f" .Amount}}</td>
	                    <td class="text-right">0%</td> <!-- Assuming 0% tax for simplicity -->
	                    <td class="text-right">0.00</td>
	                </tr>
	                <tr>
	                    <td colspan="5" class="text-right bold">合计</td>
	                    <td class="text-right bold">{{printf "%.2f" .Amount}}</td>
	                    <td colspan="2" class="text-right bold">0.00</td>
	                </tr>
	            </tbody>
	        </table>

	        <div class="footer-notes">
	            <p>备注: 订单号: {{.OrderSn}}, 总金额: {{printf "%.2f" .Amount}} 元</p>
	            <!-- You can add more detailed notes here if available in your data -->
	        </div>

	        <div class="footer-issuer">
	            <p>开票人: {{.IssuerName}}</p>
	        </div>
	    </div>
	</body>
	</html>
	`

	// 准备模板数据
	data := struct {
		InvoiceNo      string
		IssuedAt       string
		BuyerName      string
		BuyerTaxID     string
		SellerName     string
		SellerTaxID    string
		VehicleName    string
		RentalDays     int32
		DailyRate      float64
		Amount         float64
		OrderSn        string
		IssuerName     string
		TaxNumber      string // For consistency with original invoice model
		InvoiceTypeStr string // For consistency with original invoice model
		StatusStr      string // For consistency with original invoice model
		PickupTime     string
		ReturnTime     string
	}{
		InvoiceNo:      invoice.InvoiceNo,
		IssuedAt:       invoice.IssuedAt.Format("2006年01月02日"), // 匹配示例日期格式
		BuyerName:      buyerName,
		BuyerTaxID:     buyerTaxID,
		SellerName:     sellerName,
		SellerTaxID:    sellerTaxID,
		VehicleName:    invoice.VehicleName,
		RentalDays:     invoice.RentalDays,
		DailyRate:      invoice.DailyRate,
		Amount:         invoice.Amount,
		OrderSn:        order.OrderSn,
		IssuerName:     issuerName,
		TaxNumber:      invoice.TaxNumber,
		InvoiceTypeStr: getInvoiceTypeString(invoice.InvoiceType),
		StatusStr:      getInvoiceStatusString(invoice.Status),
		PickupTime:     invoice.PickupTime.Format("2006-01-02 15:04"),
		ReturnTime:     invoice.ReturnTime.Format("2006-01-02 15:04"),
	}

	// 解析并执行模板
	tmpl, err := template.New("invoice").Parse(htmlTemplate)
	if err != nil {
		return "", fmt.Errorf("解析 HTML 模板失败: %v", err)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", fmt.Errorf("执行 HTML 模板失败: %v", err)
	}

	// 将 HTML 内容写入临时文件
	htmlFileName := fmt.Sprintf("invoice_%s.html", invoice.InvoiceNo)
	htmlFilePath := filepath.Join(os.TempDir(), htmlFileName)
	// fmt.Printf("临时 HTML 文件路径: %s\n", htmlFilePath) // 调试用，可以打开
	if err := os.WriteFile(htmlFilePath, tpl.Bytes(), 0644); err != nil {
		return "", fmt.Errorf("写入临时 HTML 文件失败: %v", err)
	}

	// 确保 PDF 输出目录存在 - 使用绝对路径
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前工作目录失败: %v", err)
	}

	// 根据当前目录确定正确的invoices目录路径
	var outputDir string
	if filepath.Base(currentDir) == "invoice_srv" {
		// 如果在invoice_srv目录中运行，需要回到项目根目录
		outputDir = filepath.Join(filepath.Dir(filepath.Dir(currentDir)), "invoices")
	} else {
		// 如果在其他目录（如Api）中运行
		outputDir = filepath.Join(filepath.Dir(currentDir), "invoices")
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil { // 确保目录存在
		if !os.IsExist(err) {
			return "", fmt.Errorf("创建 PDF 输出目录失败: %v", err)
		}
		return "", fmt.Errorf("创建 PDF 输出目录失败: %v", err)
	}

	pdfFileName := fmt.Sprintf("%s.pdf", invoice.InvoiceNo)
	pdfPath := filepath.Join(outputDir, pdfFileName) // 确保路径正确

	// 调用 wkhtmltopdf 命令生成 PDF
	// 首先尝试系统PATH中的wkhtmltopdf，如果失败则尝试指定路径
	var cmd *exec.Cmd

	// 尝试使用系统PATH中的wkhtmltopdf
	cmd = exec.Command("wkhtmltopdf", "--page-size", "A4", "--encoding", "UTF-8", htmlFilePath, pdfPath)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		// 如果系统PATH中没有，尝试使用指定路径
		wkhtmltopdfPaths := []string{
			"D:\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",                // D盘安装路径
			"C:\\down\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",          // 原备用路径
			"C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe", // 默认安装路径
		}

		var lastErr error
		pdfCreated := false

		for _, wkhtmltopdfPath := range wkhtmltopdfPaths {
			cmd = exec.Command(wkhtmltopdfPath, "--page-size", "A4", "--encoding", "UTF-8", htmlFilePath, pdfPath)
			cmd.Stderr = os.Stderr
			lastErr = cmd.Run()
			if lastErr == nil {
				pdfCreated = true
				break
			}
		}

		if !pdfCreated {
			// 如果所有路径都失败，创建HTML文件作为替代
			fmt.Printf("警告: wkhtmltopdf 不可用，将创建HTML文件作为替代。请安装 wkhtmltopdf 以生成PDF文件。\n")
			fmt.Printf("尝试的路径: %v\n", wkhtmltopdfPaths)
			fmt.Printf("最后的错误: %v\n", lastErr)

			// 将HTML文件复制到输出目录作为替代
			htmlOutputPath := filepath.Join(outputDir, fmt.Sprintf("%s.html", invoice.InvoiceNo))
			if copyErr := os.WriteFile(htmlOutputPath, tpl.Bytes(), 0644); copyErr != nil {
				return "", fmt.Errorf("创建HTML文件失败: %v", copyErr)
			}
			pdfPath = htmlOutputPath
		}
	}

	// 清理临时 HTML 文件
	os.Remove(htmlFilePath)

	// 返回可访问的URL路径而不是本地文件路径
	fileName := filepath.Base(pdfPath)
	pdfURL := fmt.Sprintf("http://localhost:8888/invoices/%s", fileName)
	return pdfURL, nil
}

// 辅助函数：根据发票类型代码获取字符串描述
func getInvoiceTypeString(invoiceType int32) string {
	switch invoiceType {
	case 1:
		return "电子发票"
	case 2:
		return "纸质发票"
	default:
		return "未知类型"
	}
}

// 辅助函数：根据发票状态代码获取字符串描述
func getInvoiceStatusString(status int32) string {
	switch status {
	case 1:
		return "待开"
	case 2:
		return "已开"
	case 3:
		return "已作废"
	default:
		return "未知状态"
	}
}
