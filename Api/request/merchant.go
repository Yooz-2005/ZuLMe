package request

type MerchantRegisterRequest struct {
	Name        string `json:"name" form:"name"`
	Phone       string `json:"phone" form:"phone"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	ConfirmPass string `json:"confirm_pass" form:"confirm_pass"`
}

type MerchantLoginRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
