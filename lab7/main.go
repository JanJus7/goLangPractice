package main

import (
	"flag"
	"fmt"
	"lab7/tools"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	city := flag.String("city", "Gdynia", "City to get weather for")
	daysForward := flag.Int("days", 7, "Number of days to forecast")
	histStartDate := flag.String("histStart", "2025-06-01", "Start date for historical data (YYYY-MM-DD)")
	histEndDate := flag.String("histEnd", "2025-06-07", "End date for historical data (YYYY-MM-DD)")
	typeOfForcast := flag.String("type", "future", "Type of weather to fetch (future, hourly, historical (by dates)")
	flag.Parse()


	lat, lng := tools.GetCityData(*city)

	config, err := tools.LoadConfig("config.json")
	if err != nil {
		panic(fmt.Errorf("error with config.json: %v", err))
	}

	if *typeOfForcast == "future" {
		date, tempMax, tempMin, windSpeed, rain, err := tools.GetFutureWeatherData(lat, lng, *daysForward)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Weather forecast for %s:\n", *city)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Data", "Max Temp (°C)", "Min Temp (°C)", "Wind (m/s)", "Rain (mm)"})

		for i := 0; i < len(date); i++ {
			row := []string{
				date[i],
				fmt.Sprintf("%.2f", tempMax[i]),
				fmt.Sprintf("%.2f", tempMin[i]),
				fmt.Sprintf("%.2f", windSpeed[i]),
				fmt.Sprintf("%.2f", rain[i]),
			}
			table.Append(row)

			alerts := []string{}
			if tempMax[i] >= config.TemperatureThreshold {
				alerts = append(alerts, fmt.Sprintf("High temperature (%.2f°C)", tempMax[i]))
			}
			if windSpeed[i] >= config.WindThreshold {
				alerts = append(alerts, fmt.Sprintf("Strong wind (%.2f m/s)", windSpeed[i]))
			}
			if rain[i] >= config.RainThreshold {
				alerts = append(alerts, fmt.Sprintf("Intense rainfall (%.2f mm)", rain[i]))
			}

			if len(alerts) > 0 {
				fmt.Printf("Warning [%s]: %s\n", date[i], alerts)
			}
		}
		table.Render()
	} else if *typeOfForcast == "hourly" {
		date, temp, rain, windSpeed, err := tools.GetHourlyWeatherData(lat, lng)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hourly weather for %s:\n", *city)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Hour", "Temp (°C)", "Wind (m/s)", "Rain (mm)"})

		for i := 0; i < len(date); i++ {
			row := []string{
				date[i],
				fmt.Sprintf("%.2f", temp[i]),
				fmt.Sprintf("%.2f", windSpeed[i]),
				fmt.Sprintf("%.2f", rain[i]),
			}
			table.Append(row)
		}
		table.Render()
	} else if *typeOfForcast == "historical" {
		date, tempMax, tempMin, windSpeed, rain, err := tools.GetHistoricalWeatherData(lat, lng, *histStartDate, *histEndDate)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Historical weather for %s from %s to %s:\n", *city, *histStartDate, *histEndDate)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Data", "Max Temp (°C)", "Min Temp (°C)", "Wind (m/s)", "Rain (mm)"})

		for i := 0; i < len(date); i++ {
			row := []string{
				date[i],
				fmt.Sprintf("%.2f", tempMax[i]),
				fmt.Sprintf("%.2f", tempMin[i]),
				fmt.Sprintf("%.2f", windSpeed[i]),
				fmt.Sprintf("%.2f", rain[i]),
			}
			table.Append(row)
		}
		table.Render()
	}
}