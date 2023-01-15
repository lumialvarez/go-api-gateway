package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"strings"
)

const unknown = "unknown"

type Prometheus struct {
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total_path",
		Help: "Number of get requests.",
	},
	[]string{"path", "method", "target"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status_path",
		Help: "Status of HTTP response",
	},
	[]string{"status", "path", "method", "target"},
)

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds_path",
	Help: "Duration of HTTP requests.",
}, []string{"path", "method", "target"})

func (prometheusProvider *Prometheus) PrometheusMiddleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	application := unknown

	uriSegments := strings.Split(path, "/")
	for i := range uriSegments {
		if len(uriSegments[i]) > 0 {
			if len(uriSegments) > (i+1) && uriSegments[i+1] == "api" {
				application = uriSegments[i]
			}
			break
		}
	}

	totalRequests.WithLabelValues(path, method, application).Inc()

	timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, method, application))
	defer timer.ObserveDuration()

	ctx.Next()

	statusCode := ctx.Writer.Status()

	responseStatus.WithLabelValues(strconv.Itoa(statusCode), path, method, application).Inc()

	fmt.Print(statusCode)
}

func NewPrometheusProvider() Prometheus {
	err := prometheus.Register(totalRequests)
	if err != nil {
		log.Print("Error to register prometheus TotalRequest metric Collector", err)
	}
	err = prometheus.Register(responseStatus)
	if err != nil {
		log.Print("Error to register prometheus ResponseStatus metric Collector", err)
	}
	err = prometheus.Register(httpDuration)
	if err != nil {
		log.Print("Error to register prometheus HttpDuration metric Collector", err)
	}
	return Prometheus{}
}
