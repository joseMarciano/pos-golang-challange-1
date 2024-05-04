package presenters

type TemperatureRequestOutput struct {
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
	Kelvin     float32 `json:"kelvin"`
}
