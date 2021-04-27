package providers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {

}

func TestGetCountryClientError(t *testing.T) {
	// Init

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

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	// Init

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJSONResponse(t *testing.T) {
	// Init

	// Test
	country, err := GetCountry("AR")

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, http.StatusOK, err.Status)
	assert.EqualValues(t, country.Id, "AR")
	assert.EqualValues(t, country.Name, "Argentina")
	assert.EqualValues(t, country.TimeZone, "GMT-03:00")
	assert.EqualValues(t, 24, len(country.States))
}
