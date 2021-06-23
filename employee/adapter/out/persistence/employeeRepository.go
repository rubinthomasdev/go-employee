package persistence

type EmployeeRepository interface {
	FindByID(int) EmployeeEntity
	FindAll() []EmployeeEntity
}
