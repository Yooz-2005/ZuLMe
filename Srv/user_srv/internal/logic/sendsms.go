package logic

import (
	"Common/appconfig"
	"errors"
	"log"
	"math/rand"
	"models/model_redis"
	"strconv"
	user "user_srv/proto_user"
)

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func SendCode(in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	log.Println("SendCode: Received request for phone:", in.Phone) // Log entry

	// 2. 生成验证码
	code := generateCode()

	//短信发送
	//_, err := pkg.SendSms(code, in.Phone)
	//if err != nil {
	//	log.Println("短信发送失败")
	//	return nil, errors.New("短信发送失败")
	//}
	//if *sms.Body.Code != "OK" {
	//	log.Println("短信发送失败")
	//	return nil, errors.New(*sms.Body.Message)
	//}

	//存储验证码（5分钟有效期）
	err := model_redis.SaveVerificationCode(in.Source, in.Phone, code)
	if err != nil {
		return nil, errors.New("验证码存储失败")
	}

	//增加发送次数（24小时内有效）
	count, err := model_redis.IncrementSMSCount(in.Phone, in.Source)
	if err != nil {
		return nil, errors.New("系统错误")
	}

	//检查是否超过限制5次
	if count > appconfig.ConfData.MaxSend.Count {
		return nil, errors.New("发送次数过多，请5分钟后再试")
	}
	return &user.SendCodeResponse{
		Message: "验证码发送成功",
	}, nil
}
