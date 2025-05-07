package departures

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Departure struct {
	EstimatedTime string `json:"estimatedTime"`
	Headsign      string `json:"headsign"`
	RouteID       int    `json:"routeId"`
	VehicleCode   int    `json:"vehicleCode"`
	Delay         int    `json:"delayInSeconds"`
}

func GetDeparturesForStop(stopId string) ([]Departure, error) {
	url := fmt.Sprintf("https://ckan2.multimediagdansk.pl/departures?stopId=%s", stopId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw struct {
		LastUpdate string      `json:"lastUpdate"`
		Departures []Departure `json:"departures"`
	}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	return raw.Departures, nil
}
