package metric

import (
	"net/http"
	"time"
)

// MetricsMiddleware - Middleware для сбора метрик
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Создаем обертку для Writer, чтобы отслеживать статус ответа
		ww := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		// Передаем управление следующему обработчику
		next.ServeHTTP(ww, r)

		// Увеличиваем счетчики
		RequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
		ResponseStatus.WithLabelValues(http.StatusText(ww.statusCode)).Inc()

		duration := time.Since(start).Seconds()
		RequestExecTime.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}
