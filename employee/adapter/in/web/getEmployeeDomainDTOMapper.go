package web

import "github.com/rubinthomasdev/go-employee/employee/domain"

type GetEmployeeDomainDTOMapper struct {
}

func (g GetEmployeeDomainDTOMapper) MapDomainToDTO(empDomain domain.Employee) GetEmployeeDTO {
	return GetEmployeeDTO{
		EmployeeID:  empDomain.EmployeeID.ID,
		FirstName:   empDomain.Name.FirstName,
		LastName:    empDomain.Name.LastName,
		BaseSalary:  empDomain.BaseSalary.Amount,
		Bonus:       empDomain.Bonus.Amount,
		TotalSalary: empDomain.CalculateTotalSalary().Amount,
	}
}
