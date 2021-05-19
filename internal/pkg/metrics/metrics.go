package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

const (
	namespace = "xinliangnote"
	subsystem = "go_gin_api"
)

// metricsRequestsTotal metrics for request total counter（Counter）
var metricsRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "request(s) total",
	},
	[]string{"method", "path"},
)

// metricsRequestsCost metrics for requests cost cumulative histogram (Histogram)
var metricsRequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost",
		Help:      "request(s) cost seconds",
	},
	[]string{"method", "path", "success", "http_code", "business_code", "cost_seconds", "trace_id"},
)

func init() {
	prometheus.MustRegister(metricsRequestsTotal, metricsRequestsCost)
}

// RecordMetrics record index
func RecordMetrics(method, uri string, success bool, httpCode, businessCode int, costSeconds float64, traceId string) {
	metricsRequestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   uri,
	}).Inc()

	metricsRequestsCost.With(prometheus.Labels{
		"method":        method,
		"path":          uri,
		"success":       cast.ToString(success),
		"http_code":     cast.ToString(httpCode),
		"business_code": cast.ToString(businessCode),
		"cost_seconds":  cast.ToString(costSeconds),
		"trace_id":      traceId,
	}).Observe(costSeconds)
}
