package routes

import (
	"proyecto-topicos-backend/controllers"

	"github.com/gin-gonic/gin"
)

// configura todos los endpoints de nuestra API
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/employees", controllers.GetEmployees)
		api.GET("/employees/:id", controllers.GetEmployee)
		api.POST("/employees", controllers.CreateEmployee)
		api.PUT("/employees/:id", controllers.UpdateEmployee)
		api.DELETE("/employees/:id", controllers.DeleteEmployee)
	}
}