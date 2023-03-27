package main

import (
	"github.com/gin-gonic/gin"
	"todo-hexagonal-arch/internal/infrastructure/mongodb"
	"todo-hexagonal-arch/internal/infrastructure/router"
)

func main() {

	dbAdapter := mongodb.NewMongoDBAdapter("todo-list", "mongodb://localhost:27017")
	defer dbAdapter.CloseDBConnection()

	engine := gin.Default()

	router.Setup(dbAdapter.GetDB(), engine)

	engine.Run(":8080")
}
