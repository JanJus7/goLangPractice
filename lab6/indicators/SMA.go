package indicators

func SMA(prices []Price, period int) []float64 {
	n := len(prices)
	sma := make([]float64, n)

	var sum float64

	for i := 0; i < n; i++ {
		sum += prices[i].Close
		if i >= period {
			sum -= prices[i-period].Close
		}

		if i >= period-1 {
			sma[i] = sum / float64(period)
		}
	}
	return sma
}