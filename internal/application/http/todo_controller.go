package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-hexagonal-arch/internal/application/http/dto"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type todoController struct {
	todoService interfaces.TodoService
}

func NewTodoController(todoService interfaces.TodoService) *todoController {
	return &todoController{todoService: todoService}
}

func (tc *todoController) FindOne(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.GetString("x-user-id")

	result, err := tc.todoService.FindOne(todoID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	response := dto.TodoResponse{
		ID:      result.ID,
		Title:   result.Title,
		Content: result.Content,
		UserID:  result.UserID,
	}

	c.JSON(http.StatusOK, response)
}

func (tc *todoController) Create(c *gin.Context) {
	var request dto.TodoRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	todo := dto.ToTodoCreate(request, userID)

	err := tc.todoService.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToResponse("created", 201, nil))
}

func (tc *todoController) Update(c *gin.Context) {
	var request dto.TodoRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	todoID := c.Param("id")
	userID := c.GetString("x-user-id")

	todo := dto.ToTodoUpdate(request, todoID, userID)

	updatedTodo, err := tc.todoService.Update(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ToResponse("error", 500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.ToResponse("success", 200, updatedTodo))
}

func (tc *todoController) Delete(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.GetString("x-user-id")

	err := tc.todoService.Delete(todoID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "deleted"})
}
