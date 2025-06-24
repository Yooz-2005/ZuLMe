package constants

// 支付方式常量
const (
	PaymentMethodAlipay = 1 // 支付宝
	PaymentMethodWechat = 2 // 微信支付
)

// PaymentMethodNames 支付方式名称映射
var PaymentMethodNames = map[int32]string{
	PaymentMethodAlipay: "支付宝",
	PaymentMethodWechat: "微信支付",
}

// PaymentMethodCodes 支付方式代码映射
var PaymentMethodCodes = map[string]int32{
	"alipay": PaymentMethodAlipay,
	"wechat": PaymentMethodWechat,
}

// IsValidPaymentMethod 验证支付方式是否有效
func IsValidPaymentMethod(method int32) bool {
	return method == PaymentMethodAlipay || method == PaymentMethodWechat
}

// GetPaymentMethodName 获取支付方式名称
func GetPaymentMethodName(method int32) string {
	if name, exists := PaymentMethodNames[method]; exists {
		return name
	}
	return "未知支付方式"
}

// GetPaymentMethodCode 根据字符串获取支付方式代码
func GetPaymentMethodCode(methodStr string) (int32, bool) {
	code, exists := PaymentMethodCodes[methodStr]
	return code, exists
}
