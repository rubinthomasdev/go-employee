package service

import (
	"github.com/rubinthomasdev/go-employee/employee/application/port/in"
	"github.com/rubinthomasdev/go-employee/employee/application/port/out"
	"github.com/rubinthomasdev/go-employee/employee/domain"
)

type GetEmployeeDetailsService struct {
	LoadEmployeeDataPort out.LoadEmployeeDataPort
}

func (g GetEmployeeDetailsService) GetEmployeeDetails(inputID in.EmployeeQueryID) domain.Employee {
	return g.LoadEmployeeDataPort.GetEmployeeDataFromPersistence(domain.EmployeeID{ID: inputID.ID})
}

func (g GetEmployeeDetailsService) GetAllEmployees() []domain.Employee {
	return g.LoadEmployeeDataPort.GetAllEmployees()
}
