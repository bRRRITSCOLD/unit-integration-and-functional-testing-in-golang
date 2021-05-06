package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"unit-integration-and-functional-testing-in-golang/internal/api/domains"
	"unit-integration-and-functional-testing-in-golang/internal/api/services"
	"unit-integration-and-functional-testing-in-golang/internal/api/utils/errors"
	"unit-integration-and-functional-testing-in-golang/internal/clients"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*domains.Country, *errors.APIError)
)

type countriesServiceMock struct {
}

func (*countriesServiceMock) GetCountry(countryId string) (*domains.Country, *errors.APIError) {
	return getCountryFunc(countryId)
}

func TestMain(m *testing.M) {
	httpClient := clients.GetHTTPClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	// Init
	countryId := "AR"

	getCountryFunc = func(countryId string) (*domains.Country, *errors.APIError) {
		return nil, &errors.APIError{
			Status:  http.StatusNotFound,
			Message: "Country not found",
			Error:   "not_found",
		}
	}
	services.CountriesService = &countriesServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "countryId", Value: countryId},
	}

	// Test
	GetCountry(c)

	// Validation
	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.APIError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountry(t *testing.T) {
	// Init
	countryId := "AR"

	getCountryFunc = func(countryId string) (*domains.Country, *errors.APIError) {
		return &domains.Country{
			Id:   "AR",
			Name: "Argentina",
		}, nil
	}
	services.CountriesService = &countriesServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "countryId", Value: countryId},
	}

	// Test
	GetCountry(c)

	// Validation
	assert.EqualValues(t, http.StatusOK, response.Code)

	var country domains.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
