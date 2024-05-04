package gateways

import "pos-gollang-challenge-1/internal/model"

type TemperatureGateway interface {
	FindByCityCode(cityCode string) (*model.Temperature, error)
}

type CityGateway interface {
	FindByZipCode(zipCode string) (*model.City, error)
}
