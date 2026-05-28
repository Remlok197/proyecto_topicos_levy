package controllers

import (
	"net/http"
	"proyecto-topicos-backend/config"
	"proyecto-topicos-backend/models"

	"github.com/gin-gonic/gin"
)

// 1. GET
func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	config.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

// 2. GET by ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// 3. POST
func CreateEmployee(c *gin.Context) {
	var input models.Employee

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// 4. PUT
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}

	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&employee).Updates(input)
	c.JSON(http.StatusOK, employee)
}

// 5. DELETE
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}

	config.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Empleado eliminado correctamente"})
}