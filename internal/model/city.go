package model

type City struct {
	Id      string
	ZipCode string
}

func NewCity(id, zipCode string) *City {
	return &City{
		Id:      id,
		ZipCode: zipCode,
	}
}
