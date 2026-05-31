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

// GetEmployees godoc
// @Summary Obtiene la lista de empleados
// @Description Retorna los empleados con su salario y título actual, con paginación opcional
// @Tags employees
// @Produce json
// @Param page query int false "Página" default(1)
// @Param limit query int false "Límite de empleados" default(50)
// @Success 200 {object} schemas.PaginatedEmployeesResponse "Lista de empleados paginada"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /employees [get]
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

// GetEmployee godoc
// @Summary Obtiene un empleado por su ID
// @Description Retorna un solo empleado y sus detalles históricos
// @Tags employees
// @Produce json
// @Param id path int true "ID del Empleado"
// @Success 200 {object} schemas.EmployeeDetailResponse "Detalles del empleado"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Empleado no encontrado"
// @Router /employees/{id} [get]
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

// CreateEmployee godoc
// @Summary Registra un nuevo empleado
// @Description Da de alta a un empleado en la base de datos
// @Tags employees
// @Accept json
// @Produce json
// @Param request body schemas.NewHireRequest true "Datos del empleado"
// @Success 201 {object} map[string]string "Empleado registrado"
// @Failure 400 {object} map[string]string "Datos inválidos"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /employees [post]
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

// UpdateEmployee godoc
// @Summary Actualiza el historial laboral
// @Description Actualiza el salario y/o título de un empleado, manteniendo el historial (Type 2 SCD)
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "ID del Empleado"
// @Param request body schemas.UpdateEmployeeRequest true "Nuevos datos a actualizar"
// @Success 200 {object} map[string]string "Historial actualizado"
// @Failure 400 {object} map[string]string "JSON o ID inválido"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /employees/{id} [put]
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

// DeleteEmployee godoc
// @Summary Baja lógica de un empleado
// @Description Da de baja a un empleado y cierra sus historiales de salario/título
// @Tags employees
// @Produce json
// @Param id path int true "ID del Empleado"
// @Success 200 {object} map[string]string "Baja lógica exitosa"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /employees/{id} [delete]
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
