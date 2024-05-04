package model

import (
	"testing"
)

func TestShouldCalculateTempsCorrectly(t *testing.T) {
	expectedCelsius := []float32{-30, -20, -10, 0, 10, 20, 30, 40, 50, 60}
	expectedFahrenheit := []float32{-22, -4, 14, 32, 50, 68, 86, 104, 122, 140}
	expectedKelvin := []float32{243.15, 253.15, 263.15, 273.15, 283.15, 293.15, 303.15, 313.15, 323.15, 333.15}

	for i, celsius := range expectedCelsius {
		fahrenheit := expectedFahrenheit[i]
		kelvin := expectedKelvin[i]

		temp := NewTemperature(celsius)

		if celsius != temp.Celsius {
			t.Errorf("Expected Celsius %f, got %f", celsius, temp.Celsius)
		}

		if fahrenheit != temp.Fahrenheit {
			t.Errorf("Expected Fahrenheit %f, got %f", fahrenheit, temp.Fahrenheit)
		}

		if kelvin != temp.Kelvin {
			t.Errorf("Expected Kelvin %f, got %f", kelvin, temp.Kelvin)
		}
	}

}
