package main

import (
	"github.com/gin-gonic/gin"
	"todo-hexagonal-arch/internal/application/http/utils"
	"todo-hexagonal-arch/internal/infrastructure/mongodb"
	"todo-hexagonal-arch/internal/infrastructure/router"
)

func main() {

	env := utils.NewEnv()

	dbAdapter := mongodb.NewMongoDBAdapter(env.DBName, env.DBUrl)
	defer dbAdapter.CloseDBConnection()

	engine := gin.Default()

	router.Setup(dbAdapter.GetDB(), engine)

	engine.Run(env.AppPort)
}
