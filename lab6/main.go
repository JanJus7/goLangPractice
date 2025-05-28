package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"lab6/indicators"
	"os"
	"lab6/tools"
	"lab6/functions"
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

	prices, err := tools.ReadCSV(*csvPath)
	if err != nil {
		panic(err)
	}

	sma := indicators.SMA(prices, *smaP)
	rsi := indicators.RSI(prices, *rsiP)
	atr := indicators.ATR(prices, *atrP)

    start := functions.MaxI(*smaP, *rsiP, *atrP) - 1
	if len(prices) < start {
		panic("Not enough data to calculate indicators...")
	}

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
