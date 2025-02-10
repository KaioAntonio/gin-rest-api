package routes

import (
	"github.com/KaioAntonio/gin-rest-api/controllers"
	_ "github.com/KaioAntonio/gin-rest-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAulo)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.Run(":8081")
}
