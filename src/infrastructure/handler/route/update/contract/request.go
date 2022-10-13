package contract

type UpdateRouteRequest struct {
	Id           int64  `json:"id" binding:"required"`
	RelativePath string `json:"relative_path"`
	UrlTarget    string `json:"url_target" binding:"required"`
	TypeTarget   string `json:"type_target"`
	Secure       bool   `json:"secure"`
	Enable       bool   `json:"enable"`
}
