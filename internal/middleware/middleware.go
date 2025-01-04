package middlware

import (
	"context"
	"github.com/google/uuid"
	"net/http"

	"aggreagtor/internal/costants"
)

func SetXRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logID := uuid.New().String()
		ctx := context.WithValue(context.Background(), constants.XRequestID, logID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
