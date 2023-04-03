package dto

import "todo-hexagonal-arch/internal/core/domain"

type TodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

func ToTodoCreate(request TodoRequest, userID string) domain.Todo {
	return domain.Todo{
		Title:   request.Title,
		Content: request.Content,
		UserID:  userID,
	}
}

func ToTodoUpdate(request TodoRequest, todoID, userID string) domain.Todo {
	return domain.Todo{
		ID:      todoID,
		Title:   request.Title,
		Content: request.Content,
		UserID:  userID,
	}
}

func ToTodoResponse(todo domain.Todo) TodoResponse {
	return TodoResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Content: todo.Content,
		UserID:  todo.UserID,
	}
}
