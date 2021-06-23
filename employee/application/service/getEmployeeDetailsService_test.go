package service

import (
	"reflect"
	"testing"

	"github.com/rubinthomasdev/go-employee/employee/application/port/in"
	"github.com/rubinthomasdev/go-employee/employee/domain"
)

type MockLoadEmployeeDataPort struct{}

func (m MockLoadEmployeeDataPort) GetEmployeeDataFromPersistence(id domain.EmployeeID) domain.Employee {
	return domain.Employee{
		EmployeeID: domain.EmployeeID{ID: 1},
		Name: domain.EmployeeName{
			FirstName: "jane",
			LastName:  "doe",
		},
		BaseSalary: domain.Money{Amount: 10.5},
		Bonus:      domain.Money{Amount: 3.8},
	}
}

func (m MockLoadEmployeeDataPort) GetAllEmployees() []domain.Employee {
	return []domain.Employee{
		{
			EmployeeID: domain.EmployeeID{ID: 1},
			Name: domain.EmployeeName{
				FirstName: "john",
				LastName:  "doe",
			},
			BaseSalary: domain.Money{Amount: 10.5},
			Bonus:      domain.Money{Amount: 3.8},
		},
		{
			EmployeeID: domain.EmployeeID{ID: 2},
			Name: domain.EmployeeName{
				FirstName: "jane",
				LastName:  "doe",
			},
			BaseSalary: domain.Money{Amount: 10.5},
			Bonus:      domain.Money{Amount: 3.8},
		},
	}
}

func TestGetEmployeeDetails(t *testing.T) {
	inputID := in.EmployeeQueryID{ID: 1}
	service := GetEmployeeDetailsService{LoadEmployeeDataPort: MockLoadEmployeeDataPort{}}

	want := MockLoadEmployeeDataPort{}.GetEmployeeDataFromPersistence(domain.EmployeeID(inputID))

	got := service.GetEmployeeDetails(inputID)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Get Employee Details failed. Wanted : %v, Got %v", want, got)
	}

}

func TestGetAllEmployeeDetails(t *testing.T) {
	service := GetEmployeeDetailsService{LoadEmployeeDataPort: MockLoadEmployeeDataPort{}}

	want := MockLoadEmployeeDataPort{}.GetAllEmployees()

	got := service.GetAllEmployees()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Get Employee Details failed. Wanted : %v, Got %v", want, got)
	}

}
