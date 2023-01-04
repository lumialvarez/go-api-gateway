package contract

type UpdateAuthRequest struct {
	User User `json:"user" binding:"required"`
}

type User struct {
	UserId   int64  `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
	Role     string `json:"role" binding:"required"`
	Status   bool   `json:"status" binding:"required"`
}
