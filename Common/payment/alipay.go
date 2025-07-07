package payment

import (
	"fmt"

	"github.com/smartwalle/alipay/v3"
)

type Pay interface {
	Pay(subject, outTradeNo, totalAmount string) string
}

type AliPay struct {
	AppId      string
	PrivateKey string
	NotifyUrl  string
	ReturnUrl  string
}

func NewAliPay() *AliPay {
	return &AliPay{
		AppId:      "9021000142691060",
		PrivateKey: "MIIEowIBAAKCAQEAkMDIOh01PDKprXYXIpMRmtbNokRXZ811bgy7gwOtBq0C0OUtNG/VKWYnVeX62lUgDlOsSxerIojqd1VMwGJYdqJebRxv48rdLnQNhghVi96fVwoPNmpuhWQrpPtldaty8/eBh5lu/aEhsdzwrfsA1pSy+fqC+m2qofOdRUewpN7ioLp9BkzWKhBYjv2VGbUGicdKaomBV/itIPj2l4dtN79YiXqGmUEY9S2L8eXWuijNFzn0f8Pjhk1rBKdSKx33zamvEvsdYjpK0FrnqgO2bcBPQSJCpo1LNoKvHEfuwmlz17D50FEG9o5H+Kg9cwrsvsiSqeV+3qd5fVSGJe0SpQIDAQABAoIBAApXW3abY5Q7YlfcGecEbruLIQG4tC3cRclo63R+rLvPXGYkGUinvKjKfYDrfzJd5n5fh+2NcVJ0SZvg4izUWdGP5AMThwLVGVwaOmUa71Ggw+1p8JBIpkCtWVjeBO+VIg8/3WcclQtVmZylViR3zCWIVgj51qWhdBFB4Dh02K/OiNT9RLuB4hZUIT3FhlqcOFLdJ2VWxDZxr0xMNctxva2doIPway1EMIEIm52HNKszAVD4r6asdT2N9spykuhKAhsw2hVXTdkZJohUWZ+8fwONpsQXOaGNOVkiN9tQMhVnn6YodZId9Il8kxhgnDoUStYUy+LerIdjBQPp4/0xX4ECgYEA2liFpyf2krMZ/PboZsh0GI6lwYQ7EEIezEMfAq8gLE30yhViDDP0I3gVWla2oE+IP/BRnRh4kImLOoF3mHiDtTkfLHpMOIBmqVAwRlF3E/QhMFSVDqK2Aefx7wn9ZS4pC8vA+7FKQq6mkQ/5mEOVxuO1siG4X2YELcqFhJ1ZpPUCgYEAqbdOOJLZLwD02a+NWu1FlLAkXSF1ZYlx5VnjBjBiYW1GHtNm7xg9cW4pU7euYbLtulEeVChovs+Hcm5VmPQOX5pk0oY8VZro+atQQDvkS9z9JgvT3JkQUgkmEzT0LhIimrCewbnKJfcj7/CuZLzTQKCWws49/xlOFhBB8YIoqPECgYA5Wq1o9i9n45H9B+KONTOBy96wkYpuP+AVKcB4lQXvfV7CwpEpwW/s7Ts2qrZ4L8wLd5YInQf2d5rR+HYw399A+Es/BLUG1nuhGAZGQln0LNmW93DcElOa9pFviAE+1bxEc/YyZySplXT9f+PBYmdyghgVPZRPwt1wJdWiuy50DQKBgEuf7dARx4tFXtW9fzx8LBw0XQ/Ov/Qtyb3MTvhMCRqmya8kvmJeJ8rqrqmqWJ8aTwpN6TjRBNO5v/5CogvU/K6nKrQQssPmklfmeY0V4wXXBEq2zIIpBne3seqvFvuMgzTw7N0gP66pMK9TTTKAXZYXuPY3VrLcgMoeFnRabmnRAoGBALNQBntY5tZgWhiFWXhB/oOGmL/WOXw8aoESQTKsTDBipGqXmAZ6qQlfHQFsZhQBEcxZW9YBUT6yLsNASfhFYr8wxa9tcrVqN17AN1o88xdmdV2LXEObzWyqJJMJNHICciqya0qLbJvZTW8j6ItyhESB09rS9pHKuECn1YznJTPx",
		NotifyUrl:  "http://7baa3296.r27.cpolar.top/payment/alipay/notify",
		ReturnUrl:  "http://localhost:3000/payment/success",
	}
}

func (a *AliPay) Pay(subject, outTradeNo, totalAmount string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyUrl
	p.ReturnURL = a.ReturnUrl
	p.Subject = subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
