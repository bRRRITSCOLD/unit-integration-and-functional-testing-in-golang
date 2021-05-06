package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unit-integration-and-functional-testing-in-golang/internal/api/domains"
	"unit-integration-and-functional-testing-in-golang/internal/api/utils/errors"
	"unit-integration-and-functional-testing-in-golang/internal/clients"
)

const (
	GetCountryUrl = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*domains.Country, *errors.APIError) {
	client := clients.GetHTTPClient()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf(GetCountryUrl, countryId))
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
			s := string(body)
			fmt.Println(s) // ABC€
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
