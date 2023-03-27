package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type jwtUtil struct {
}

const (
	TokenSecret = "nt1HTxRgdL4HnFm0RfNBSx9VkPQzXVQWkd_dqPGSfNZUQl9a2Gxc4xb6SLlMie6vYCM"
)

func NewJwtUtil() interfaces.TokenUtil {
	return &jwtUtil{}
}

func (tu *jwtUtil) GenerateToken(userID string, email string) (string, error) {

	duration := 30 * time.Minute
	exp := jwt.NewNumericDate(time.Now().Add(duration))
	now := jwt.NewNumericDate(time.Now())

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: exp,
			IssuedAt:  now,
		},
	})

	token, err := claims.SignedString([]byte(TokenSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (tu *jwtUtil) ValidateToken(tokenString string) (*domain.CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSecret), nil
	})
	if err != nil {
		return &domain.CustomClaims{}, err
	}

	if claims, ok := token.Claims.(*domain.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return &domain.CustomClaims{}, errors.New("can't cast token to custom claims")
}
