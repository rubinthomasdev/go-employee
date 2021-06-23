package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	EmployeeID int     `json:"employeeID"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	BaseSalary float64 `json:"baseSalary"`
	Bonus      float32 `json:"bonus"`
}

type InMemData struct {
	Employees map[int]Employee
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	empID, err := strconv.Atoi(r.URL.Path[len("/employee/"):])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("getting employee data for employee id : %d \n", empID)
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err = enc.Encode(inMemDB.Employees[empID])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

var inMemDB InMemData

func main() {
	// create employee data
	employees := make(map[int]Employee)
	employees[1] = Employee{EmployeeID: 1, FirstName: "jane", LastName: "doe", BaseSalary: 100.0, Bonus: 10.0}
	employees[2] = Employee{EmployeeID: 2, FirstName: "john", LastName: "doe", BaseSalary: 100.0, Bonus: 10.0}
	inMemDB = InMemData{Employees: employees}

	//http handler registration
	http.HandleFunc("/employee/", getEmployeeHandler)

	//start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
