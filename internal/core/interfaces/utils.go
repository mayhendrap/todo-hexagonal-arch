package interfaces

import "todo-hexagonal-arch/internal/core/domain"

type TokenUtil interface {
	GenerateToken(userID string, email string) (string, error)
	ValidateToken(tokenString string) (*domain.CustomClaims, error)
}
