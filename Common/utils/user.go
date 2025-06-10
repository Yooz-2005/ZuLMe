package utils

// 加密手机号
func EncryptPhone(phone string) string {
	if len(phone) < 11 {
		return phone
	}
	start := phone[0:3]
	end := phone[7:]
	return start + "****" + end
}
