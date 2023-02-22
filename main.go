package main

import (
	"flag"
	"fmt"
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
			celsius = (fahr - 32) * 5 / 9
		} else if kelvin != 0 {
			celsius = kelvin - 273.15
		}
		fmt.Printf("%.2f%s er %.2f%s\n", fahr, "°F", celsius, "°C")
	case "F":
		if celsius != 0 {
			fahr = celsius*9/5 + 32
		} else if kelvin != 0 {
			celsius = kelvin - 273.15
			fahr = celsius*9/5 + 32
		}
		fmt.Printf("%.2f%s er %.2f%s\n", celsius, "°C", fahr, "°F")
	case "K":
		if celsius != 0 {
			kelvin = celsius + 273.15
		} else if fahr != 0 {
			celsius = (fahr - 32) * 5 / 9
			kelvin = celsius + 273.15
		}
		fmt.Printf("%.2f%s er %.2f%s\n", celsius, "°C", kelvin, "K")
	default:
		fmt.Println("Error: ugyldig temperaturskala valgt")
	}

}
