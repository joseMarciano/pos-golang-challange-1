package main

import (
	"github.com/gofiber/fiber/v2"
	"pos-gollang-challenge-1/application/usecase/find"
	"pos-gollang-challenge-1/infra/controller"
	"pos-gollang-challenge-1/infra/gateway"
)

func main() {
	temperatureGateway := gateway.NewTemperatureGatewayImpl()
	cityGateway := gateway.NewCityGatewayImpl()
	useCase := find.NewFindTemperatureUseCase(temperatureGateway, cityGateway)
	tempController := controller.NewTemperatureController(useCase)

	app := fiber.New()

	app.Get("/:zip", tempController.FindTemperature)
	app.Listen(":3000")
}
