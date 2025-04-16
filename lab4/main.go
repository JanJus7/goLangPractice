package main

import (
	"lab4/help"
	"sync"
)

const (
	numWorkers = 5
	numOrders  = 20
)

func main() {

	orders := make(chan help.Order, numOrders)
	results := make(chan help.ProcessResult, numOrders)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go help.Worker(i, orders, results, &wg)
	}

	go help.GenerateOrders(numOrders, orders)

	wg.Wait()
	close(results)

	help.ProcessResults(results)
}
