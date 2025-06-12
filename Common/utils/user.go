package utils

import "regexp"

func EncryptPhone(phone string) string {
	// 检查手机号格式是否为11位数字
	re := regexp.MustCompile(`^1\d{10}$`)
	if !re.MatchString(phone) {
		return phone // 非标准手机号格式，返回原始值
	}
	// 替换中间四位为星号
	return phone[:3] + "****" + phone[7:]
}
