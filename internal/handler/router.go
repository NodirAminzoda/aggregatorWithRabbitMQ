package handler

import (
	"aggreagtor/internal/config"
	"aggreagtor/internal/metric"
	"aggreagtor/internal/middleware"
	"aggreagtor/internal/modules/service"
	"aggreagtor/pkg/zerolog"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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
	//initialization metrics
	metric.InitMetrics()

	router := mux.NewRouter()

	router.Handle("/metrics", promhttp.Handler())
	router.Use(middlware.SetXRequestID, metric.MetricsMiddleware) // add middleware

	router.HandleFunc("/precheck", h.PreCheck).Methods(http.MethodPost)

	// add route for metrics
	return router
}
