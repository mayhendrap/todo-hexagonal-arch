package interfaces

import (
	"todo-hexagonal-arch/internal/core/domain"
)

type AuthService interface {
	Register(user domain.User) (string, error)
	Login(email, password string) (string, error)
}

type TodoService interface {
	FindOne(todoID, userID string) (domain.Todo, error)
	FindAll(todoID string) ([]domain.Todo, error)
	Create(todo domain.Todo) error
	Update(todo domain.Todo) (domain.Todo, error)
	Delete(todoID, userID string) error
}
