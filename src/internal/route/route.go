package route

type Route struct {
	id           int64
	relativePath string
	urlTarget    string
	typeTarget   string
	secure       bool
	enable       bool
}

func NewRoute(id int64, relativePath string, urlTarget string, typeTarget string, secure bool, enable bool) *Route {
	return &Route{id: id, relativePath: relativePath, urlTarget: urlTarget, typeTarget: typeTarget, secure: secure, enable: enable}
}

func (r *Route) Id() int64 {
	return r.id
}

func (r *Route) RelativePath() string {
	return r.relativePath
}

func (r *Route) UrlTarget() string {
	return r.urlTarget
}

func (r *Route) TypeTarget() string {
	return r.typeTarget
}

func (r *Route) Secure() bool {
	return r.secure
}

func (r *Route) Enable() bool {
	return r.enable
}

func (r *Route) UpdateRoute(route Route) {
	r.urlTarget = route.urlTarget
	r.secure = route.secure
	r.enable = route.enable
}
