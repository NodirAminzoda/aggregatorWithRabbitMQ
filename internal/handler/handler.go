package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"aggreagtor/internal/costants"
	"aggreagtor/internal/domain/precheck"
	"aggreagtor/pkg/response"
	"aggreagtor/utils"
)

func (h *Handler) PreCheck(w http.ResponseWriter, r *http.Request) {

	requestID := utils.GetRequestID(r.Context())
	ctx := context.WithValue(context.Background(), constants.XRequestID, requestID)

	var request precheck.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.ErrorWithPrefix("", "[PreCheck] badRequest body", err)
		response.WriteResponse(w, h.log, http.StatusBadRequest, "Неверный формат входных данных", "", nil)
		return
	}

	h.log.InfoWithPrefix("", "[PreCheck] Request Body", fmt.Sprintf("%v", request))

	resp, err := h.service.PreCheckService(ctx, request)
	if err != nil {
		h.log.ErrorWithPrefix(requestID, "[PreCheck] Error", err)
		response.WriteResponse(w, h.log, http.StatusBadRequest, "Неверный формат входных данных", requestID, nil)
		return
	}

	response.WriteResponse(w, h.log, http.StatusOK, "Ok", requestID, resp)
}
