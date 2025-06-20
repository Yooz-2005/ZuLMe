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

	// Lua脚本：INCR后检查是否为1，若是则设置过期时间（24小时 = 86400秒）
	script := `
		local count = redis.call('INCR', KEYS[1])
		if count == 1 then
			redis.call('EXPIRE', KEYS[1], ARGV[1])
		end
		return count
	`
	// 执行Lua脚本（参数1：过期时间秒数，24小时 = 86400秒）
	result, err := global.Rdb.Eval(ctx, script, []string{countKey}, int64(86400)).Result()

	if err != nil {
		return 0, fmt.Errorf("执行Lua脚本失败: %v", err)
	}

	return result.(int64), nil
}

// GetSMSCount 获取当前发送次数
func GetSMSCount(phone string, source string) (int64, error) {
	ctx := context.Background()
	countKey := fmt.Sprintf("sms:count:%s:%s", source, phone)

	result, err := global.Rdb.Get(ctx, countKey).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return 0, nil // 如果key不存在，返回0
		}
		return 0, fmt.Errorf("获取发送次数失败: %v", err)
	}

	count := int64(0)
	if result != "" {
		if n, err := fmt.Sscanf(result, "%d", &count); err != nil || n != 1 {
			return 0, fmt.Errorf("解析发送次数失败: %v", err)
		}
	}

	return count, nil
}
