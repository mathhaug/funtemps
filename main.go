package main

import (
	"flag"
	"fmt"

	"github.com/mathhaug/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr, celsius, kelvin float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()

	// Sjekk at kun én temperaturskala er valgt
	numTemps := 0
	if fahr != 0 {
		numTemps++
	}
	if celsius != 0 {
		numTemps++
	}
	if kelvin != 0 {
		numTemps++
	}
	if numTemps != 1 {
		fmt.Println("Error: må spesifisere nøyaktig én temperaturskala")
		return
	}

	// Konverter temperatur til ønsket skala
	switch out {
	case "C":
		if fahr != 0 {
			celsius = conv.FahrenheitToCelsius(fahr)
		} else if kelvin != 0 {
			celsius = conv.KelvinToCelsius(kelvin)
		}
		fmt.Printf("%.2f%s er %.2f%s\n", fahr, "°F", celsius, "°C")
	case "F":
		if celsius != 0 {
			fahr = conv.CelsiusToFahrenheit(celsius)
		} else if kelvin != 0 {
			celsius = conv.KelvinToCelsius(kelvin)
			fahr = conv.CelsiusToFahrenheit(celsius)
		}
		fmt.Printf("%.2f%s er %.2f%s\n", celsius, "°C", fahr, "°F")
	case "K":
		if celsius != 0 {
			kelvin = conv.CelsiusToKelvin(celsius)
		} else if fahr != 0 {
			celsius = conv.FahrenheitToCelsius(fahr)
			kelvin = conv.CelsiusToKelvin(celsius)
		}
		fmt.Printf("%.2f%s er %.2f%s\n", celsius, "°C", kelvin, "K")
	default:
		fmt.Println("Error: ugyldig temperaturskala valgt")
	}

}
