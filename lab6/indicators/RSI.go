package indicators

func RSI(prices []Price, period int) []float64 {
	n := len(prices)
	rsi := make([]float64, n)

	var lossSum, gainSum float64

	for i := 1; i < n; i++ {
		change := prices[i].Close - prices[i-1].Close
		if change > 0 {
			gainSum += change
		} else {
			lossSum -= change
		}

		if i == period {
			avgGain := gainSum / float64(period)
			avgLoss := lossSum / float64(period)

			rs := avgGain / avgLoss

			rsi[i] = 100 - (100 / ( 1 + rs))

		}
	}

	for i := period; i < n; i++ {
		    change := prices[i].Close - prices[i-1].Close
			var gain, loss float64
			if change > 0 {
				gain = change
			} else {
				loss = -change
			}

			gainSum = (gainSum*float64(period-1) + gain) / float64(period)
			lossSum = (lossSum*float64(period-1) + loss) / float64(period)

			rs := gainSum / lossSum
			rsi[i] = 100 - (100 / (1 + rs))
	}

	return rsi
}