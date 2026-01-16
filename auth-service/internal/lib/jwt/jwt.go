package jwt

import (
	"auth-service/internal/domain/entity"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoSigningKey         = errors.New("empty signing key")
	ErrNoClaims             = errors.New("error getting claims from token")
	ErrUnexpectedSignMethod = errors.New("unexpected signing method")
)

type Claims struct {
	UserID int    `json:"uid"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
	jwt.RegisteredClaims
}

func NewToken(
	user *entity.User,
	duration time.Duration,
	signingKey string,
) (string, error) {
	if signingKey == "" {
		return "", ErrNoSigningKey
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}