package inmemdb

import "github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"

type InMemDB struct {
	Employees map[int]persistence.EmployeeEntity
}
