package services

import (
	"unit-integration-and-functional-testing-in-golang/internal/api/domains"
	"unit-integration-and-functional-testing-in-golang/internal/api/providers"
	"unit-integration-and-functional-testing-in-golang/internal/api/utils/errors"
)

type countriesService struct{}

type countriesServiceInterface interface {
	GetCountry(countryId string) (*domains.Country, *errors.APIError)
}

var (
	CountriesService countriesServiceInterface
)

func init() {
	CountriesService = &countriesService{}
}

func (s *countriesService) GetCountry(countryId string) (*domains.Country, *errors.APIError) {
	country, err := providers.GetCountry(countryId)
	if err != nil {
		return nil, err
	}
	return country, nil
}
