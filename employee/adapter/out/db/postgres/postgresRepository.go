package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"
)

type PostgresRepository struct {
	DB *sql.DB
}

func (p PostgresRepository) FindByID(id int) persistence.EmployeeEntity {
	employeeEntity := persistence.EmployeeEntity{}
	statement := `select id,firstName,lastName,baseSalary,bonus from public.employee where ID=$1;`
	row := p.DB.QueryRow(statement, id)
	switch err := row.Scan(&employeeEntity.EmployeeID, &employeeEntity.FirstName, &employeeEntity.LastName, &employeeEntity.BaseSalary, &employeeEntity.Bonus); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(employeeEntity)
	default:
		fmt.Println("err :", err)
	}
	return employeeEntity
}

func (p PostgresRepository) FindAll() []persistence.EmployeeEntity {
	employees := []persistence.EmployeeEntity{}
	statement := `select id,firstName,lastName,baseSalary,bonus from public.employee limit 20`
	rows, err := p.DB.Query(statement)
	if err != nil {
		log.Println("Error fetching employees data. ", err)
		return employees
	}
	defer rows.Close()

	for rows.Next() {
		employeeEntity := persistence.EmployeeEntity{}
		switch err = rows.Scan(&employeeEntity.EmployeeID, &employeeEntity.FirstName, &employeeEntity.LastName, &employeeEntity.BaseSalary, &employeeEntity.Bonus); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(employeeEntity)
			employees = append(employees, employeeEntity)
		default:
			fmt.Println("err :", err)
		}
	}
	return employees
}
