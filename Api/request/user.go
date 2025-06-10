package request

type UserRegisterRequest struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

type SendCodeRequest struct {
	Phone  string `json:"phone" form:"phone"`
	Source string `json:"source" form:"source"`
}
