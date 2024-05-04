package gateway

import (
	"encoding/json"
	"net/http"
	"pos-gollang-challenge-1/internal/model"
	"strings"
)

const baseUrl = "https://viacep.com.br/ws/{zipCode}/json/"

type CityGatewayImpl struct{}

type ResponseApi struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewCityGatewayImpl() *CityGatewayImpl {
	return &CityGatewayImpl{}
}

func (c *CityGatewayImpl) FindByZipCode(zipCode string) (*model.City, error) {
	var apiResp *ResponseApi
	replace := strings.Replace(baseUrl, "{zipCode}", zipCode, -1)
	resp, err := http.Get(replace)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	return model.NewCity(apiResp.Siafi, apiResp.Cep), nil
}
