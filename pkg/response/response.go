package response

import (
	"encoding/json"
	"net/http"

	"aggreagtor/pkg/zerolog"
)

func WriteResponse(w http.ResponseWriter, logg *zerolog.Logger, statusCode int, message string, requestID string, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	if message == "" {
		message = Errors[statusCode]
	}

	w.WriteHeader(statusCode)

	resp := map[string]interface{}{
		"status_code": statusCode,
		"message":     message,
	}

	// Если тело передано, добавляем его в ответ
	if body != nil {
		resp["body"] = body
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		logg.ErrorWithPrefix(requestID, "WriteResponse: Error marshaling response, Error", err)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		logg.ErrorWithPrefix(requestID, "WriteResponse: Error writing response body, Error", err)
	}
}
