package persistence

type EmployeeEntity struct {
	EmployeeID int     `json:"employeeID"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	BaseSalary float64 `json:"baseSalary"`
	Bonus      float64 `json:"bonus"`
}