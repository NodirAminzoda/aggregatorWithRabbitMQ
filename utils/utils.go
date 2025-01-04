package utils

import (
	"context"
	"github.com/google/uuid"

	"aggreagtor/internal/costants"
)

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(constants.XRequestID).(string); ok {
		return requestID
	}
	return ""
}

func GenerateUUID() string {
	return uuid.New().String()
}
