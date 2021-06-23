package main

import (
	"log"
	"net/http"

	"github.com/rubinthomasdev/go-employee/employee/adapter/in/web"
	inmemdb "github.com/rubinthomasdev/go-employee/employee/adapter/out/db/inMemDB"
	"github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"
	"github.com/rubinthomasdev/go-employee/employee/application/service"
)

func main() {

	// create DB
	inMemDB := inmemdb.InMemDB{}

	// create repo
	getEmpRepo := inmemdb.InMemRepository{Db: &inMemDB}
	getEmpRepo.Initialize()

	// create adapter
	getEmpAdapter := persistence.EmployeePersistenceAdapter{EmployeeRepo: getEmpRepo, Mapper: persistence.EmployeeDataMapper{}}

	// create service
	getEmpSvc := service.GetEmployeeDetailsService{LoadEmployeeDataPort: getEmpAdapter}

	// create handler
	getEmpHandler := web.GetEmployeeDataHandler{GetEmployeeDataUseCase: getEmpSvc}

	//http handler registration
	http.HandleFunc("/api/v1/employees/", getEmpHandler.GetEmployeeDetails)

	//start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
