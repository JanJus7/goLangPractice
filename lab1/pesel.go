package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GeneratePESEL: geneuje numer PESEL
// Parametry:
// - birthDate: time.Time: reprezentacja daty urodzenia
// - płeć: znak "M" lub "K"
// Wyjscie:
//Tablica z cyframi numeru PESEL

func GenerateCheckDigit(cyfryPESEL [10]int) int {
	weights := [10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	sum := 0

	for i := 0; i < 10; i++ {
		sum += cyfryPESEL[i] * weights[i]
	}

	controlDigit := (10 - (sum % 10)) % 10
	return controlDigit
}


func GenerujPESEL(birthDate time.Time, gender string) [11]int {

	var cyfryPESEL [11]int

	year := birthDate.Year()
	month := int(birthDate.Month())
	day := birthDate.Day()

	randomSerial := rand.Intn(900) + 100

	cutYear := year % 100

	a := cutYear / 10
	b := cutYear % 10
	var cutMonth = 0

	switch {
	case 1800 <= year && year < 1900:
		cutMonth = month + 80
	case 1900 <= year && year < 2000:
		cutMonth = month
	case 2000 <= year && year < 2100:
		cutMonth = month + 20
	case 2100 <= year && year < 2200:
		cutMonth = month + 40
	case 2200 <= year && year < 2300:
		cutMonth = month + 60
	}

	c := cutMonth / 10
	d := cutMonth % 10

	e := day / 10
	f := day % 10

	g := randomSerial / 100
	twoDigits := randomSerial % 100
	h := twoDigits / 10
	i := twoDigits % 10

	var reqNumber = 0
	switch {
	case gender == "K":
		reqNumber = rand.Intn(5) * 2
	case gender == "M":
		reqNumber = rand.Intn(5) * 2 + 1
	}


	fmt.Println(year, "=>", cutYear, "=>", a, b)
	fmt.Println(month, "=>", cutMonth, "=>", c, d)
	fmt.Println(day, "=>", e, f)
	fmt.Println(randomSerial, "=>", g, h, i)
	fmt.Println(gender, "=>", reqNumber)

	cyfryPESEL[0] = a
	cyfryPESEL[1] = b
	cyfryPESEL[2] = c
	cyfryPESEL[3] = d
	cyfryPESEL[4] = e
	cyfryPESEL[5] = f
	cyfryPESEL[6] = g
	cyfryPESEL[7] = h
	cyfryPESEL[8] = i
	cyfryPESEL[9] = reqNumber
	var cyfryPierwsze10 [10]int
	copy(cyfryPierwsze10[:], cyfryPESEL[:10])
	cyfryPESEL[10] = GenerateCheckDigit(cyfryPierwsze10)
	

	return cyfryPESEL
}

// WeryfikujPESEL: weryfikuje poprawność numeru PESEL
// Parametry:
// - cyfryPESEL: Tablica z cyframi numeru PESEL
// Wyjscie:
//zmienna bool

func WeryfikujPESEL(cyfryPESEL [11]int) bool {

	var czyPESEL bool

	switch {
	case len(cyfryPESEL) == 11:
		czyPESEL = true
	default:
		czyPESEL = false
	}

	return czyPESEL
}

func main() {
	
	birthDate := time.Date(1981, 2, 26, 0, 0, 0, 0, time.UTC)
	pesel := GenerujPESEL(birthDate, "M")

	fmt.Println("Wygenerowany PESEL:", pesel)

	fmt.Println("Czy numer PESEL jest poprawny:", WeryfikujPESEL(pesel))

}
