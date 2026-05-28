package models

type Employee struct {
	EmployeeNumber int    `gorm:"primaryKey;column:employeeNumber" json:"employeeNumber"`
	LastName       string `gorm:"column:lastName;not null" json:"lastName"`
	FirstName      string `gorm:"column:firstName;not null" json:"firstName"`
	Extension      string `gorm:"column:extension;not null" json:"extension"`
	Email          string `gorm:"column:email;not null" json:"email"`
	OfficeCode     string `gorm:"column:officeCode;not null" json:"officeCode"`
	ReportsTo      *int   `gorm:"column:reportsTo" json:"reportsTo"` 
	JobTitle       string `gorm:"column:jobTitle;not null" json:"jobTitle"`
}