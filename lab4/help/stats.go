package help

import "fmt"

func ProcessResults(results <-chan ProcessResult) {
	var total, success int

	for res := range results {
		total++
		if res.Success {
			success++
		} else {
			fmt.Printf("Order #%d was lost in the process: %v\n", res.OrderID, res.Error)
		}
	}

	fail := total - success
	successRate := float64(success) / float64(total) * 100

	fmt.Println("\n===Stats===")
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Fails: %d\n", fail)
	fmt.Printf("Success rate: %.2f%%\n", successRate)
}
