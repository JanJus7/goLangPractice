package busdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StopTime struct {
	StopID        int    `json:"stopId"`
	ArrivalTime   string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	Sequence      int    `json:"stopSequence"`
	TripID        int `json:"tripId"`
}

func GetTravelTimes(date string, routeId string) ([]StopTime, error) {
	url := fmt.Sprintf("https://ckan2.multimediagdansk.pl/stopTimes?date=%s&routeId=%s", date, routeId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var raw struct {
		LastUpdate string     `json:"lastUpdate"`
		StopTimes  []StopTime `json:"stopTimes"`
	}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	return raw.StopTimes, nil
}