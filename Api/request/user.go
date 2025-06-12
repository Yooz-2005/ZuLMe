package request

type UserRegisterRequest struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

type SendCodeRequest struct {
	Phone  string `json:"phone" form:"phone"`
	Source string `json:"source" form:"source"`
}

type UpdateUserProfileRequest struct {
	RealName       string `json:"real_name" form:"real_name"`
	IdType         string `json:"id_type" form:"id_type"`
	IdNumber       string `json:"id_number" form:"id_number"`
	IdExpireDate   string `json:"id_expire_date" form:"id_expire_date"`
	Email          string `json:"email" form:"email"`
	Province       string `json:"province" form:"province"`
	City           string `json:"city" form:"city"`
	District       string `json:"district" form:"district"`
	EmergencyName  string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone string `json:"emergency_phone" form:"emergency_phone"`
}

type UpdateUserPhoneRequest struct {
	Phone string `json:"phone" form:"phone"`
}
