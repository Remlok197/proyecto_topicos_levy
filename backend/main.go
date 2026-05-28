package main

import (
	"fmt"
	"proyecto-topicos-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Server Gin y conexión a MySQL funcionando mi gente",
		})
	})
	fmt.Println("Arrancando server en http://localhost:8080")
	r.Run(":8080")
}