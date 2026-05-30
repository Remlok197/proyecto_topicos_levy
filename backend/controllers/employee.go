package controllers

import (
	"fmt"
	"net/http"
	"proyecto-topicos-backend/config"
	"proyecto-topicos-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 1. GET
func GetEmployees(c *gin.Context) {
	type EmployeeResponse struct {
		EmpNo     int    `json:"employeeNumber"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Title     string `json:"jobTitle"`
		Salary    int    `json:"salary"`
	}

	var results []EmployeeResponse

	// JOIN para traer el empleado y puesto actual
	query := `
		SELECT e.emp_no, e.first_name, e.last_name, t.title, s.salary
		FROM employees e
		JOIN titles t ON e.emp_no = t.emp_no
		WHERE t.to_date = '9999-01-01'
	`

	config.DB.Raw(query).Scan(&results)
	c.JSON(http.StatusOK, results)
}

// 2. GET by ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	var titles []models.Title
	var salaries []models.Salary

	if err := config.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}

	config.DB.Where("emp_no = ?", id).Order("from_date ASC").Find(&titles)
	config.DB.Where("emp_no = ?", id).Order("from_date ASC").Find(&salaries)

	c.JSON(http.StatusOK, gin.H{
		"personal_info":  employee,
		"titles_history": titles,
		"salary_history": salaries,
	})
}

// 3. POST
func CreateEmployee(c *gin.Context) {
	var req models.NewHireRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos en el formulario"})
		return
	}

	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, req.BirthDate)
	hireDate, _ := time.Parse(layout, req.HireDate)
	infinityDate, _ := time.Parse(layout, "9999-01-01")

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		emp := models.Employee{
			EmpNo:     req.EmpNo,
			BirthDate: birthDate,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    req.Gender,
			HireDate:  hireDate,
		}
		if err := tx.Create(&emp).Error; err != nil {
			return err
		}

		deptEmp := models.DeptEmp{
			EmpNo:    req.EmpNo,
			DeptNo:   req.DeptNo,
			FromDate: hireDate,
			ToDate:   infinityDate,
		}
		if err := tx.Create(&deptEmp).Error; err != nil {
			return err
		}

		title := models.Title{
			EmpNo:    req.EmpNo,
			Title:    req.Title,
			FromDate: hireDate,
			ToDate:   infinityDate,
		}
		if err := tx.Create(&title).Error; err != nil {
			return err
		}

		salary := models.Salary{
			EmpNo:    req.EmpNo,
			Salary:   req.Salary,
			FromDate: hireDate,
			ToDate:   infinityDate,
		}
		if err := tx.Create(&salary).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al registrar al empleado, la operación fue cancelada: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Empleado registrado con éxito"})
}

// 4. PUT
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	today := time.Now()
	infinityDate, _ := time.Parse("2006-01-02", "9999-01-01")

	// Mapa generico para el cambio
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Ascenso
		if newSalary, ok := updateData["salary"]; ok {
			tx.Model(&models.Salary{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
			tx.Create(&models.Salary{EmpNo: convertToInt(id), Salary: int(newSalary.(float64)), FromDate: today, ToDate: infinityDate})
		}

		// Ascenso
		if newTitle, ok := updateData["title"]; ok {
			tx.Model(&models.Title{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
			tx.Create(&models.Title{EmpNo: convertToInt(id), Title: newTitle.(string), FromDate: today, ToDate: infinityDate})
		}

		// Transferencia
		if newDept, ok := updateData["dept_no"]; ok {
			tx.Model(&models.DeptEmp{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
			tx.Create(&models.DeptEmp{EmpNo: convertToInt(id), DeptNo: newDept.(string), FromDate: today, ToDate: infinityDate})
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo al actualizar historial"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Historial actualizado correctamente"})
}

// 5. DELETE
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	today := time.Now()

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Cerramos ciclo en todas las tablas
		tx.Model(&models.Title{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
		tx.Model(&models.Salary{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
		tx.Model(&models.DeptEmp{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo dar de baja al empleado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Empleado dado de baja exitosamente (baja lógica registrada)"})
}

// Función para convertir el ID string de Gin a int para GORM
func convertToInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
