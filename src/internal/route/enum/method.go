package enum

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

func (m Method) Value() string {
	return string(m)
}

func (m Method) IsValid() bool {
	switch m {
	case GET:
		return true
	case POST:
		return true
	case PUT:
		return true
	case DELETE:
		return true
	default:
		return false
	}
}
