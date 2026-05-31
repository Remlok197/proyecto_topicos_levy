package routes

import (
	_ "employees-api/docs"
	"employees-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configura todos los endpoints de nuestra API
func SetupRoutes(router *gin.Engine, empHandler *handlers.EmployeeHandler) {
	// Redirigir la raíz (/) a la documentación
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})

	// Redirigir /docs
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})

	// Endpoint de Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		api.GET("/employees", empHandler.GetEmployees)
		api.GET("/employees/:id", empHandler.GetEmployee)
		api.POST("/employees", empHandler.CreateEmployee)
		api.PUT("/employees/:id", empHandler.UpdateEmployee)
		api.DELETE("/employees/:id", empHandler.DeleteEmployee)
	}
}
