package logic

import (
	"Common/appconfig"
	"Common/utils"
	"errors"
	"log"
	"models/model_redis"
	user "user_srv/proto_user"
)

func SendCode(in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	// 1. 检查当前发送次数是否超过限制
	currentCount, err := model_redis.GetSMSCount(in.Phone, in.Source)
	if err != nil {
		return nil, errors.New("系统错误")
	}
	if currentCount >= appconfig.ConfData.SendSms.Count {
		return nil, errors.New("发送次数过多，请24小时后再试")
	}

	// 2. 生成验证码
	code, err := utils.GenerateNumericCode(int(appconfig.ConfData.SendSms.Count))
	if err != nil {
		return nil, err
	}

	// 3. 存储验证码（5分钟有效期）
	err = model_redis.SaveVerificationCode(in.Source, in.Phone, code)
	if err != nil {
		return nil, errors.New("验证码存储失败")
	}

	// 4. 增加发送次数（24小时内有效）
	newCount, err := model_redis.IncrementSMSCount(in.Phone, in.Source)
	if err != nil {
		return nil, errors.New("系统错误")
	}
	log.Printf("SendCode: SMS count incremented for phone %s, new count: %d", in.Phone, newCount)

	return &user.SendCodeResponse{
		Message: "验证码发送成功",
	}, nil
}
