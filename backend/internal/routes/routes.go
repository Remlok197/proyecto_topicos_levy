package routes

import (
	"employees-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todos los endpoints de nuestra API
func SetupRoutes(router *gin.Engine, empHandler *handlers.EmployeeHandler) {
	api := router.Group("/api")
	{
		api.GET("/employees", empHandler.GetEmployees)
		api.GET("/employees/:id", empHandler.GetEmployee)
		api.POST("/employees", empHandler.CreateEmployee)
		api.PUT("/employees/:id", empHandler.UpdateEmployee)
		api.DELETE("/employees/:id", empHandler.DeleteEmployee)
	}
}
