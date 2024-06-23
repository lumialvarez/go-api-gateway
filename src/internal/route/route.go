package route

import (
	"github.com/lumialvarez/go-api-gateway/src/internal/route/enum"
	"strings"
)

type Route struct {
	id           int64
	relativePath string
	urlTarget    string
	typeTarget   string
	secure       bool
	enable       bool
	methods      []enum.Method
}

func NewRoute(id int64, relativePath string, urlTarget string, typeTarget string, secure bool, enable bool, methods []enum.Method) *Route {
	return &Route{id: id, relativePath: relativePath, urlTarget: urlTarget, typeTarget: typeTarget, secure: secure, enable: enable, methods: methods}
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

func (r *Route) Methods() []enum.Method {
	return r.methods
}

func (r *Route) UpdateRoute(route Route) {
	r.urlTarget = route.urlTarget
	r.secure = route.secure
	r.enable = route.enable
	r.methods = route.methods
}

func (r *Route) GetStringMethods() []string {
	methods := make([]string, 0)
	for _, method := range r.methods {
		methods = append(methods, method.Value())
	}
	return methods
}

func (r *Route) MethodExists(inputMethod string) bool {
	inputMethod = strings.ToUpper(inputMethod)
	for _, method := range r.methods {
		if method.Value() == inputMethod {
			return true
		}
	}
	return false
}
