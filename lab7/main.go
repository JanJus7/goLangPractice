package main

import (
	"flag"
	"fmt"
	"lab7/tools"
)

func main() {
	city := flag.String("city", "Gdynia", "City to get weather for")
	daysForward := flag.Int("days", 7, "Number of days to forecast")
	// histStartDate := flag.String("histStart", "2025-06-01", "Start date for historical data (YYYY-MM-DD)")
	// histEndDate := flag.String("histEnd", "2025-06-07", "End date for historical data (YYYY-MM-DD)")
	flag.Parse()


	lat, lng := tools.GetCityData(*city)

	date, tempMax, tempMin, windSpeed, rain, err := tools.GetFutureWeatherData(lat, lng, *daysForward)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Weather forecast for %s:\n", *city)
	for i := 0; i < len(date); i++ {
		fmt.Printf("Date: %s, Max Temp: %.2f°C, Min Temp: %.2f°C, Wind Speed: %.2f m/s, Rain: %.2f mm\n",
			date[i], tempMax[i], tempMin[i], windSpeed[i], rain[i])
	}




}