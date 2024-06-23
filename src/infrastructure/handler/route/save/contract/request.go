package contract

type SaveRouteRequest struct {
	Id           int64    `json:"id"`
	RelativePath string   `json:"relative_path" binding:"required"`
	UrlTarget    string   `json:"url_target" binding:"required"`
	TypeTarget   string   `json:"type_target" binding:"required"`
	Secure       bool     `json:"secure"`
	Enable       bool     `json:"enable"`
	Methods      []string `json:"methods"`
}
