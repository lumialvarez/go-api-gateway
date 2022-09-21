package apierror

type BadRequest struct {
	message string
	Cause   []string
}

func NewBadRequest(message string, cause ...string) BadRequest {
	return BadRequest{message: message, Cause: cause}
}

func (err BadRequest) Error() string {
	return err.message
}
