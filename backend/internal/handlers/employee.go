package handlers

import (
	"net/http"
	"employees-api/internal/schemas"
	"employees-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeHandler expone los endpoints HTTP para la gestión de empleados
type EmployeeHandler struct {
	service services.EmployeeService
}

// NewEmployeeHandler inicializa el handler con su correspondiente servicio
func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// GetEmployees maneja: GET /employees (con paginación opcional)
func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	res, err := h.service.GetEmployees(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetEmployee maneja: GET /employees/:id
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	res, err := h.service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado o error de lectura"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateEmployee maneja: POST /employees
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req schemas.NewHireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos en el formulario: " + err.Error()})
		return
	}

	if err := h.service.RegisterEmployee(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar al empleado: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Empleado registrado con éxito"})
}

// UpdateEmployee maneja: PUT /employees/:id
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	var req schemas.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	if err := h.service.UpdateEmployee(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo al actualizar historial: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Historial actualizado correctamente"})
}

// DeleteEmployee maneja: DELETE /employees/:id (Baja lógica)
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	if err := h.service.TerminateEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo dar de baja al empleado: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Empleado dado de baja exitosamente (baja lógica registrada)"})
}
