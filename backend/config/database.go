package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// usuario:contraseña@tcp(servidor:puerto)/nombre_bd
	// parseTime=True es para que GORM traduzca fechas de MySQL a structs de Go
	dsn := "root:Qwertyz..@tcp(127.0.0.1:3306)/employees_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	fmt.Println("Conexion exitosa a MySQL")
	DB = database
}