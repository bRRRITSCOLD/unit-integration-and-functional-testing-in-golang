package providers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"unit-integration-and-functional-testing-in-golang/internal/clients"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	httpClient := clients.GetHTTPClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	os.Exit(m.Run())
}

func TestGetCountry(t *testing.T) {
	// Init
	countryId := "AR"

	httpmock.Reset()
	responder := httpmock.NewStringResponder(200, `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`)
	httpmock.RegisterResponder("GET", fmt.Sprintf(GetCountryUrl, countryId), responder)

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, country.Id, "AR")
	assert.EqualValues(t, country.Name, "Argentina")
	assert.EqualValues(t, country.TimeZone, "GMT-03:00")
	assert.EqualValues(t, 24, len(country.States))
}

func TestGetCountryClientError(t *testing.T) {
	// Init
	countryId := "AR"

	httpmock.Reset()
	responder := httpmock.NewErrorResponder(errors.New("Client Error"))
	httpmock.RegisterResponder("GET", fmt.Sprintf(GetCountryUrl, countryId), responder)

	// Test
	country, err := GetCountry(countryId)

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid client error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	// Init
	countryId := "ARS"

	httpmock.Reset()
	responder := httpmock.NewStringResponder(404, `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`)
	httpmock.RegisterResponder("GET", fmt.Sprintf(GetCountryUrl, countryId), responder)

	// Test
	country, err := GetCountry(countryId)

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorResponse(t *testing.T) {
	// Init
	countryId := "AR"

	httpmock.Reset()
	responder, _ := httpmock.NewJsonResponder(500, 1)
	httpmock.RegisterResponder("GET", fmt.Sprintf(GetCountryUrl, countryId), responder)

	// Test
	country, err := GetCountry(countryId)

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error response when getting country AR", err.Message)
}

func TestGetCountryInvalidJSONResponse(t *testing.T) {
	// Init
	countryId := "AR"

	httpmock.Reset()
	responder := httpmock.NewStringResponder(200, `{"id": 123, "name": "Argentina", "time_zone": "GMT-03:00"}`)
	httpmock.RegisterResponder("GET", fmt.Sprintf(GetCountryUrl, countryId), responder)

	// Test
	country, err := GetCountry(countryId)

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}
