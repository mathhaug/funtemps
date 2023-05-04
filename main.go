package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	conv "github.com/mathhaug/funtemps/conv"
	fun "github.com/mathhaug/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64
var out string
var funfacts string
var t string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&out, "out", "C", "calculate temperature in C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, Luna - Månen og terra - Jorden")
	flag.StringVar(&t, "T", "C", "temperature unit for funfacts")
}

// Command example: "go run main.go -F 100 -out C"

func main() {

	flag.Parse()

	// Tar for seg konvertering av gradene
	if isFlagPassed("F") && isFlagPassed("out") {
		if out == "C" {
			celsius = conv.FahrenheitToCelsius(fahr)
			fmt.Printf("%g°F is equal to %s°C\n", fahr, NumFormat(celsius))
		} else if out == "K" {
			kelvin = conv.FahrenheitToKelvin(fahr)
			fmt.Printf("%g°F is equal to %s°K\n", fahr, NumFormat(kelvin))
		} else {
			fmt.Println("Invalid output unit\nCommand example: 'go run main.go -F/C/K 100 -out C/F/K'")
			os.Exit(1)
		}
	} else if isFlagPassed("C") && isFlagPassed("out") {
		if out == "F" {
			fahr = conv.CelsiusToFahrenheit(celsius)
			fmt.Printf("%g°C is equal to %s°F\n", celsius, NumFormat(fahr))
		} else if out == "K" {
			kelvin = conv.CelsiusToKelvin(celsius)
			fmt.Printf("%g°C is equal to %s°K\n", celsius, NumFormat(kelvin))
		} else {
			fmt.Println("Invalid output unit\nCommand example: 'go run main.go -F/C/K 100 -out C/F/K'")
			os.Exit(1)

		}
	} else if isFlagPassed("K") && isFlagPassed("out") {
		if out == "F" {
			fahr = conv.KelvinToFarhenheit(kelvin)
			fmt.Printf("%g°K is equal to %s°F\n", kelvin, NumFormat(fahr))
		} else if out == "C" {
			celsius = conv.KelvinToCelsius(kelvin)
			fmt.Printf("%g°K is equal to %s°C\n", kelvin, NumFormat(celsius))
		} else {
			fmt.Println("Invalid output unit\nCommand example: 'go run main.go -F/C/K 100 -out C/F/K'")
			os.Exit(1)
		}
	}

	// tar for seg hvilke funfacts og hvilke grader de skal komme i
	if isFlagPassed("funfacts") {
		// månen er ikke en planet, men var beste ordet jeg kom på
		planet := ""
		switch funfacts {
		case "sun":
			planet = "Sun"
		case "luna":
			planet = "Luna"
		case "terra":
			planet = "Terra"
		}
		if planet != "" {
			facts := fun.GetFunFacts(funfacts, 0, t)
			fmt.Printf("Fun facts about %s in %s°\n", planet, t)
			for _, fact := range facts {
				fmt.Println(fact)
			}
		}
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// Formaterer verdine
func NumFormat(num float64) string {

	// Runder nummeret til 2 desimaler
	roundedNum := fmt.Sprintf("%.2f", num)

	// Fjerner unødvendige nuller om det finnes
	trimmedNum := strings.TrimRight(roundedNum, "0")

	// Fjerner punktum visst det er nødvendig
	if trimmedNum[len(trimmedNum)-1] == '.' {
		trimmedNum = trimmedNum[:len(trimmedNum)-1]
	}
	return trimmedNum
}
