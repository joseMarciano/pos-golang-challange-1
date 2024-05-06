package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"pos-gollang-challenge-1/application/usecase/find"
	"pos-gollang-challenge-1/infra/presenters"
)

type TemperatureController struct {
	findTemperatureUseCase *find.FindTemperatureUseCase
}

func NewTemperatureController(findTemperatureUseCase *find.FindTemperatureUseCase) *TemperatureController {
	return &TemperatureController{
		findTemperatureUseCase: findTemperatureUseCase,
	}
}

func (controller *TemperatureController) FindTemperature(c *fiber.Ctx) error {
	var outputBody *presenters.TemperatureRequestOutput
	zipCode := c.Params("zip")

	temperature, err := controller.findTemperatureUseCase.Execute(&find.FindTemperatureUseCaseInput{ZipCode: zipCode})
	if err != nil && err.Error() == "can not find zipcode" {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err != nil && err.Error() == "invalid zipcode" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	outputBody = &presenters.TemperatureRequestOutput{
		Celsius:    temperature.Celsius,
		Fahrenheit: temperature.Fahrenheit,
		Kelvin:     temperature.Kelvin,
	}

	err = json.NewEncoder(c).Encode(outputBody)
	if err != nil {
		return err
	}

	return nil
}
