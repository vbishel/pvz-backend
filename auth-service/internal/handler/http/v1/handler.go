package v1

import (
	"auth-service/config"
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const apiPath = "/v1"

func SetupHandlers(
	handler *gin.Engine,
	log *slog.Logger,
	cfg *config.Config,
	authService AuthService,
	usersService UsersService,
) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = strings.Split(cfg.HTTP.CORSAllowOrigins, " ")
	corsCfg.AllowCredentials = true

	handler.Use(cors.New(corsCfg))
}

