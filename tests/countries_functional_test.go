package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unit-integration-and-functional-testing-in-golang/internal/api/domains"
	"unit-integration-and-functional-testing-in-golang/internal/api/providers"
	"unit-integration-and-functional-testing-in-golang/internal/api/utils/errors"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {
	// Init
	countryId := "AR"

	httpmock.Reset()
	responder := httpmock.NewStringResponder(200, `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`)
	httpmock.RegisterResponder("GET", fmt.Sprintf(providers.GetCountryUrl, countryId), responder)

	// Execution
	response, err := http.Get("http://localhost:8080/countries/AR")

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, response)

	var country domains.Country
	bytes, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bytes, &country)

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, country.Id, "AR")
	assert.EqualValues(t, country.Name, "Argentina")
	assert.EqualValues(t, country.TimeZone, "GMT-03:00")
	assert.EqualValues(t, 24, len(country.States))
}

func TestGetCountryNotFound(t *testing.T) {
	// Init
	countryId := "ARS"

	httpmock.Reset()
	responder := httpmock.NewStringResponder(http.StatusNotFound, `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`)
	httpmock.RegisterResponder("GET", fmt.Sprintf(providers.GetCountryUrl, countryId), responder)

	// Execution
	response, err := http.Get("http://localhost:8080/countries/ARS")

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, response)

	var apiError errors.APIError
	bytes, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bytes, &apiError)

	assert.Nil(t, err)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, "not_found", apiError.Error)
	assert.EqualValues(t, "Country not found", apiError.Message)
}
