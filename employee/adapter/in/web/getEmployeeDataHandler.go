package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/rubinthomasdev/go-employee/employee/application/port/in"
)

type GetEmployeeDataHandler struct {
	GetEmployeeDataUseCase in.GetEmployeeDetailsQuery
	Mapper                 GetEmployeeDomainDTOMapper
}

func (g GetEmployeeDataHandler) GetEmployeeDetails(w http.ResponseWriter, r *http.Request) {

	// return if method is not GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// create an encoder and set the response type as json
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	// if no id passed return all employees
	if r.URL.Path == "/api/v1/employees/" || r.URL.Path == "/api/v1/employees" {
		log.Println("getting all employees")
		empDTOSlice := []GetEmployeeDTO{}
		for _, empDomain := range g.GetEmployeeDataUseCase.GetAllEmployees() {
			empDTOSlice = append(empDTOSlice, g.Mapper.MapDomainToDTO(empDomain))
		}
		err := enc.Encode(empDTOSlice)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	// return employee for the given id
	empID, err := strconv.Atoi(r.URL.Path[len("/api/v1/employees/"):])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("getting employee data for employee id : %d \n", empID)
	empModel := g.GetEmployeeDataUseCase.GetEmployeeDetails(in.EmployeeQueryID{ID: empID})

	err = enc.Encode(g.Mapper.MapDomainToDTO(empModel))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
