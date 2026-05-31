package schemas

type NewHireRequest struct {
	EmpNo     int    `json:"emp_no" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required,datetime=2006-01-02"`
	FirstName string `json:"first_name" binding:"required,max=14"`
	LastName  string `json:"last_name" binding:"required,max=16"`
	Gender    string `json:"gender" binding:"required,oneof=M F"`
	HireDate  string `json:"hire_date" binding:"required,datetime=2006-01-02"`
	DeptNo    string `json:"dept_no" binding:"required,len=4"`
	Title     string `json:"title" binding:"required,max=50"`
	Salary    int    `json:"salary" binding:"required,gt=0"`
}

type UpdateEmployeeRequest struct {
	Salary *int    `json:"salary" binding:"omitempty,gt=0"`
	Title  *string `json:"title" binding:"omitempty,max=50"`
	DeptNo *string `json:"dept_no" binding:"omitempty,len=4"`
}
