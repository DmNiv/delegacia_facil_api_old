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
		// delegaciasRoutes.GET("/:id", controllers.GetDelegaciaByID)
		// delegaciasRoutes.GET("/horario24h", controllers.GetDelegaciasByHorario)
		// delegaciasRoutes.GET("/tipo", controllers.GetDelegaciasByTipo)
	}

	return r
}
