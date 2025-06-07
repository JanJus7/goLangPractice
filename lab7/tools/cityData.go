package tools

import (
	"encoding/csv"
	"os"
)

func GetCityData(cityName string) (Lat string, Lng string){
	file, err := os.Open("pl.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
    reader.Comma = ','

	for{
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		if record[0] == cityName {
			Lat = record[1]
			Lng = record[2]
			return
		}
		
	}

	return Lat, Lng
}