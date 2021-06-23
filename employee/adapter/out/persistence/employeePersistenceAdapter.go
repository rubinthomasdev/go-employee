package persistence

import "github.com/rubinthomasdev/go-employee/employee/domain"

type EmployeePersistenceAdapter struct {
	EmployeeRepo EmployeeRepository
	Mapper       EmployeeDataMapper
}

func (e EmployeePersistenceAdapter) GetEmployeeDataFromPersistence(id domain.EmployeeID) domain.Employee {
	empEntity := e.EmployeeRepo.FindByID(id.ID)
	return e.Mapper.MapToDomain(empEntity)
}

func (e EmployeePersistenceAdapter) GetAllEmployees() []domain.Employee {
	empDomainSlice := []domain.Employee{}
	for _, empEntity := range e.EmployeeRepo.FindAll() {
		empDomainSlice = append(empDomainSlice, e.Mapper.MapToDomain(empEntity))
	}
	return empDomainSlice
}
