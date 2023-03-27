package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-hexagonal-arch/internal/application/http/middleware"
	"todo-hexagonal-arch/internal/application/http/utils"
	"todo-hexagonal-arch/internal/core/domain"
)

func Setup(db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewAuthRouter(db, domain.UserCollection, publicRouter)

	tokenUtil := utils.NewJwtUtil()
	authMiddleware := middleware.NewAuthMiddleware(tokenUtil)

	privateRouter := gin.Group("")
	privateRouter.Use(authMiddleware.AuthMiddleware())
	NewTodoRouter(db, domain.TodoCollection, privateRouter)
}
