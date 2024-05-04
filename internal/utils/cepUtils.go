package utils

import "regexp"

func IsValidCep(cep string) bool {
	cepRegex := regexp.MustCompile(`^\d{8}$`)
	return cepRegex.MatchString(cep)
}
