package v1

import (
	"auth-service/internal/config"
	"auth-service/internal/domain/user"
	"auth-service/internal/lib/logger/sl"
	"auth-service/pkg/apperrors"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	log          *slog.Logger
	cfg          *config.Config
	authService  AuthService
	usersService UsersService
}

func newAuthHandler(handler *gin.RouterGroup, log *slog.Logger, cfg *config.Config, auth AuthService, users UsersService) {
	h := &authHandler{
		log:          log,
		cfg:          cfg,
		authService:  auth,
		usersService: users,
	}

	g := handler.Group("/auth")
	g.POST("/login", h.login)
	g.POST("/register", h.register)

	protected := g.Group("/").Use(tokenMiddleware(log, cfg))
	protected.POST("/logout", h.logout)
	protected.GET("/me", h.me)
	protected.GET("/validate-token", h.validateToken)
}

type registerDTO struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func (h *authHandler) register(c *gin.Context) {
	const op = "handler.http.v1.auth.register"

	log := h.log.With(
		slog.String("op", op),
	)

	log.Info("trying to register user")
	var data registerDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Warn("validation failed", sl.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := h.authService.Register(c.Request.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserAlreadyExists) {
			log.Warn("user already exists", "email", data.Email)
			c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrUserAlreadyExists.Error()})
			return
		}

		log.Error("registration failed", "email", data.Email, sl.Err(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	log.Info("user registered successfully", "userID", uid)
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

type loginDTO struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func (h *authHandler) login(c *gin.Context) {
	const op = "handler.http.v1.auth.login"

	log := h.log.With(
		slog.String("op", op),
	)

	log.Info("trying to login user")
	var data loginDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Warn("validation failed", sl.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserIncorrectEmailOrPassword) {
			log.Warn("login failed: incorrect email or password", "email", data.Email)
			c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrUserIncorrectEmailOrPassword.Error()})
			return
		}

		if errors.Is(err, apperrors.ErrUserNotFound) {
			log.Warn("login failed: user does not exist", "email", data.Email)
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			return
		}

		log.Error("login failed: internal server error", sl.Err(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	isSecure := h.cfg.App.Env == "production"

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		h.cfg.AccessToken.CookieKey,
		token,
		int(h.cfg.AccessToken.TTL),
		"/",
		"",
		isSecure,
		true,
	)
	log.Info("login successful", "email", data.Email)
	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully"})
}

func (h *authHandler) logout(c *gin.Context) {
	const op = "handler.http.v1.auth.logout"
	log := h.log.With(
		slog.String("op", op),
	)

	_, err := c.Cookie(h.cfg.AccessToken.CookieKey)
	id := c.GetInt(UserIDContextKey)
	if err != nil {
		log.Error("logout failed: no cookie provided", "userID", id)
		c.JSON(http.StatusUnauthorized, "no cookie provided")
		return
	}

	isSecure := h.cfg.App.Env == "production"
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		h.cfg.AccessToken.CookieKey,
		"",
		-1,
		"/",
		"",
		isSecure,
		true,
	)
	log.Info("logout successful", "userID", id)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *authHandler) me(c *gin.Context) {
	const op = "handler.http.v1.auth.me"

	log := h.log.With(
		slog.String("op", op),
	)

	log.Info("getting info about user")

	uid, ok := c.Get(UserIDContextKey)
	if !ok {
		log.Warn("failed to get info about user")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	uidint, ok := uid.(int)
	if !ok || uid == 0 {
		log.Error("invalid user id", "userID", uid)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	u, err := h.usersService.Find(c.Request.Context(), user.UserID(uidint))
	if err != nil {
		log.Error("failed to get info about user", "userID", uidint, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("info returned successfully", "userID", uidint)

	c.JSON(http.StatusOK, gin.H{
		"id":      u.ID,
		"email":   u.Email,
		"role_id": u.RoleID,
	})
}

func (h *authHandler) validateToken(c *gin.Context) {
	const op = "handlers.HandleValidateToken"

	log := h.log.With(
		slog.String("op", op),
	)

	email, ok := c.Get(UserEmailContextKey)
	if !ok {
		log.Warn("failed to parse user email from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id, ok := c.Get(UserIDContextKey)
	if !ok {
		log.Warn("failed to parse user id from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idint, ok := id.(int)
	if !ok {
		log.Warn("failed to convert uid to int", "userID", id)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	emailstr, ok := email.(string)
	if !ok {
		log.Warn("failed to convert email to string", "email", email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Header("X-User-ID", strconv.Itoa(idint))
	c.Header("X-User-Email", emailstr)

	log.Info("token successfuly validated", "userID", idint)
	c.JSON(http.StatusOK, gin.H{
		"email": email,
		"id":    id,
	})
}
