package main

import (
	"TODO_APP/internal/config"
	handlers "TODO_APP/internal/handlers"
	"TODO_APP/internal/repository"
	"TODO_APP/internal/service"
	"TODO_APP/internal/storage"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		log.Error("failed to connect storage", "error", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handlers.NewHandle(serv)

	router := handlers.InitRoutes()

	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Idle_timeout,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			log.Error("failde to start the server")
		}
	}()

	log.Error("server running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Error("server stoped")

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
