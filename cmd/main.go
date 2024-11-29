package main

import (
	"TODO_APP/internal/config"
	handlers "TODO_APP/internal/handlers"
	"TODO_APP/internal/repository"
	"TODO_APP/internal/service"
	"TODO_APP/internal/storage"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "develop"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setUpLogger(cfg.Env)

	db, err := storage.New(*cfg)
	if err != nil {
		log.Error("failed to connect storage", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handler := handlers.NewHandle(serv)

	_ = handler

	// TODO: router gin

	// TODO: rus server
}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
