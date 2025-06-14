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
	if lat == "" || lng == "" {
    	panic(fmt.Errorf("could not find coordinates for city: %s", *city))
	}

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
			if tempMax[i] >= config.HighTemperatureThreshold {
				alerts = append(alerts, fmt.Sprintf("High temperature (%.2f°C)", tempMax[i]))
			}
			if tempMin[i] <= config.LowTemperatureThreshold {
				alerts = append(alerts, fmt.Sprintf("Low temperature (%.2f°C)", tempMin[i]))
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

		tools.GeneratePlot(date, tempMax, "Forecast - Temp Max", "Temp (°C)", "plots/futureTemp.png")
		tools.GeneratePlot(date, windSpeed, "Forecast - Wind", "Wind (m/s)", "plots/futureWind.png")
		tools.GeneratePlot(date, rain, "Forecast - Rainfall", "Rainfall (mm)", "plots/futureRain.png")

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

		tools.GeneratePlot(date, temp, "Hourly - Temperature", "Temp (°C)", "plots/hourlyTemp.png")
		tools.GeneratePlot(date, windSpeed, "Hourly - Wind", "Wind (m/s)", "plots/hourlyWind.png")
		tools.GeneratePlot(date, rain, "Hourly - Rainfall", "Rainfall (mm)", "plots/hourlyRain.png")

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

		tools.GeneratePlot(date, tempMax, "Historical - Temp Max", "Temp (°C)", "plots/historicalTemp.png")
		tools.GeneratePlot(date, windSpeed, "Historical - Wind", "Wind (m/s)", "plots/historicalWind.png")
		tools.GeneratePlot(date, rain, "Historical - Rainfall", "Rainfall (mm)", "plots/historicalRain.png")

	}
}