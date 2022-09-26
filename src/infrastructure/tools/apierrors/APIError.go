package apierrors

type APIError struct {
	Status  int
	Message string
	Err     string
	Cause   []string
}

func (err APIError) Error() string {
	return err.Message
}
