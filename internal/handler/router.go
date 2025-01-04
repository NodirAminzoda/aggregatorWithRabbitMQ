package handler

import (
	"github.com/gorilla/mux"

	"aggreagtor/internal/config"
	"aggreagtor/internal/middleware"
	"aggreagtor/internal/modules/service"
	"aggreagtor/pkg/zerolog"
)

type Handler struct {
	cfg     *config.Server
	log     *zerolog.Logger
	service service.IService
}

func New(cfg *config.Server, log *zerolog.Logger, service service.IService) *Handler {
	return &Handler{
		cfg:     cfg,
		log:     log,
		service: service,
	}
}

func (h *Handler) InitRoute() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlware.SetXRequestID)
	router.HandleFunc("/precheck", h.PreCheck)
	return router
}
