package in

import "github.com/rubinthomasdev/go-employee/employee/domain"

type GetEmployeeDetailsQuery interface {
	GetEmployeeDetails(id EmployeeQueryID) domain.Employee
	GetAllEmployees() []domain.Employee
}
