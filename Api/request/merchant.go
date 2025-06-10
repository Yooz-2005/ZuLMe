package request

type MerchantRegisterRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Phone       string `json:"phone" form:"phone" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Password    string `json:"password" form:"password" binding:"required,min=6"`
	ConfirmPass string `json:"confirm_pass" form:"confirm_pass" binding:"required,eqfield=Password"`
}

type MerchantLoginRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
