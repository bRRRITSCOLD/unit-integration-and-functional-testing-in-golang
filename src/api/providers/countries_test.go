package providers

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {
	// Init

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
	httpClient := getClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	responder := httpmock.NewStringResponder(0, "")
	httpmock.RegisterResponder("GET", getCountryUrl+"/AR", responder)

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid client error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	// Init
	httpClient := getClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	responder := httpmock.NewStringResponder(0, `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`)
	httpmock.RegisterResponder("GET", getCountryUrl+"/ARS", responder)

	// Test
	country, err := GetCountry("ARS")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	// Init
	httpClient := getClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	responder := httpmock.NewStringResponder(0, `{"message": "invalid client error when getting country AR", "error": "", "status": "404", "cause": []}`)
	httpmock.RegisterResponder("GET", getCountryUrl+"/AR", responder)

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid client error when getting country AR", err.Message)
}

func TestGetCountryInvalidJSONResponse(t *testing.T) {
	// Init
	httpClient := getClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())
	responder := httpmock.NewStringResponder(200, `{"id": 123, "name": "Argentina", "time_zone": "GMT-03:00"}`)
	httpmock.RegisterResponder("GET", getCountryUrl+"/AR", responder)

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}
