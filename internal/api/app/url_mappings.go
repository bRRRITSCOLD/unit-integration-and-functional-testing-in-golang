package app

import "unit-integration-and-functional-testing-in-golang/internal/api/controllers"

func mapUrls() {
	router.GET("/countries/:countryId", controllers.GetCountry)
}
