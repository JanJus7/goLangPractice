package tools

import (
	"encoding/csv"
	"lab6/indicators"
	"os"
	"strconv"
)
func ReadCSV(path string) ([]indicators.Price, error) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	header := rows[0]

	x := map[string]int{}
	for i, col := range header {
		x[col] = i
	}

	var out []indicators.Price
	for _, row := range rows[1:] {
		open,_ := strconv.ParseFloat(row[x["Otwarcie"]], 64)
		close,_ := strconv.ParseFloat(row[x["Zamkniecie"]], 64)
		high,_ := strconv.ParseFloat(row[x["Najwyzszy"]], 64)
		low,_ := strconv.ParseFloat(row[x["Najnizszy"]], 64)
		volume,_ := strconv.ParseFloat(row[x["Wolumen"]], 64)

		out = append(out, indicators.Price{
            Date:   row[x["Data"]],
            Open:   open,
            High:   high,
            Low:    low,
            Close:  close,
            Volume: volume,
        })
	}

	return out, nil
}