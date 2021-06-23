package persistence

import (
	"github.com/rubinthomasdev/go-employee/employee/domain"
)

type EmployeeDataMapper struct{}

func (e EmployeeDataMapper) MapToDomain(empEntity EmployeeEntity) domain.Employee {
	return domain.Employee{
		EmployeeID: domain.EmployeeID{ID: empEntity.EmployeeID},
		Name:       domain.EmployeeName{FirstName: empEntity.FirstName, LastName: empEntity.LastName},
		BaseSalary: domain.Money{Amount: empEntity.BaseSalary},
		Bonus:      domain.Money{Amount: float64(empEntity.Bonus)},
	}
}
