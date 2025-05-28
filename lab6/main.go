package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"lab6/indicators"
	"os"
	"strconv"
)

func main() {
	csvPath := flag.String("csv", "", "Path to csv")
	outPath := flag.String("out", "output.csv", "Output csv file")
	smaP := flag.Int("sma", 20, "SMA period")
	rsiP := flag.Int("rsi", 21, "RSI period")
	atrP := flag.Int("atr", 14, "ATR period")
	flag.Parse()

	if *csvPath == "" {
		panic("Please provide a valid path to CSV.")
	}

	prices, err := readCSV(*csvPath)
	if err != nil {
		panic(err)
	}

	sma := indicators.SMA(prices, *smaP)
	rsi := indicators.RSI(prices, *rsiP)
	atr := indicators.ATR(prices, *atrP)

    start := max(*smaP, *rsiP, *atrP) - 1

	f, err := os.Create(*outPath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    w := csv.NewWriter(f)
    defer w.Flush()

    header := []string{"Date", fmt.Sprintf("SMA(%d)", *smaP), fmt.Sprintf("RSI(%d)", *rsiP), fmt.Sprintf("ATR(%d)", *atrP)}
    if err := w.Write(header); err != nil {
        panic(err)
    }

    for i := start; i < len(prices); i++ {
        row := []string{
            prices[i].Date,
            fmt.Sprintf("%.2f", sma[i]),
            fmt.Sprintf("%.2f", rsi[i]),
            fmt.Sprintf("%.2f", atr[i]),
        }
        if err := w.Write(row); err != nil {
            panic(err)
        }
    }
}

func readCSV(path string) ([]indicators.Price, error) {
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

func max(a, b, c int) int {
    m := a
    if b > m {
        m = b
    }
    if c > m {
        m = c
    }
    return m
}