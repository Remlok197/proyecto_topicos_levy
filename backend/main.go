package main

import (
	"fmt"
	"proyecto-topicos-backend/config"
	"proyecto-topicos-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	routes.SetupRoutes(r)
	fmt.Println("Arrancando server en http://localhost:8080")
	r.Run(":8080")
}