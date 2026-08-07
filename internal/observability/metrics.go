package observability

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "redira_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "redira_http_request_duration_seconds",
			Help: "Total HTTP request duration in seconds",
		},
		[]string{"method", "path"},
	)
)

func MetricsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		requestCount.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			strconv.Itoa(c.Writer.Status()),
		).Inc()

		requestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}

func RegisterMetrics(router *gin.Engine) {

	router.GET(
		"/metrics",
		gin.WrapH(
			promhttp.Handler(),
		),
	)

}
