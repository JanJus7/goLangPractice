package indicators

func ATR(prices []Price, period int) []float64 {
	n := len(prices)
	atr := make([]float64, n)

	var trSum float64

	for i := 1; i < n; i++ {
		highLow := prices[i].High - prices[i].Low
		highClose := abs(prices[i].High - prices[i-1].Close)
		lowClose := abs(prices[i].Low - prices[i-1].Close)
		
		tr := max(highLow, max(highClose, lowClose))

		if i <= period {
            trSum += tr
        	if i == period {
                atr[i] = trSum / float64(period)
            }
        } else {
            atr[i] = (atr[i-1]*float64(period-1) + tr) / float64(period)
        }
    }
    return atr
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}