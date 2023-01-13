package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type PrometheusProvider struct {
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total_path",
		Help: "Number of get requests.",
	},
	[]string{"path", "method"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status_path",
		Help: "Status of HTTP response",
	},
	[]string{"status", "path", "method"},
)

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds_path",
	Help: "Duration of HTTP requests.",
}, []string{"path", "method"})

func (provider *PrometheusProvider) PrometheusMiddleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	totalRequests.WithLabelValues(path, method).Inc()

	timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, method))
	defer timer.ObserveDuration()

	ctx.Next()

	statusCode := ctx.Writer.Status()

	responseStatus.WithLabelValues(strconv.Itoa(statusCode), path, method).Inc()

	fmt.Print(statusCode)
}

func NewPrometheusProvider() PrometheusProvider {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
	return PrometheusProvider{}
}
