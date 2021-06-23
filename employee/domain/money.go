package domain

import "math"

type Money struct {
	Amount float64 `json:"amount"`
}

func (m Money) Add(n Money) Money {
	return Money{math.Round((m.Amount+n.Amount)*10) / 10}
}
