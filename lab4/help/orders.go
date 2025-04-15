package help

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sampleItems = []string{"Laptop", "Telefon", "Tablet", "Monitor", "Myszka", "Klawiatura"}

func GenerateOrders(numOrders int, orders chan<- Order) {
	for i := 1; i <= numOrders; i++ {
		order := Order{
			ID: i,
			CustomerName: fmt.Sprintf("Client %d", i),
			Items: randomItems(),
			TotalAmount: float64(rand.Intn(1000) + 100) + rand.Float64(),
		}
		orders <- order
		fmt.Printf("Order #%d: %s, Price: %.2f, Items: %v\n", order.ID, order.CustomerName, order.TotalAmount, order.Items)

		time.Sleep(time.Duration(rand.Intn(800)+800) * time.Millisecond)
	}
	close(orders)
}

func randomItems() []string {
	count := rand.Intn(3) + 1
	items := make([]string, count)
	for i := 0; i < count; i++ {
		items[i] = sampleItems[rand.Intn(len(sampleItems))]
	}
	return items
}

func Worker(id int, orders <-chan Order, results chan<- ProcessResult, wg *sync.WaitGroup) {
	defer wg.Done()

	const maxRetries = 10

	for order := range orders {
		start := time.Now()
		attempts := 0
		success := false

		for attempts < maxRetries {
			attempts++
			processTime := time.Duration(rand.Intn(1000)+500) * time.Millisecond
			time.Sleep(processTime)

			success = rand.Float32() < 0.2
			if success {
				break
			}

			fmt.Printf("Worker %d has processed the order #%d (success: %v)\n", id, order.ID, success)
		}

		result := ProcessResult{
			OrderID:      order.ID,
			CustomerName: order.CustomerName,
			Success:      success,
			ProcessTime:  time.Since(start),
			Attempts:     attempts,
		}
		if !success {
			result.Error = fmt.Errorf("error completing the order #%d after %d attempts", order.ID, attempts)
		}

		fmt.Printf("Worker %d has processed the order #%d (success: %v)\n", id, order.ID, success)
		results <- result
	}
}
