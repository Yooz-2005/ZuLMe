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

	// 1. 检查当前发送次数是否超过限制
	currentCount, err := model_redis.GetSMSCount(in.Phone, in.Source)
	if err != nil {
		log.Printf("SendCode: Failed to get SMS count for phone %s: %v", in.Phone, err)
		return nil, errors.New("系统错误")
	}

	log.Printf("SendCode: Current SMS count for phone %s: %d, limit: %d", in.Phone, currentCount, appconfig.ConfData.SendSms.Count)

	if currentCount >= appconfig.ConfData.SendSms.Count {
		log.Printf("SendCode: SMS limit exceeded for phone %s", in.Phone)
		return nil, errors.New("发送次数过多，请24小时后再试")
	}

	// 2. 生成验证码
	code := generateCode()
	log.Printf("SendCode: Generated verification code for phone %s: %s", in.Phone, code)

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

	// 3. 存储验证码（5分钟有效期）
	err = model_redis.SaveVerificationCode(in.Source, in.Phone, code)
	if err != nil {
		log.Printf("SendCode: Failed to save verification code for phone %s: %v", in.Phone, err)
		return nil, errors.New("验证码存储失败")
	}
	log.Printf("SendCode: Verification code saved for phone %s", in.Phone)

	// 4. 增加发送次数（24小时内有效）
	newCount, err := model_redis.IncrementSMSCount(in.Phone, in.Source)
	if err != nil {
		log.Printf("SendCode: Failed to increment SMS count for phone %s: %v", in.Phone, err)
		return nil, errors.New("系统错误")
	}
	log.Printf("SendCode: SMS count incremented for phone %s, new count: %d", in.Phone, newCount)

	return &user.SendCodeResponse{
		Message: "验证码发送成功",
	}, nil
}
