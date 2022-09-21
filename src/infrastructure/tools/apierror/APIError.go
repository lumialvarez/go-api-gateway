package apierror

type APIError struct {
	Status  int64
	Message string
	Err     string
	Cause   []string
}
