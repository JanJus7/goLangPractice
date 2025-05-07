package main

import (
	"bufio"
	"fmt"
	"lab5/modules/busdata"
	"lab5/modules/departures"
	"lab5/modules/stops"
	"os"
	"strings"
	"time"
)

func parseTime(value string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05", value)
	if err != nil {
		return time.Time{}
	}
	return t
}

func FormatRouteInfo(routeID, date string) string {
	stopTimes, err := busdata.GetTravelTimes(date, routeID)
	if err != nil {
		panic(err)
	}

	if len(stopTimes) == 0 {
		fmt.Println("Brak danych.")
		return ""
	}

	firstTripId := stopTimes[0].TripID
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Linia %s: przystanki i czasy przejazdu:\n", routeID))

	var previous *busdata.StopTime
	for _, stopTime := range stopTimes {
		if stopTime.TripID != firstTripId {
			break
		}

		tArrival := parseTime(stopTime.ArrivalTime)
		tDeparture := parseTime(stopTime.DepartureTime)

		sb.WriteString(fmt.Sprintf("  Przystanek ID: %d): Przyjazd %02d:%02d, Odjazd %02d:%02d\n",
			stopTime.StopID, tArrival.Hour(), tArrival.Minute(), tDeparture.Hour(), tDeparture.Minute()))

		if previous != nil {
			prevDep, _ := time.Parse("2006-01-02T15:04:05", previous.DepartureTime)
			diff := tArrival.Sub(prevDep).Minutes()
			sb.WriteString(fmt.Sprintf("  -> Czas przejazdu z przystanku %d: %.0f minut\n", previous.StopID, diff))
		}

		previous = &stopTime
	}

	return sb.String()
}

func main() {
	err := stops.LoadStops()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Menu:")
	fmt.Println("1 -> Wyszukaj przystanek")
	fmt.Println("2 -> Rozkład jazdy linii na dzisiejszy dzień")
	fmt.Println("3 -> Porównanie dwóch linii (równolegle)")

	mInput, _ := reader.ReadString('\n')
	mInput = strings.Replace(mInput, "\n", "", -1)
	switch mInput {
	case "1":
		fmt.Println("Podaj nazwę przystanku:")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		found := stops.FindStopsByName(input)
		if len(found) == 0 {
			fmt.Println("Nie znaleziono przystanku o podanej nazwie.")
			return
		}

		fmt.Println("Znalezione przystanki:")
		for _, stop := range found {
			fmt.Printf("ID: %d, Nazwa: %s\n", stop.StopID, stop.StopName)
		}
		fmt.Println("Podaj ID przystanku:")
		var stopID int
		fmt.Scanf("%d", &stopID)
		depList, err := departures.GetDeparturesForStop(fmt.Sprintf("%d", stopID))
		if err != nil {
			panic(err)
		}
		fmt.Println("Odjazdy z przystanku:")
		for _, dep := range depList {
			t, err := time.Parse(time.RFC3339, dep.EstimatedTime)
			if err != nil {
				fmt.Println("Błąd parsowania czasu:", err)
				continue
			}
			fmt.Printf("Linia: %d, Kierunek: %s, Czas odjazdu: %02d:%02d\n", dep.RouteID, dep.Headsign, t.Hour(), t.Minute())
		}
	case "2":
		fmt.Println("Podaj ID linii:")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		today := time.Now().Format("2006-01-02")
		fmt.Println(FormatRouteInfo(input, today))		

	case "3":
		fmt.Println("Podaj ID pierwszej linii:")
		line1, _ := reader.ReadString('\n')
		line1 = strings.Replace(line1, "\n", "", -1)

		fmt.Println("Podaj ID drugiej linii:")
		line2, _ := reader.ReadString('\n')
		line2 = strings.Replace(line2, "\n", "", -1)

		today := time.Now().Format("2006-01-02")
		results := make(chan string)

		go func() {
			results <- FormatRouteInfo(line1, today)
		}()
		go func() {
			results <- FormatRouteInfo(line2, today)
		}()

		fmt.Println(<-results)
		fmt.Println(<-results)

	default:
		fmt.Println("Nieznana opcja.")
		return

	}
}
