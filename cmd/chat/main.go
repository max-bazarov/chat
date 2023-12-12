package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/max-bazarov/chat/internal/app"
	"github.com/max-bazarov/chat/internal/database"
	"github.com/max-bazarov/chat/internal/database/postgres"
	"github.com/max-bazarov/chat/internal/database/redis"
	"github.com/max-bazarov/chat/internal/service"

	"github.com/max-bazarov/chat/internal/chat"
	"github.com/max-bazarov/chat/internal/config"

	"github.com/max-bazarov/chat/internal/transport/rest"
	"github.com/max-bazarov/chat/internal/transport/ws"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Error(fmt.Sprintf("failed to initialize DB: %s", err.Error()))
	}

	rdb, err := redis.NewRedisDB()
	if err != nil {
		log.Error(fmt.Sprintf("failed to initialize DB: %s", err.Error()))
	}

	arepo := database.NewRepository(db, rdb)
	asvc := service.NewService(arepo)
	ah := rest.NewHandler(asvc)

	hub := chat.NewHub()
	wsh := ws.NewHandler(hub)
	go hub.Run()

	if err := app.Run(cfg.Port, rest.InitRoutes(ah, wsh)); err != nil {
		log.Error(fmt.Sprintf("error occured while running http server: %s", err.Error()))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
