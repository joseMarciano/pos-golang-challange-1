package find

import (
	"fmt"
	"pos-gollang-challenge-1/internal/gateways"
	"pos-gollang-challenge-1/internal/utils"
)

type FindTemperatureUseCase struct {
	temperatureGateway gateways.TemperatureGateway
	cityGateway        gateways.CityGateway
}

func NewFindTemperatureUseCase(gateway gateways.TemperatureGateway, cityGateway gateways.CityGateway) *FindTemperatureUseCase {
	return &FindTemperatureUseCase{gateway, cityGateway}
}

func (uc *FindTemperatureUseCase) Execute(input *FindTemperatureUseCaseInput) (*FindTemperatureUseCaseOutput, error) {
	if !utils.IsValidCep(input.ZipCode) {
		return nil, fmt.Errorf("invalid zipcode")
	}

	city, err := uc.cityGateway.FindByZipCode(input.ZipCode)
	if err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}

	temperature, err := uc.temperatureGateway.FindByCityCode(city.Id)
	if err != nil {
		return nil, fmt.Errorf("can not find temperature")
	}

	return &FindTemperatureUseCaseOutput{
		Celsius:    temperature.Celsius,
		Fahrenheit: temperature.Fahrenheit,
		Kelvin:     temperature.Kelvin,
	}, nil
}
