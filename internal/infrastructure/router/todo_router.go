package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-hexagonal-arch/internal/application/http"
	"todo-hexagonal-arch/internal/core/service"
	"todo-hexagonal-arch/internal/infrastructure/mongodb"
)

func NewTodoRouter(db *mongo.Database, collection string, group *gin.RouterGroup) {
	todoRepository := mongodb.NewTodoRepository(db, collection)
	todoService := service.NewTodoService(todoRepository)
	todoController := http.NewTodoController(todoService)

	group.GET("/todos/:id", todoController.FindOne)
	group.POST("/todos", todoController.Create)
	group.PATCH("/todos/:id", todoController.Update)
	group.DELETE("/todos/:id", todoController.Delete)
}
