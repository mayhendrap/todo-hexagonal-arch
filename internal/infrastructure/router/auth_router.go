package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-hexagonal-arch/internal/application/http"
	"todo-hexagonal-arch/internal/application/http/utils"
	"todo-hexagonal-arch/internal/core/service"
	"todo-hexagonal-arch/internal/infrastructure/mongodb"
)

func NewAuthRouter(db *mongo.Database, collection string, group *gin.RouterGroup) {
	userRepository := mongodb.NewUserRepository(db, collection)
	tokenUtil := utils.NewJwtUtil()
	userService := service.NewAuthService(userRepository, tokenUtil)
	userController := http.NewAuthController(userService)

	group.POST("/register", userController.Register)
	group.POST("/login", userController.Login)
}
