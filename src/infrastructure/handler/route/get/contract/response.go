package contract

type GetRouteResponse struct {
	Id           int64
	RelativePath string
	UrlTarget    string
	TypeTarget   string
	Enable       bool
}
