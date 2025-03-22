package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/nimilgp/URLcommentary/cmd/api"
	"github.com/nimilgp/URLcommentary/internal/config"
	database "github.com/nimilgp/URLcommentary/internal/db"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))
	ctx := context.Background()
	db := database.New(config.Cfg.Dsn, ctx)
	defer db.Close(ctx)

	logger.Info("starting server", "base url", config.Cfg.BaseURL)
	if err := api.GetAPIServer(db.Queries, ctx).Run(logger); err != nil {
		log.Fatal(err)
	}
}
