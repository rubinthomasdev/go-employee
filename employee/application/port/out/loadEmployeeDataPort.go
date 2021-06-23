package out

import "github.com/rubinthomasdev/go-employee/employee/domain"

type LoadEmployeeDataPort interface {
	GetEmployeeDataFromPersistence(id domain.EmployeeID) domain.Employee
	GetAllEmployees() []domain.Employee
}
