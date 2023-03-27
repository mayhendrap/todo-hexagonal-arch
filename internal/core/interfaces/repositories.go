package interfaces

import (
	"todo-hexagonal-arch/internal/core/domain"
)

type TodoRepository interface {
	FindOne(todoID string, userID string) (domain.Todo, error)
	FindAll(userID string) ([]domain.Todo, error)
	Create(todo domain.Todo) error
	Update(todo domain.Todo) (domain.Todo, error)
	Delete(todoID, userID string) error
}

type UserRepository interface {
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (string, error)
	Update(user domain.User) (domain.User, error)
	Delete(userID string) error
}
