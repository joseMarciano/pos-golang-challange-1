package find

type FindTemperatureUseCaseInput struct {
	ZipCode string
}

type FindTemperatureUseCaseOutput struct {
	Celsius    float32
	Fahrenheit float32
	Kelvin     float32
}
