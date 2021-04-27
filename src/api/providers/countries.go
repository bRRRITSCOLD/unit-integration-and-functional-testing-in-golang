package providers

import (
	"encoding/json"
	"net/http"
	"unit-integration-and-functional-testing-in-golang/src/api/domains"
	"unit-integration-and-functional-testing-in-golang/src/api/utils/errors"

	"github.com/go-resty/resty/v2"
)

var restyClient *resty.Client

func getClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}

func GetCountry(countryId string) (*domains.Country, *errors.APIError) {
	client := getClient()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&domains.Country{}). // or SetResult(AuthSuccess{}).
		SetError(&errors.APIError{}).  // or SetError(AuthError{}).
		Get("https://api.mercadolibre.com/countries/AR")
	if err != nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: "error",
		}
	}

	if resp.StatusCode() > 299 {
		var apiErr errors.APIError = errors.APIError{}
		body := resp.Body()
		if body != nil {
			apiErr, err = json.Unmarshal(&apiErr, body)
		}
	}
	return resp, nil
}
