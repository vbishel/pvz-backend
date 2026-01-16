package v1

import (
	"auth-service/config"
	"auth-service/pkg/apperrors"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	log         *slog.Logger
	cfg         *config.Config
	authService AuthService
}

func newAuthHandler(handler *gin.RouterGroup, log *slog.Logger, cfg *config.Config, auth AuthService) {
	h := &authHandler{
		log:         log,
		cfg:         cfg,
		authService: auth,
	}

	g := handler.Group("/auth")
	g.POST("/login", h.login)
	g.POST("/register", h.register)
	g.POST("/logout", h.logout)
	g.GET("/me", h.me)
}

type registerDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

func (h *authHandler) register(c *gin.Context) {
	var data registerDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authService.Register(context.Background(), data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

type loginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

func (h *authHandler) login(c *gin.Context) {
	var data loginDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(context.Background(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserIncorrectEmailOrPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect email or password"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	isSecure := h.cfg.App.Env == "production"

	c.SetCookie(
		h.cfg.AccessToken.CookieKey,
		token,
		int(h.cfg.AccessToken.TTL),
		"/",
		"",
		isSecure,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully"})
}

func (h *authHandler) logout(c *gin.Context) {
	const op = "handler.http.v1.logout"
	log := h.log.With(
		slog.String("op", op),
	)

	isSecure := h.cfg.App.Env == "production"
	_, err := c.Cookie(h.cfg.AccessToken.CookieKey)
	email := c.GetString(h.cfg.GinCtx.EmailKey)
	id := c.GetInt(h.cfg.GinCtx.UserIDKey)
	if err != nil {
		log.Error(fmt.Sprintf("logout failed: no cookie provided, id: %v, email: %s", id, email))
		c.JSON(http.StatusUnauthorized, "no cookie provided")
		return
	}

	c.SetCookie(
		h.cfg.AccessToken.CookieKey,
		"",
		-1,
		"/",
		"",
		isSecure,
		true,
	)
	log.Info(fmt.Sprintf("logout successful for user id: %s", id))
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *authHandler) me(c *gin.Context) {
	
}