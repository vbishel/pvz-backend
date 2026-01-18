package v1

import (
	"auth-service/internal/config"
	"auth-service/internal/lib/jwt"
	"auth-service/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func tokenMiddleware(l *slog.Logger, cfg *config.Config) gin.HandlerFunc {
	const op = "handler.http.v1.tokenMiddleware"
	log := l.With(
		slog.String("op", op),
	)
	return func(c *gin.Context) {
		token, err := c.Cookie(cfg.AccessToken.CookieKey)
		if err != nil {
			log.Warn("no access token provided", sl.Err(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no access token provided"})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token, cfg.AccessToken.SigningKey)
		if err != nil {
			log.Warn("invalid access token", sl.Err(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			c.Abort()
			return
		}

		log.Info("middleware received correct token", "userID", claims.UserID)
		c.Set(UserIDContextKey, claims.UserID)
		c.Set(UserEmailContextKey, claims.Email)

		c.Next()
	}
}
