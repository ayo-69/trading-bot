package types

import "time"

type Trade struct {
	Timestamp time.Time
	Symbol    string
	Side      string
	Price     float64
	Quantity  float64
}
