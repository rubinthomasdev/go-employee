package domain

import "testing"

func TestMoneyAdd(t *testing.T) {
	testCases := []struct {
		a        Money
		b        Money
		expected Money
	}{
		{Money{10.5}, Money{11.5}, Money{22.0}},
		{Money{10.56}, Money{11.517}, Money{22.1}},
		{Money{10.12}, Money{11.34}, Money{21.5}},
		{Money{10.12}, Money{-11.34}, Money{-1.2}},
		{Money{10.12}, Money{-11.77}, Money{-1.7}},
		{Money{10.12}, Money{-11.76}, Money{-1.6}},
	}

	for _, tc := range testCases {
		want := tc.expected
		got := tc.a.Add(tc.b)

		if want != got {
			t.Errorf("Money add failed. Wanted %.1f, Got %.1f", want.Amount, got.Amount)
		}
	}
}
