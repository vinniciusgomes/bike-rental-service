package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

// PrometheusMiddleware registers Prometheus metrics for each HTTP request.
//
// It is a Gin middleware function that calculates the duration of the request and
// registers metrics for the number of requests and the duration of the requests.
// The metrics are registered using the Prometheus client library.
//
// Returns a Gin handler function.
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process the request
		c.Next()

		// Calculate the request duration
		duration := time.Since(start).Seconds()

		// Get request details
		path := c.FullPath()
		method := c.Request.Method
		status := http.StatusText(c.Writer.Status())

		// Record metrics
		httpRequestsTotal.WithLabelValues(path, method, status).Inc()
		httpRequestDuration.WithLabelValues(path, method).Observe(duration)
	}
}

// PrometheusHandler returns a Gin handler function that wraps the promhttp.Handler.
//
// The promhttp.Handler is a middleware that exposes Prometheus metrics.
// It registers the metrics and serves them on a specific HTTP endpoint.
//
// Returns a Gin handler function.
func PrometheusHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
