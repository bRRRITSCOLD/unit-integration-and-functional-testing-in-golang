package controllers

import (
	"net/http"
	"unit-integration-and-functional-testing-in-golang/internal/api/services"

	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {
	country, err := services.CountriesService.GetCountry(c.Param("countryId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
