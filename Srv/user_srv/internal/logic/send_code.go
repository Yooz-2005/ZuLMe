package logic

import (
	"Common/appconfig"
	"errors"
	"golang.org/x/exp/rand"
	"models/model_redis"
	"strconv"
	user "user_srv/proto_user"
)

// todo 生成五位数随机邀请码
func generateCode() string {
	return strconv.Itoa(rand.Intn(1000) + 9000)
}

func SendCode(in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	code := generateCode()
	//_, err := pkg.SendSms(code, in.Phone)
	//if err != nil {
	//	log.Println("短信发送失败")
	//	return nil, errors.New("短信发送失败")
	//}
	//if *sms.Body.Code != "OK" {
	//	log.Println("短信发送失败")
	//	return nil, errors.New(*sms.Body.Message)
	//}
	//存储验证码
	err := model_redis.SaveVerificationCode(in.Source, in.Phone, code)
	if err != nil {
		return nil, errors.New("验证码存储失败")
	}

	// 增加发送次数
	count, err := model_redis.IncrementSMSCount(in.Phone, in.Source)
	if err != nil {
		return nil, errors.New("系统错误")
	}
	//判断发送次数是否超过限制
	if count > appconfig.ConfData.MaxSend.Count {
		return nil, errors.New("发送次数过多,请稍后再试")
	}
	return &user.SendCodeResponse{
		Message: "验证码发送成功",
	}, nil
}
