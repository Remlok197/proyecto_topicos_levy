package schemas

// EmployeeSummaryResponse es el DTO para listados (GET /employees)
type EmployeeSummaryResponse struct {
	EmployeeNumber int    `json:"employeeNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	JobTitle       string `json:"jobTitle"`
	Salary         int    `json:"salary"`
}

// Metadata
type PaginatedEmployeesResponse struct {
	Page  int                       `json:"page"`
	Limit int                       `json:"limit"`
	Data  []EmployeeSummaryResponse `json:"data"`
}

// EmployeeDetailResponse (GET /employees/:id)
type EmployeeDetailResponse struct {
	PersonalInfo  EmployeeInfoResponse `json:"personal_info"`
	TitlesHistory []TitleResponse      `json:"titles_history"`
	SalaryHistory []SalaryResponse     `json:"salary_history"`
}

// Sub-structs
type EmployeeInfoResponse struct {
	EmpNo     int    `json:"emp_no"`
	BirthDate string `json:"birth_date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	HireDate  string `json:"hire_date"`
}

type TitleResponse struct {
	Title    string `json:"title"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type SalaryResponse struct {
	Salary   int    `json:"salary"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}
