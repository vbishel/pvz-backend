package app

import (
	"auth-service/internal/config"
	v1 "auth-service/internal/handler/http/v1"
	"auth-service/internal/lib/logger/sl"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/database/postgres"
	"auth-service/pkg/httpserver"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func MustRun(cfg *config.Config) {
	log := configureLogger(cfg)
	dsn := buildPostgresDSN(cfg)

	pg, err := postgres.New(dsn, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to postgres: %s", err.Error()))
	}

	userRepository := repository.NewUserRepository(pg)

	usersService := service.NewUsersService(log, cfg, userRepository)
	authService := service.NewAuthService(log, cfg, usersService, usersService)

	handler := gin.New()
	v1.SetupHandlers(handler, log, cfg, authService, usersService)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-interrupt:
		log.Info("app.MustRun signal", "signal", s.String())
	case err = <-httpServer.Notify():
		log.Error("app.MustRun httpServer.Notify", sl.Err(err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app.MustRun httpServer.Shutdown", sl.Err(err))
	}
}

func configureLogger(cfg *config.Config) *slog.Logger {
	var level slog.Leveler

	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	return log
}
