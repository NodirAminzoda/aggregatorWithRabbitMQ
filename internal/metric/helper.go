package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

// Обертка для ResponseWriter, чтобы захватить статус ответа
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// InitMetrics регистрирует метрики
func InitMetrics() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(ResponseStatus)
}
