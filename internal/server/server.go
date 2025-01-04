package server

import (
	"context"
	"net/http"

	"aggreagtor/internal/config"
	"aggreagtor/internal/handler"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func New(cfg config.Server, handler *handler.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:           cfg.Host + ":" + cfg.Port,
		Handler:        handler.InitRoute(),
		ReadTimeout:    cfg.ReadTimeOut,
		WriteTimeout:   cfg.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}}
}
