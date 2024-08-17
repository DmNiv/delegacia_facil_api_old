package routes

import (
	"delegaciafacil/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	delegaciasRoutes := r.Group("/delegacias")
	{
		delegaciasRoutes.GET("", controllers.GetDelegacias)
		delegaciasRoutes.GET("/horario", controllers.GetDelegaciasPorHorario)
	}
	return r
}