package routes

import (
	_ "employees-api/docs"
	"employees-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CORSMiddleware permite peticiones cruzadas desde el frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // O ajusta al puerto de Vite (ej. http://localhost:5173)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRoutes configura todos los endpoints de nuestra API
func SetupRoutes(router *gin.Engine, empHandler *handlers.EmployeeHandler) {
	// Aplicar CORS
	router.Use(CORSMiddleware())
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
