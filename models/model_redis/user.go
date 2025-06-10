package model_redis

import (
	"Common/global"
	"context"
	"fmt"
	"time"
)

// SaveVerificationCode 存储验证码到 Redis
func SaveVerificationCode(source, phone, code string) error {
	ctx := context.Background()
	key := source + phone
	return global.Rdb.Set(ctx, key, code, 5*time.Minute).Err()
}

// GetVerificationCode 从redis中获取验证码
func GetVerificationCode(source, phone string) (string, error) {
	ctx := context.Background()
	key := source + phone
	return global.Rdb.Get(ctx, key).Result()
}

// DeleteVerificationCode 删除验证码
func DeleteVerificationCode(source, phone string) error {
	ctx := context.Background()
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
	// 执行Lua脚本（参数1：过期时间秒数）
	result, err := global.Rdb.Eval(ctx, script, []string{countKey}, int64(5)).Result()

	if err != nil {
		return 0, fmt.Errorf("执行Lua脚本失败: %v", err)
	}

	return result.(int64), nil
}
