package main

import (
	"employees-api/config"
	"employees-api/internal/handlers"
	"employees-api/internal/routes"
	"employees-api/internal/services"
	"employees-api/internal/storage"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Employees API
// @version 1.0
// @description API REST para la gestión de empleados (MySQl test_db).
// @host localhost:8080
// @BasePath /

func main() {
	// 1. Conectar a la base de datos
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("No se pudo iniciar la base de datos: ", err)
	}

	// 2. Inyección de dependencias (de abajo hacia arriba)
	empStore := storage.NewEmployeeStorage(db)
	empService := services.NewEmployeeService(empStore)
	empHandler := handlers.NewEmployeeHandler(empService)

	// 3. Configurar servidor HTTP y enrutar
	r := gin.Default()
	routes.SetupRoutes(r, empHandler)

	fmt.Println("Arrancando server en http://localhost:8080")
	r.Run(":8080")
}
