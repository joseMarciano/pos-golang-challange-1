package model

type Temperature struct {
	Celsius    float32
	Kelvin     float32
	Fahrenheit float32
}

func NewTemperature(celsiusTemp float32) *Temperature {
	return &Temperature{
		Celsius:    celsiusTemp,
		Kelvin:     fromCelsiusToKelvin(celsiusTemp),
		Fahrenheit: fromCelsiusToFahrenheit(celsiusTemp),
	}
}

func fromCelsiusToFahrenheit(celsiusTemp float32) float32 {
	return (celsiusTemp * 1.8) + 32
}

func fromCelsiusToKelvin(celsiusTemp float32) float32 {
	return celsiusTemp + 273.15
}
