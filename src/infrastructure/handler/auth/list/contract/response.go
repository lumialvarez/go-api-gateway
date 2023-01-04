package contract

type ListAuthResponse struct {
	Users []User `json:"users"`
}

type User struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}
