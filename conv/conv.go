package conv

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

// Konverterer Celsius til Farhenheit
func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

// Konverterer Kelvin til Farhenheit
func KelvinToFahrenheit(value float64) float64 {
	return value*9/5 - 459.67
}

// Konverterer Farhenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}
