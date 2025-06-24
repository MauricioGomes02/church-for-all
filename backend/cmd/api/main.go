package main

import (
	usecases "church-for-all/internal/application/use_cases"
	"church-for-all/internal/infra/api/handlers"
	"church-for-all/internal/infra/data/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	poolConnectionManager := mysql.NewPoolConnectionManager()
	createChurchMySQL := mysql.NewCreateChurchMySQL(poolConnectionManager.Database)
	getChurchMySQL := mysql.NewGetChurchsMySQL(poolConnectionManager.Database)
	createChurchUseCase := usecases.NewCreateChurchUseCase(createChurchMySQL)
	getChurchsUseCase := usecases.NewGetChurchsUseCase(getChurchMySQL)
	createChurchHandler := handlers.NewCreateChurchHandler(createChurchUseCase)
	getChurchsHandler := handlers.NewGetChurchsHandler(getChurchsUseCase)

	churchsGroup := engine.Group("churchs")
	churchsGroup.POST("", createChurchHandler.Create)
	churchsGroup.GET("", getChurchsHandler.Get)

	// Responsável por iniciar o servidor e escutar na porta 8080
	engine.Run(":8080")
}
