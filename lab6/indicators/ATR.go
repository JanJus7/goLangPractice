package indicators

import "lab6/functions"

func ATR(prices []Price, period int) []float64 {
	n := len(prices)
	atr := make([]float64, n)

	var trSum float64

	for i := 1; i < n; i++ {
		highLow := prices[i].High - prices[i].Low
		highClose := functions.Abs(prices[i].High - prices[i-1].Close)
		lowClose := functions.Abs(prices[i].Low - prices[i-1].Close)
		
		tr := functions.MaxF(highLow, functions.MaxF(highClose, lowClose))

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
