package main

import (
	"flag"
	"fmt"
	"lab7/tools"
)

func main() {
	city := flag.String("city", "Gdynia", "City to get weather for")
	daysForward := flag.Int("days", 7, "Number of days to forecast")
	histStartDate := flag.String("histStart", "2025-06-01", "Start date for historical data (YYYY-MM-DD)")
	histEndDate := flag.String("histEnd", "2025-06-07", "End date for historical data (YYYY-MM-DD)")
	typeOfForcast := flag.String("type", "future", "Type of weather to fetch (future, hourly, historical (by dates)")
	flag.Parse()


	lat, lng := tools.GetCityData(*city)

	if *typeOfForcast == "future" {
		date, tempMax, tempMin, windSpeed, rain, err := tools.GetFutureWeatherData(lat, lng, *daysForward)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Weather forecast for %s:\n", *city)
		for i := 0; i < len(date); i++ {
			fmt.Printf("Date: %s, Max Temp: %.2f°C, Min Temp: %.2f°C, Wind Speed: %.2f m/s, Rain: %.2f mm\n",
				date[i], tempMax[i], tempMin[i], windSpeed[i], rain[i])
		}
	} else if *typeOfForcast == "hourly" {
		date, temp, rain, windSpeed, err := tools.GetHourlyWeatherData(lat, lng)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hourly weather for %s:\n", *city)
		for i := 0; i < len(date); i++ {
			fmt.Printf("Date: %s, Temp: %.2f°C, Wind Speed: %.2f m/s, Rain: %.2f mm\n",
				date[i], temp[i], windSpeed[i], rain[i])
		}
	} else if *typeOfForcast == "historical" {
		date, tempMax, tempMin, windSpeed, rain, err := tools.GetHistoricalWeatherData(lat, lng, *histStartDate, *histEndDate)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Historical weather for %s from %s to %s:\n", *city, *histStartDate, *histEndDate)
		for i := 0; i < len(date); i++ {
			fmt.Printf("Date: %s, Max Temp: %.2f°C, Min Temp: %.2f°C, Wind Speed: %.2f m/s, Rain: %.2f mm\n",
				date[i], tempMax[i], tempMin[i], windSpeed[i], rain[i])
		}
	}
}