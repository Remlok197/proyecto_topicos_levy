package storage

import (
	"employees-api/internal/models"
	"time"

	"gorm.io/gorm"
)

// EmployeeStorage define las operaciones de base de datos disponibles
type EmployeeStorage interface {
	GetPaginated(limit, offset int) ([]EmployeeSummaryRow, error)
	GetByID(id int) (models.Employee, []models.Title, []models.Salary, error)
	CreateNewHire(emp models.Employee, deptEmp models.DeptEmp, title models.Title, salary models.Salary) error
	UpdateHistory(id int, today time.Time, newSalary *models.Salary, newTitle *models.Title, newDeptEmp *models.DeptEmp) error
	TerminateEmployee(id int, today time.Time) error
}

// gormEmployeeStorage implementa la interfaz usando GORM
type gormEmployeeStorage struct {
	db *gorm.DB
}

// NewEmployeeStorage crea una instancia del almacenamiento
func NewEmployeeStorage(db *gorm.DB) EmployeeStorage {
	return &gormEmployeeStorage{db: db}
}

// Estructura temporal interna para mapear el Query RAW de paginación
type EmployeeSummaryRow struct {
	EmpNo     int
	FirstName string
	LastName  string
	Title     string
	Salary    int
}

// --- Implementación de los métodos ---

func (s *gormEmployeeStorage) GetPaginated(limit, offset int) ([]EmployeeSummaryRow, error) {
	var results []EmployeeSummaryRow

	query := `
        SELECT e.emp_no, e.first_name, e.last_name, t.title, s.salary
        FROM employees e
        JOIN titles t ON e.emp_no = t.emp_no
        JOIN salaries s ON e.emp_no = s.emp_no
        WHERE t.to_date = '9999-01-01' AND s.to_date = '9999-01-01'
        LIMIT ? OFFSET ?
    `
	err := s.db.Raw(query, limit, offset).Scan(&results).Error
	return results, err
}

func (s *gormEmployeeStorage) GetByID(id int) (models.Employee, []models.Title, []models.Salary, error) {
	var employee models.Employee
	var titles []models.Title
	var salaries []models.Salary

	if err := s.db.First(&employee, id).Error; err != nil {
		return employee, nil, nil, err
	}

	if err := s.db.Where("emp_no = ?", id).Order("from_date ASC").Find(&titles).Error; err != nil {
		return employee, nil, nil, err
	}

	if err := s.db.Where("emp_no = ?", id).Order("from_date ASC").Find(&salaries).Error; err != nil {
		return employee, nil, nil, err
	}

	return employee, titles, salaries, nil
}

func (s *gormEmployeeStorage) CreateNewHire(emp models.Employee, deptEmp models.DeptEmp, title models.Title, salary models.Salary) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&emp).Error; err != nil {
			return err
		}
		if err := tx.Create(&deptEmp).Error; err != nil {
			return err
		}
		if err := tx.Create(&title).Error; err != nil {
			return err
		}
		if err := tx.Create(&salary).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *gormEmployeeStorage) UpdateHistory(id int, today time.Time, newSalary *models.Salary, newTitle *models.Title, newDeptEmp *models.DeptEmp) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		infinityDate, _ := time.Parse("2006-01-02", "9999-01-01")

		if newSalary != nil {
			if err := tx.Model(&models.Salary{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
				return err
			}
			newSalary.EmpNo = id
			newSalary.FromDate = today
			newSalary.ToDate = infinityDate
			if err := tx.Create(newSalary).Error; err != nil {
				return err
			}
		}

		if newTitle != nil {
			if err := tx.Model(&models.Title{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
				return err
			}
			newTitle.EmpNo = id
			newTitle.FromDate = today
			newTitle.ToDate = &infinityDate
			if err := tx.Create(newTitle).Error; err != nil {
				return err
			}
		}

		if newDeptEmp != nil {
			if err := tx.Model(&models.DeptEmp{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
				return err
			}
			newDeptEmp.EmpNo = id
			newDeptEmp.FromDate = today
			newDeptEmp.ToDate = infinityDate
			if err := tx.Create(newDeptEmp).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *gormEmployeeStorage) TerminateEmployee(id int, today time.Time) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Title{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Salary{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.DeptEmp{}).Where("emp_no = ? AND to_date = ?", id, "9999-01-01").Update("to_date", today).Error; err != nil {
			return err
		}
		return nil
	})
}
