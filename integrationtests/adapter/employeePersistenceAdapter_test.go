package adapter

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rubinthomasdev/go-employee/employee/adapter/out/db/postgres"
	"github.com/rubinthomasdev/go-employee/employee/adapter/out/persistence"
	"github.com/rubinthomasdev/go-employee/employee/domain"
)

func TestGetEmployeeDataFromPersistence(t *testing.T) {
	// create DB connection
	host := os.Getenv("TEST_PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("TEST_PG_PORT"))
	user := os.Getenv("TEST_PG_USER")
	password := os.Getenv("TEST_PG_PASSWORD")
	dbname := os.Getenv("TEST_PG_DBNAME")
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

	wantEmp := domain.Employee{
		EmployeeID: domain.EmployeeID{ID: 1},
		Name:       domain.EmployeeName{FirstName: "test1", LastName: "user"},
		BaseSalary: domain.Money{Amount: 12.12},
		Bonus:      domain.Money{Amount: 13.13},
	}
	gotEmp := getEmpAdapter.GetEmployeeDataFromPersistence(domain.EmployeeID{ID: 1})

	if !reflect.DeepEqual(wantEmp, gotEmp) {
		t.Errorf("Data mismatched from DB. Wanted : %v, Got : %v\n", wantEmp, gotEmp)
	}
}

func TestGetAllEmployees(t *testing.T) {
	// create DB connection
	host := os.Getenv("TEST_PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("TEST_PG_PORT"))
	user := os.Getenv("TEST_PG_USER")
	password := os.Getenv("TEST_PG_PASSWORD")
	dbname := os.Getenv("TEST_PG_DBNAME")
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

	wantEmp := []domain.Employee{
		{
			EmployeeID: domain.EmployeeID{ID: 1},
			Name:       domain.EmployeeName{FirstName: "test1", LastName: "user"},
			BaseSalary: domain.Money{Amount: 12.12},
			Bonus:      domain.Money{Amount: 13.13},
		},
		{
			EmployeeID: domain.EmployeeID{ID: 2},
			Name:       domain.EmployeeName{FirstName: "test2", LastName: "user"},
			BaseSalary: domain.Money{Amount: 12.56},
			Bonus:      domain.Money{Amount: 13.18},
		},
	}
	gotEmp := getEmpAdapter.GetAllEmployees()

	if !reflect.DeepEqual(wantEmp, gotEmp) {
		t.Errorf("Data mismatched from DB. Wanted : %v, Got : %v\n", wantEmp, gotEmp)
	}
}
