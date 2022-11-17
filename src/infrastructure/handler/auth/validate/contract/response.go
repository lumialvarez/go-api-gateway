package contract

type ValidateAuthResponse struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}
