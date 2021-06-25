package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rubinthomasdev/go-employee/employee/adapter/in/web"
	"github.com/rubinthomasdev/go-employee/employee/adapter/out/db/postgres"
	"github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"
	"github.com/rubinthomasdev/go-employee/employee/application/service"
)

func main() {

	// ** inmem start **
	// create DB connection
	// inMemDB := inmemdb.InMemDB{}
	// create repo
	// getEmpRepo := inmemdb.InMemRepository{Db: &inMemDB}
	// getEmpRepo.Initialize()
	// ** inmem end **

	// ** postgres start **
	// create DB connection
	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// create repo
	getEmpRepo := postgres.PostgresRepository{DB: db}
	// ** postgres end **

	// create adapter
	getEmpAdapter := persistence.EmployeePersistenceAdapter{EmployeeRepo: getEmpRepo, Mapper: persistence.EmployeeDataMapper{}}

	// create service
	getEmpSvc := service.GetEmployeeDetailsService{LoadEmployeeDataPort: getEmpAdapter}

	// create handler
	getEmpHandler := web.GetEmployeeDataHandler{GetEmployeeDataUseCase: getEmpSvc}

	//http handler registration
	// http.HandleFunc("/api/v1/employees/", getEmpHandler.GetEmployeeDetails)

	// gorilla mux start
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/employees/{empid}", getEmpHandler.GetEmployeeDetails)
	r.HandleFunc("/api/v1/employees", getEmpHandler.GetEmployeeDetails)
	http.Handle("/", r)
	// gorilla mux end

	//start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
