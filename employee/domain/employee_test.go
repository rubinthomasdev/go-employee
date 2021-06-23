package domain

import "testing"

func TestGetTotalSalary(t *testing.T) {
	testCases := []struct {
		employee Employee
		expected Money
	}{
		{
			Employee{EmployeeID: EmployeeID{1}, Name: EmployeeName{"john", "doe"}, BaseSalary: Money{10.0}, Bonus: Money{1.2}},
			Money{11.2},
		},
	}

	for _, tc := range testCases {
		want := tc.expected
		got := tc.employee.CalculateTotalSalary()

		if want != got {
			t.Errorf("Employee get total Salary failed. Wanted %.1f, Got %.1f", want.Amount, got.Amount)
		}
	}
}
