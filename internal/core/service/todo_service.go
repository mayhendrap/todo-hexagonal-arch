package service

import (
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type todoService struct {
	todoRepository interfaces.TodoRepository
}

func NewTodoService(todoRepository interfaces.TodoRepository) interfaces.TodoService {
	return &todoService{todoRepository: todoRepository}
}

func (s *todoService) FindOne(todoID, userID string) (domain.Todo, error) {
	result, err := s.todoRepository.FindOne(todoID, userID)
	return result, err
}

func (s *todoService) FindAll(userID string) ([]domain.Todo, error) {
	return s.todoRepository.FindAll(userID)
}

func (s *todoService) Create(todo domain.Todo) error {
	return s.todoRepository.Create(todo)
}

func (s *todoService) Update(todo domain.Todo) (domain.Todo, error) {
	return s.todoRepository.Update(todo)
}

func (s *todoService) Delete(todoID, userID string) error {
	return s.todoRepository.Delete(todoID, userID)
}
