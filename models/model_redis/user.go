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
	return global.Rdb.Del(ctx, "sms:"+phone).Err()
}
