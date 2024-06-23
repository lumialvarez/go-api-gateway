package contract

type GetRouteResponse struct {
	Id           int64    `json:"id"`
	RelativePath string   `json:"relative_path"`
	UrlTarget    string   `json:"url_target"`
	TypeTarget   string   `json:"type_target"`
	Secure       bool     `json:"secure"`
	Enable       bool     `json:"enable"`
	Methods      []string `json:"methods"`
}
