package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCSVLineByLine(filepath string) {

	// otwarcie pliku
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Uwaga: separator to , a nie ;
	reader := csv.NewReader(file)

	// czytanie linia po linii
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func main() {
	readCSVLineByLine("weatherref-france-vigilance-meteo-departement.csv")
}