package request

type UserRegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
