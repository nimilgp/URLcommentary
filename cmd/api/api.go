package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/nimilgp/URLcommentary/internal/config"
	"github.com/nimilgp/URLcommentary/internal/dblayer"
)

type APIServer struct {
	baseURL string
	version string
	querier *dblayer.Queries
	ctx     context.Context
	logger  *slog.Logger
}

func GetAPIServer(querier *dblayer.Queries, ctx context.Context) *APIServer {
	return &APIServer{
		baseURL: config.Cfg.BaseURL,
		version: config.Cfg.Version,
		querier: querier,
		ctx:     ctx,
	}
}

func (s *APIServer) Run(logger *slog.Logger) error {

	server := http.Server{
		Addr:         s.baseURL,
		Handler:      s.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	return server.ListenAndServe()
}
