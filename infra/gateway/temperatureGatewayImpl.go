package gateway

import (
	"log"
	"math/rand"
	"pos-gollang-challenge-1/internal/model"
	"time"
)

type TemperatureGatewayImpl struct{}

func NewTemperatureGatewayImpl() TemperatureGatewayImpl {
	return TemperatureGatewayImpl{}
}

func (t TemperatureGatewayImpl) FindByCityCode(_ string) (*model.Temperature, error) {
	// Seed the random number generator with the current time
	rand.NewSource(time.Now().UnixNano())

	// Generate a random float32 value between -10 and 50
	randomValue := rand.Float32()*(50+10) - 10

	log.Println("Random float32 value between -10 and 50:", randomValue)

	return model.NewTemperature(randomValue), nil
}
