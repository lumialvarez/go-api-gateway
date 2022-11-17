package contract

type ValidateAuthRequest struct {
	Token string `json:"token" binding:"required"`
}
