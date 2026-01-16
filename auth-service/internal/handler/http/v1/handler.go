package v1

import (
	"auth-service/config"
	"auth-service/internal/domain/user"
	"context"
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const apiPath = "/v1"

type AuthService interface {
	Register(ctx context.Context, email, password string) (user.UserID, error)
	Login(ctx context.Context, email, password string) (string, error)
}

func SetupHandlers(
	handler *gin.Engine,
	log *slog.Logger,
	cfg *config.Config,
	authService AuthService,
) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = strings.Split(cfg.HTTP.CORSAllowOrigins, " ")
	corsCfg.AllowCredentials = true

	handler.Use(cors.New(corsCfg))
}

