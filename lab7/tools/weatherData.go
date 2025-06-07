package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func GetFutureWeatherData(Lat string, Lng string, daysForward int) (date []string, tempMax []float64, tempMin []float64, windSpeed []float64, rain []float64, err error) {
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + Lat + "&longitude=" + Lng + "&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,wind_speed_10m_max&timezone=Europe%2FWarsaw&forecast_days=" + strconv.Itoa(daysForward)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, nil, nil, nil, err
	}
	
	var returnedData Weather

	err = json.NewDecoder(resp.Body).Decode(&returnedData)
    if err != nil {
        return nil, nil, nil, nil, nil, fmt.Errorf("błąd dekodowania JSON: %v", err)
    }

	return returnedData.Daily.Time, 
       returnedData.Daily.TemperatureMax, 
       returnedData.Daily.TemperatureMin, 
       returnedData.Daily.WindSpeedMax,
       returnedData.Daily.PrecipitationSum, 
       nil
}

func GetHourlyWeatherData(Lat string, Lng string) (date []string, temp []float64, windSpeed []float64, rain []float64, err error) {
	today := time.Now().Format("2006-01-02")
	
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + Lat + "&longitude=" + Lng + "&hourly=temperature_2m,precipitation,wind_speed_10m&timezone=Europe%2FWarsaw&start_date=" + today + "&end_date=" + today

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, nil, nil, err
	}
	
	var returnedData HourlyWeather

	err = json.NewDecoder(resp.Body).Decode(&returnedData)
    if err != nil {
        return nil, nil, nil, nil, fmt.Errorf("błąd dekodowania JSON: %v", err)
    }

	return returnedData.Hourly.Time,
	       returnedData.Hourly.Temperature,
		   returnedData.Hourly.Precipitation,
		   returnedData.Hourly.WindSpeed,
		   nil
}

func GetHistoricalWeatherData(Lat string, Lng string, startDate string, endDate string) (date []string, tempMax []float64, tempMin []float64, windSpeed []float64, rain []float64, err error) {
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + Lat + "&longitude=" + Lng + "&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,wind_speed_10m_max&timezone=Europe%2FWarsaw&start_date=" + startDate + "&end_date=" + endDate

		resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, nil, nil, nil, err
	}
	
	var returnedData Weather

	err = json.NewDecoder(resp.Body).Decode(&returnedData)
    if err != nil {
        return nil, nil, nil, nil, nil, fmt.Errorf("błąd dekodowania JSON: %v", err)
    }

	return returnedData.Daily.Time, 
       returnedData.Daily.TemperatureMax, 
       returnedData.Daily.TemperatureMin, 
       returnedData.Daily.WindSpeedMax,
       returnedData.Daily.PrecipitationSum, 
       nil
}

