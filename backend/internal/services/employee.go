package services

import (
	"employees-api/internal/models"
	"employees-api/internal/schemas"
	"employees-api/internal/storage"
	"fmt"
	"time"
)

// EmployeeService define las acciones de negocio para empleados
type EmployeeService interface {
	GetEmployees(page, limit int) (schemas.PaginatedEmployeesResponse, error)
	GetEmployeeByID(id int) (schemas.EmployeeDetailResponse, error)
	RegisterEmployee(req schemas.NewHireRequest) error
	UpdateEmployee(id int, req schemas.UpdateEmployeeRequest) error
	TerminateEmployee(id int) error
}

// employeeServiceImpl implementa la interfaz de negocio
type employeeServiceImpl struct {
	store storage.EmployeeStorage
}

// NewEmployeeService inicializa el servicio inyectando su almacenamiento
func NewEmployeeService(store storage.EmployeeStorage) EmployeeService {
	return &employeeServiceImpl{store: store}
}

// --- Implementación de los métodos ---

func (s *employeeServiceImpl) GetEmployees(page, limit int) (schemas.PaginatedEmployeesResponse, error) {
	// Evitar límites absurdos que puedan crashear la API
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	rows, err := s.store.GetPaginated(limit, offset)
	if err != nil {
		return schemas.PaginatedEmployeesResponse{}, err
	}

	// Traducir las filas crudas del repositorio a DTOs de salida para el cliente
	summaries := make([]schemas.EmployeeSummaryResponse, len(rows))
	for i, row := range rows {
		summaries[i] = schemas.EmployeeSummaryResponse{
			EmployeeNumber: row.EmpNo,
			FirstName:      row.FirstName,
			LastName:       row.LastName,
			JobTitle:       row.Title,
			Salary:         row.Salary,
		}
	}

	return schemas.PaginatedEmployeesResponse{
		Page:  page,
		Limit: limit,
		Data:  summaries,
	}, nil
}

func (s *employeeServiceImpl) GetEmployeeByID(id int) (schemas.EmployeeDetailResponse, error) {
	emp, titles, salaries, err := s.store.GetByID(id)
	if err != nil {
		return schemas.EmployeeDetailResponse{}, err
	}

	layout := "2006-01-02"

	// Traducir información personal
	personalInfo := schemas.EmployeeInfoResponse{
		EmpNo:     emp.EmpNo,
		BirthDate: emp.BirthDate.Format(layout),
		FirstName: emp.FirstName,
		LastName:  emp.LastName,
		Gender:    emp.Gender,
		HireDate:  emp.HireDate.Format(layout),
	}

	// Traducir el historial de puestos
	titlesHistory := make([]schemas.TitleResponse, len(titles))
	for i, t := range titles {
		toDateStr := ""
		if t.ToDate != nil {
			toDateStr = t.ToDate.Format(layout)
		}
		titlesHistory[i] = schemas.TitleResponse{
			Title:    t.Title,
			FromDate: t.FromDate.Format(layout),
			ToDate:   toDateStr,
		}
	}

	// Traducir el historial de salarios
	salariesHistory := make([]schemas.SalaryResponse, len(salaries))
	for i, sal := range salaries {
		salariesHistory[i] = schemas.SalaryResponse{
			Salary:   sal.Salary,
			FromDate: sal.FromDate.Format(layout),
			ToDate:   sal.ToDate.Format(layout),
		}
	}

	return schemas.EmployeeDetailResponse{
		PersonalInfo:  personalInfo,
		TitlesHistory: titlesHistory,
		SalaryHistory: salariesHistory,
	}, nil
}

func (s *employeeServiceImpl) RegisterEmployee(req schemas.NewHireRequest) error {
	layout := "2006-01-02"

	// Conversión de fechas
	birthDate, err := time.Parse(layout, req.BirthDate)
	if err != nil {
		return fmt.Errorf("fecha de nacimiento inválida (formato YYYY-MM-DD): %w", err)
	}
	hireDate, err := time.Parse(layout, req.HireDate)
	if err != nil {
		return fmt.Errorf("fecha de contratación inválida (formato YYYY-MM-DD): %w", err)
	}
	infinityDate, _ := time.Parse(layout, "9999-01-01")

	// Crear modelos de base de datos
	emp := models.Employee{
		EmpNo:     req.EmpNo,
		BirthDate: birthDate,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
		HireDate:  hireDate,
	}

	deptEmp := models.DeptEmp{
		EmpNo:    req.EmpNo,
		DeptNo:   req.DeptNo,
		FromDate: hireDate,
		ToDate:   infinityDate,
	}

	title := models.Title{
		EmpNo:    req.EmpNo,
		Title:    req.Title,
		FromDate: hireDate,
		ToDate:   &infinityDate,
	}

	salary := models.Salary{
		EmpNo:    req.EmpNo,
		Salary:   req.Salary,
		FromDate: hireDate,
		ToDate:   infinityDate,
	}

	return s.store.CreateNewHire(emp, deptEmp, title, salary)
}

func (s *employeeServiceImpl) UpdateEmployee(id int, req schemas.UpdateEmployeeRequest) error {
	today := time.Now()

	var newSalary *models.Salary
	var newTitle *models.Title
	var newDeptEmp *models.DeptEmp

	// Mapear solo los campos que vinieron en la petición
	if req.Salary != nil {
		newSalary = &models.Salary{
			Salary: *req.Salary,
		}
	}

	if req.Title != nil {
		newTitle = &models.Title{
			Title: *req.Title,
		}
	}

	if req.DeptNo != nil {
		newDeptEmp = &models.DeptEmp{
			DeptNo: *req.DeptNo,
		}
	}

	// Si no enviaron ningún cambio, retornamos temprano sin tocar la BD
	if newSalary == nil && newTitle == nil && newDeptEmp == nil {
		return nil
	}

	return s.store.UpdateHistory(id, today, newSalary, newTitle, newDeptEmp)
}

func (s *employeeServiceImpl) TerminateEmployee(id int) error {
	today := time.Now()
	return s.store.TerminateEmployee(id, today)
}
