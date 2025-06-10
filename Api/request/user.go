package request

type UserRegisterRequest struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
	Code  string `json:"code" form:"code" binding:"required"`
}

type SendCodeRequest struct {
	Phone  string `json:"phone" form:"phone" binding:"required"`
	Source string `json:"source" form:"source" binding:"required"`
}
