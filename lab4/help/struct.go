package help

import "time"

type Order struct {
	ID           int
	CustomerName string
	Items        []string
	TotalAmount  float64
}

type ProcessResult struct {
	OrderID      int
	CustomerName string
	Success      bool
	ProcessTime  time.Duration
	Error        error
	Attempts     int
}
