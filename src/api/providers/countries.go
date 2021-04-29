package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unit-integration-and-functional-testing-in-golang/src/api/domains"
	"unit-integration-and-functional-testing-in-golang/src/api/utils/errors"

	"github.com/go-resty/resty/v2"
)

const (
	getCountryUrl = "https://api.mercadolibre.com/countries/%s"
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
		Get(fmt.Sprintf(getCountryUrl, countryId))
	if err != nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid client error when getting country %s", countryId),
		}
	}

	body := resp.Body()

	if resp.StatusCode() > 299 {
		var apiErr errors.APIError
		if body != nil {
			if err := json.Unmarshal(body, &apiErr); err != nil {
				return nil, &errors.APIError{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
				}
			}
			return nil, &apiErr
		}
	}

	var result domains.Country
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &result, nil
}
