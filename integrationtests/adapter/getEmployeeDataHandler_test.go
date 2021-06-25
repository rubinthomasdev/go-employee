package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rubinthomasdev/go-employee/employee/adapter/in/web"
	"github.com/rubinthomasdev/go-employee/employee/application/port/in"
	"github.com/rubinthomasdev/go-employee/employee/domain"
)

type MockGetEmployeeDataUseCase struct{}

func (m MockGetEmployeeDataUseCase) GetEmployeeDetails(id in.EmployeeQueryID) domain.Employee {
	return domain.Employee{
		EmployeeID: domain.EmployeeID{ID: 100},
		Name:       domain.EmployeeName{FirstName: "test", LastName: "user"},
		BaseSalary: domain.Money{Amount: 12.12},
		Bonus:      domain.Money{Amount: 14.14},
	}
}

func (m MockGetEmployeeDataUseCase) GetAllEmployees() []domain.Employee {
	return []domain.Employee{
		{
			EmployeeID: domain.EmployeeID{ID: 123},
			Name:       domain.EmployeeName{FirstName: "test", LastName: "user"},
			BaseSalary: domain.Money{Amount: 12.12},
			Bonus:      domain.Money{Amount: 14.14},
		},
	}
}

func TestGetEmployeeDetailsForID(t *testing.T) {
	getHandler := web.GetEmployeeDataHandler{
		GetEmployeeDataUseCase: MockGetEmployeeDataUseCase{},
		Mapper:                 web.GetEmployeeDomainDTOMapper{},
	}

	wantStatus := 200
	wantEmployeeID := 100

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/employees/1", nil)
	getHandler.GetEmployeeDetails(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var gotEmp web.GetEmployeeDTO
	fmt.Println(string(body))
	err := json.Unmarshal(body, &gotEmp)
	if err != nil {
		t.Errorf("error json parsing in web adapter call. %v", err)
	}

	if wantEmployeeID != gotEmp.EmployeeID {
		t.Errorf("web adaper call failed. Wanted %d, Got %d", wantEmployeeID, gotEmp.EmployeeID)
	}

	if wantStatus != w.Code {
		t.Errorf("web adaper call status failed. Wanted %d, Got %d", wantStatus, w.Code)
	}
}

func TestGetAllEmployeeDetails(t *testing.T) {
	getHandler := web.GetEmployeeDataHandler{
		GetEmployeeDataUseCase: MockGetEmployeeDataUseCase{},
		Mapper:                 web.GetEmployeeDomainDTOMapper{},
	}

	wantStatus := 200
	wantEmployeeID := 123

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/employees", nil)
	getHandler.GetEmployeeDetails(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var gotEmp []web.GetEmployeeDTO
	fmt.Println(string(body))
	err := json.Unmarshal(body, &gotEmp)
	if err != nil {
		t.Errorf("error json parsing in web adapter call. %v", err)
	}

	if wantEmployeeID != gotEmp[0].EmployeeID {
		t.Errorf("web adaper call failed. Wanted %d, Got %d", wantEmployeeID, gotEmp[0].EmployeeID)
	}

	if wantStatus != w.Code {
		t.Errorf("web adaper call status failed. Wanted %d, Got %d", wantStatus, w.Code)
	}
}

func TestGetAllReturnsIfMethodIsNotGet(t *testing.T) {
	getHandler := web.GetEmployeeDataHandler{
		GetEmployeeDataUseCase: MockGetEmployeeDataUseCase{},
		Mapper:                 web.GetEmployeeDomainDTOMapper{},
	}

	wantStatus := http.StatusMethodNotAllowed

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/employees", nil)
	getHandler.GetEmployeeDetails(w, r)

	if wantStatus != w.Code {
		t.Errorf("web adaper call status failed. Wanted %d, Got %d", wantStatus, w.Code)
	}
}
