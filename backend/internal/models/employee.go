package models

import "time"

type Employee struct {
	EmpNo     int       `gorm:"primaryKey;column:emp_no"`
	BirthDate time.Time `gorm:"column:birth_date"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Gender    string    `gorm:"column:gender"`
	HireDate  time.Time `gorm:"column:hire_date"`
}

type Department struct {
	DeptNo   string `gorm:"primaryKey;column:dept_no"`
	DeptName string `gorm:"column:dept_name"`
}

type DeptManager struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no"`
	DeptNo   string    `gorm:"primaryKey;column:dept_no"`
	FromDate time.Time `gorm:"column:from_date"`
	ToDate   time.Time `gorm:"column:to_date"`
}

func (DeptManager) TableName() string {
	return "dept_manager"
}

type DeptEmp struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no"`
	DeptNo   string    `gorm:"primaryKey;column:dept_no"`
	FromDate time.Time `gorm:"column:from_date"`
	ToDate   time.Time `gorm:"column:to_date"`
}

func (DeptEmp) TableName() string {
	return "dept_emp"
}

type Title struct {
	EmpNo    int        `gorm:"primaryKey;column:emp_no"`
	Title    string     `gorm:"primaryKey;column:title"`
	FromDate time.Time  `gorm:"primaryKey;column:from_date"`
	ToDate   *time.Time `gorm:"column:to_date"`
}

type Salary struct {
	EmpNo    int       `gorm:"primaryKey;column:emp_no"`
	Salary   int       `gorm:"column:salary"`
	FromDate time.Time `gorm:"primaryKey;column:from_date"`
	ToDate   time.Time `gorm:"column:to_date"`
}
