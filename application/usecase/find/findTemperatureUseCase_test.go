package find

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pos-gollang-challenge-1/internal/model"
	"testing"
)

type CityGatewayMock struct {
	mock.Mock
}

func (m *CityGatewayMock) FindByZipCode(zipCode string) (*model.City, error) {
	args := m.Called(zipCode)

	temp, ok := args.Get(0).(*model.City)
	if !ok {
		return nil, args.Error(1)
	}

	return temp, nil
}

type TemperatureMock struct {
	mock.Mock
}

func (m *TemperatureMock) FindByCityCode(cityCode string) (*model.Temperature, error) {
	args := m.Called(cityCode)

	city, ok := args.Get(0).(*model.Temperature)
	if !ok {
		return nil, args.Error(1)
	}

	return city, nil
}

func TestGivenInvalidZipCode_WhenFind_ThenReturnError(t *testing.T) {
	temperatureGateway := new(TemperatureMock)
	cityGateway := new(CityGatewayMock)

	useCase := NewFindTemperatureUseCase(temperatureGateway, cityGateway)

	_, err := useCase.Execute(&FindTemperatureUseCaseInput{ZipCode: "8886800"})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid zipcode")
}

func TestGivenValidZipCode_WhenCityGatewayThrows_ThenReturnError(t *testing.T) {
	zipCode := "88865000"
	temperatureGateway := new(TemperatureMock)
	cityGateway := new(CityGatewayMock)

	cityGateway.On("FindByZipCode", zipCode).Return(nil, fmt.Errorf("some error ocurred"))
	useCase := NewFindTemperatureUseCase(temperatureGateway, cityGateway)

	_, err := useCase.Execute(&FindTemperatureUseCaseInput{ZipCode: zipCode})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "can not find zipcode")
}

func TestGivenValidZipCode_WhenTemperatureGatewayThrows_ThenReturnError(t *testing.T) {
	expectedCityId := "sdasd5as465d1"
	expectedZipCode := "88865000"
	expectedCity := model.NewCity(expectedCityId, expectedZipCode)

	temperatureGateway := new(TemperatureMock)
	cityGateway := new(CityGatewayMock)

	cityGateway.On("FindByZipCode", expectedZipCode).Return(expectedCity, nil)
	temperatureGateway.On("FindByCityCode", expectedCityId).Return(nil, fmt.Errorf("some error occurred"))
	useCase := NewFindTemperatureUseCase(temperatureGateway, cityGateway)

	_, err := useCase.Execute(&FindTemperatureUseCaseInput{ZipCode: expectedZipCode})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "can not find temperature")
}

func TestGivenValidZipCode_WhenCallUseCase_ShouldReturnCorrectly(t *testing.T) {
	expectedCityId := "sdasd5as465d1"
	expectedZipCode := "88865000"
	expectedCity := model.NewCity(expectedCityId, expectedZipCode)

	var expectedCelsiusTemp float32 = 80.0
	var expectedFahrenheitTemp float32 = 176
	var expectedKelvinTemp float32 = 353.15

	expectedTemperature := model.NewTemperature(expectedCelsiusTemp)

	temperatureGateway := new(TemperatureMock)
	cityGateway := new(CityGatewayMock)

	cityGateway.On("FindByZipCode", expectedZipCode).Return(expectedCity, nil)
	temperatureGateway.On("FindByCityCode", expectedCityId).Return(expectedTemperature, nil)

	useCase := NewFindTemperatureUseCase(temperatureGateway, cityGateway)

	output, err := useCase.Execute(&FindTemperatureUseCaseInput{ZipCode: expectedZipCode})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, expectedCelsiusTemp, output.Celsius)
	assert.Equal(t, expectedFahrenheitTemp, output.Fahrenheit)
	assert.Equal(t, expectedKelvinTemp, output.Kelvin)
}
