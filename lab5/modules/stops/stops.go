package stops

import (
	"net/http"
	"strings"
	"encoding/json"
)

type Stop struct {
	StopID   int    `json:"stopId"`
	StopName string `json:"stopName"`
}

var stops []Stop

func LoadStops() error {
	url := "https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json"
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var raw map[string]struct {
		LastUpdate string `json:"lastUpdate"`
		Stops      []Stop `json:"stops"`
	}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return err
	}

	for _, data := range raw {
		stops = data.Stops
		break
	}
	return nil
}

func FindStopsByName(name string) []Stop {
	var result []Stop

	name = strings.ToLower(name)
	for _, stop := range stops {
		if strings.Contains(strings.ToLower(stop.StopName), name) {
			result = append(result, stop)
		}
	}
	return result
}
