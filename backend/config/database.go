package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDatabase inicializa la conexión con la base de datos basándose en variables de entorno.
func ConnectDatabase() (*gorm.DB, error) {
	// Carga variables desde el archivo .env si existe localmente
	loadEnv(".env")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "employees_db")

	if dbUser == "" || dbPass == "" {
		return nil, fmt.Errorf("las variables de entorno obligatorias DB_USER o DB_PASS, no están configuradas")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al abrir conexión de base de datos: %w", err)
	}

	fmt.Println("Conexión exitosa a MySQL")
	return database, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// loadEnv lee un archivo tipo .env y lo carga en las variables de entorno del sistema.
func loadEnv(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return // Silencioso: si no existe (ej. en producción/Docker), se asumen variables de entorno del sistema
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			os.Setenv(key, val)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("advertencia: error al leer el archivo .env: %v", err)
	}
}
