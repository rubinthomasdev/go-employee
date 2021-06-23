package inmemdb

import "github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"

type InMemRepository struct {
	Db *InMemDB
}

func (i InMemRepository) FindByID(id int) persistence.EmployeeEntity {
	return i.Db.Employees[id]
}

func (i InMemRepository) FindAll() []persistence.EmployeeEntity {
	employees := []persistence.EmployeeEntity{}
	for _, emp := range i.Db.Employees {
		employees = append(employees, emp)
	}
	return employees
}

func (i InMemRepository) Initialize() {
	e1 := persistence.EmployeeEntity{
		EmployeeID: 1,
		FirstName:  "john",
		LastName:   "doe",
		BaseSalary: 100.0,
		Bonus:      12.5,
	}

	e2 := persistence.EmployeeEntity{
		EmployeeID: 2,
		FirstName:  "jane",
		LastName:   "doe",
		BaseSalary: 100.0,
		Bonus:      12.57,
	}

	i.Db.Employees = make(map[int]persistence.EmployeeEntity)
	i.Db.Employees[1] = e1
	i.Db.Employees[2] = e2

}
