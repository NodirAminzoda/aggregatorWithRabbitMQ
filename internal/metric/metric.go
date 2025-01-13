package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"}, // Метки: метод запроса и конечная точка
	)

	ResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status",
			Help: "HTTP response statuses",
		},
		[]string{"status"}, // Метка: статус ответа
	)

	RequestExecTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_execution_time_seconds",
			Help: "Histogram of HTTP request execution times",
			Objectives: map[float64]float64{
				0.99: 0.001,
				0.95: 0.001,
				0.5:  0.05,
			},
		},
		[]string{"method", "endpoint"},
	)
)
