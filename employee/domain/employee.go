package domain

type Employee struct {
	EmployeeID EmployeeID   `json:"employeeID"`
	Name       EmployeeName `json:"employeeName"`
	BaseSalary Money        `json:"baseSalary"`
	Bonus      Money        `json:"bonus"`
}

func (e Employee) CalculateTotalSalary() Money {
	return e.BaseSalary.Add(e.Bonus)
}
