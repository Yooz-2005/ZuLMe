package model_redis

import (
	"Common/global"
	"context"
	"time"
)

// SaveVerificationCode 保存驗證碼到 Redis
func SaveVerificationCode(phone, code string) error {
	ctx := context.Background()
	// 設置 5 分鐘過期時間
	return global.Rdb.Set(ctx, "sms:"+phone, code, 5*time.Minute).Err()
}

// GetVerificationCode 從 Redis 獲取驗證碼
func GetVerificationCode(phone string) (string, error) {
	ctx := context.Background()
	return global.Rdb.Get(ctx, "sms:"+phone).Result()
}

// DeleteVerificationCode 刪除驗證碼
func DeleteVerificationCode(phone string) error {
	ctx := context.Background()
<<<<<<< HEAD
	return global.Rdb.Del(ctx, "sms:"+phone).Err()
=======
	key := source + phone
	return global.Rdb.Del(ctx, key).Err()
}

func IncrementSMSCount(phone string, source string) (int64, error) {
	ctx := context.Background()
	countKey := fmt.Sprintf("sms:count:%s:%s", source, phone)

	// Lua脚本：INCR后检查是否为1，若是则设置过期时间
	script := `
		local count = redis.call('INCR', KEYS[1])
		if count == 1 then
			redis.call('EXPIRE', KEYS[1], ARGV[1])
		end
		return count
	`
	// 执行Lua脚本（参数1：过期时间）
	result, err := global.Rdb.Eval(ctx, script, []string{countKey}, int64(5)).Result()
	if err != nil {
		return 0, fmt.Errorf("执行Lua脚本失败: %v", err)
	}
	return result.(int64), nil
>>>>>>> 7a8d98483f5c45085a094626c26fda6d517f3e37
}
