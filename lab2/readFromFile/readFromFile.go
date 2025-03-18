package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func readCSVLineByLine(filepath string) []Weather {

	// otwarcie pliku
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Uwaga: separator to , a nie ;
	reader := csv.NewReader(file)
	reader.Comma = ';'

	var csvData []Weather

	isHeader := true

	// czytanie linia po linii
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Błąd podczas czytania rekordu: ", err)
		}

		// fmt.Println(record)

		if isHeader {
			isHeader = false
			continue
		}

		if len(record) == 9 {
			row := Weather{
				Id:           record[0],
				Echeance:     record[1],
				PhenomenId:   record[2],
				PhenomenDesc: record[3],
				DangerCode:   record[4],
				DangerColor:  record[5],
				StartDate:    record[6],
				EndDate:      record[7],
				EmissionDate: record[8],
			}

			csvData = append(csvData, row)
		} else {
			fmt.Println("Invalid record: ", record)
		}

	}

	// for _, row := range csvData {
	// 	fmt.Printf("Row: %+v\n", row)
	// }

	return csvData
}

func writeCSV(filepath string, data []Weather) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Błąd podczas tworzenia pliku:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	header := []string{"Id", "Echeance", "PhenomenId", "PhenomenDesc", "DangerCode", "DangerColor", "StartDate", "EndDate", "EmissionDate"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Błąd podczas zapisywania nagłówka:", err)
	}

	for _, row := range data {
		record := []string{
			row.Id,
			row.Echeance,
			row.PhenomenId,
			row.PhenomenDesc,
			row.DangerCode,
			row.DangerColor,
			row.StartDate,
			row.EndDate,
			row.EmissionDate,
		}
		if err := writer.Write(record); err != nil {
			log.Fatal("Błąd podczas zapisywania rekordu:", err)
		}
	}
}

func showStatistics(data []Weather) {
	dangerCodeCount := make(map[string]int)
	for _, row := range data {
		dangerCodeCount[row.DangerCode]++
	}

	fmt.Println("Statystyka DangerCode:")
	for code, count := range dangerCodeCount {
		fmt.Printf("DangerCode %s: %d wystąpień\n", code, count)
	}
}

type Weather struct {
	Id           string
	Echeance     string
	PhenomenId   string
	PhenomenDesc string
	DangerCode   string
	DangerColor  string
	StartDate    string
	EndDate      string
	EmissionDate string
}

func main() {
	csvData := readCSVLineByLine("weatherref-france-vigilance-meteo-departement.csv")
	sortedCsvData := make([]Weather, len(csvData))
	sortedCsvData2 := make([]Weather, len(csvData))

	copy(sortedCsvData, csvData)
	copy(sortedCsvData2, csvData)

	sort.Slice(sortedCsvData, func(i, j int) bool {
		return sortedCsvData[i].Id < sortedCsvData[j].Id
	})

	sort.Slice(sortedCsvData2, func(i, j int) bool {
		if sortedCsvData2[i].DangerCode == sortedCsvData2[j].DangerCode {
			return sortedCsvData2[i].Id < sortedCsvData2[j].Id
		}
		return sortedCsvData2[i].DangerCode < sortedCsvData2[j].DangerCode
	})

	writeCSV("data.csv", csvData)
	writeCSV("data-sorted.csv", sortedCsvData)
	writeCSV("data-sorted2.csv", sortedCsvData2)

	showStatistics(csvData)

	// for _, row := range csvData {
	//     fmt.Printf("Row: %+v\n", row)
	// }
}
