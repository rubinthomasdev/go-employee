package web

type GetEmployeeDTO struct {
	EmployeeID  int     `json:"employeeID"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	BaseSalary  float64 `json:"baseSalary"`
	Bonus       float64 `json:"bonus"`
	TotalSalary float64 `json:"totalSalary"`
}
