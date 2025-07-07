package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// CustomTime 自定义时间类型，支持多种格式解析
type CustomTime struct {
	time.Time
}

// 支持的时间格式
var timeFormats = []string{
	"2006-01-02T15:04:05Z07:00", // RFC3339
	"2006-01-02T15:04:05Z",      // RFC3339 UTC
	"2006-01-02T15:04:05",       // ISO8601 without timezone
	"2006-01-02 15:04:05",       // MySQL datetime
	"2006-01-02",                // Date only
	"15:04:05",                  // Time only
}

// UnmarshalJSON 实现JSON反序列化
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		ct.Time = time.Time{}
		return nil
	}

	// 尝试各种时间格式
	for _, format := range timeFormats {
		if t, err := time.Parse(format, str); err == nil {
			ct.Time = t
			return nil
		}
	}

	return fmt.Errorf("无法解析时间格式: %s", str)
}

// MarshalJSON 实现JSON序列化
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + ct.Time.Format("2006-01-02T15:04:05Z07:00") + `"`), nil
}

// Value 实现driver.Valuer接口，用于数据库存储
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

// Scan 实现sql.Scanner接口，用于数据库读取
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		ct.Time = v
	case string:
		// 尝试解析字符串时间
		for _, format := range timeFormats {
			if t, err := time.Parse(format, v); err == nil {
				ct.Time = t
				return nil
			}
		}
		return fmt.Errorf("无法解析时间字符串: %s", v)
	default:
		return fmt.Errorf("无法将 %T 转换为时间", value)
	}

	return nil
}

// String 返回格式化的时间字符串
func (ct CustomTime) String() string {
	if ct.Time.IsZero() {
		return ""
	}
	return ct.Time.Format("2006-01-02 15:04:05")
}

// IsZero 检查时间是否为零值
func (ct CustomTime) IsZero() bool {
	return ct.Time.IsZero()
}

// Before 检查时间是否在指定时间之前
func (ct CustomTime) Before(t time.Time) bool {
	return ct.Time.Before(t)
}

// After 检查时间是否在指定时间之后
func (ct CustomTime) After(t time.Time) bool {
	return ct.Time.After(t)
}

// ParseTimeString 解析时间字符串，支持多种格式
func ParseTimeString(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	// 尝试各种时间格式
	for _, format := range timeFormats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("无法解析时间格式: %s", timeStr)
}

// FormatTime 格式化时间为标准格式
func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02T15:04:05Z07:00")
}

// FormatTimeForDB 格式化时间为数据库格式
func FormatTimeForDB(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// ParseFormTime 解析表单时间（支持datetime-local格式）
func ParseFormTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	// HTML datetime-local 格式: 2025-01-01T12:00
	if strings.Contains(timeStr, "T") && !strings.Contains(timeStr, "Z") && !strings.Contains(timeStr, "+") {
		// 添加时区信息
		timeStr += "Z"
	}

	return ParseTimeString(timeStr)
}

// NowFormatted 返回当前时间的格式化字符串
func NowFormatted() string {
	return FormatTime(time.Now())
}

// TodayStart 返回今天的开始时间
func TodayStart() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// TodayEnd 返回今天的结束时间
func TodayEnd() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
}
