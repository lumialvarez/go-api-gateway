package contract

type LoginAuthResponse struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}
