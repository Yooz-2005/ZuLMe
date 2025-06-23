package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

// 生成指定位数的数字验证码
func GenerateNumericCode(length int) (string, error) {
	// 检查长度是否有效
	if length <= 0 {
		return "", nil
	}

	var code string
	for i := 0; i < length; i++ {
		// 使用crypto/rand生成真正随机的数字
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		// 将数字转换为字符串并添加到验证码中 把生成的随机数转换为对应的 ASCII 字符
		code += string(rune('0' + n.Int64()))
	}
	return code, nil
}

// GenerateTaxNumber 根据当前时间和用户ID生成唯一的纳税人识别号，格式为 ZuLMe+时间戳+用户ID
func GenerateTaxNumber(userID int64) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("ZuLMe%d%d", timestamp, userID)
}
