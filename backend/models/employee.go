package models

import "time"

type Employee struct {
	EmpNo     int       `gorm:"primaryKey;column:emp_no" json:"emp_no"`
	BirthDate time.Time `gorm:"column:birth_date" json:"birth_date"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Gender    string    `gorm:"column:gender" json:"gender"`
	HireDate  time.Time `gorm:"column:hire_date" json:"hire_date"`
}

type Department struct {
	DeptNo   string `gorm:"primaryKey;column:dept_no" json:"dept_no"`
	DeptName string `gorm:"column:dept_name" json:"dept_name"`
}

type DeptEmp struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no" json:"emp_no"`
	DeptNo   string    `gorm:"primaryKey;column:dept_no" json:"dept_no"`
	FromDate time.Time `gorm:"primaryKey;column:from_date" json:"from_date"`
	ToDate   time.Time `gorm:"column:to_date" json:"to_date"`
}

func (DeptEmp) TableName() string {
	return "dept_emp"
}

type Title struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no" json:"emp_no"`
	Title    string    `gorm:"primaryKey;column:title" json:"title"`
	FromDate time.Time `gorm:"primaryKey;column:from_date" json:"from_date"`
	ToDate   time.Time `gorm:"column:to_date" json:"to_date"`
}

type Salary struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no" json:"emp_no"`
	Salary   int       `gorm:"column:salary" json:"salary"`
	FromDate time.Time `gorm:"primaryKey;column:from_date" json:"from_date"`
	ToDate   time.Time `gorm:"column:to_date" json:"to_date"`
}

// Este modelo es el DTO, lo vamos a usar para el formulario de altas, recibe todos los datos jeje
type NewHireRequest struct {
	EmpNo     int    `json:"emp_no"`
	BirthDate string `json:"birth_date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	HireDate  string `json:"hire_date"`
	DeptNo    string `json:"dept_no"`
	Title     string `json:"title"`
	Salary    int    `json:"salary"`
}
