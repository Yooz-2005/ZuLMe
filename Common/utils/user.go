package utils

import (
	"fmt"
	"regexp"
	"time"
)

func EncryptPhone(phone string) string {
	// 检查手机号格式是否为11位数字
	re := regexp.MustCompile(`^1\d{10}$`)
	if !re.MatchString(phone) {
		return phone // 非标准手机号格式，返回原始值
	}
	// 替换中间四位为星号
	return phone[:3] + "****" + phone[7:]
}

// GenerateTaxNumber 根據當前時間和用戶ID生成唯一納稅人識別號，格式為 ZuLMe+時間戳+用戶ID
func GenerateTaxNumber(userID int64) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("ZuLMe%d%d", timestamp, userID)
}
