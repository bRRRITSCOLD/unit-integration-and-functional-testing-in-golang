package main

import (
	"fmt"
	"unit-integration-and-functional-testing-in-golang/src/api/providers"
)

func main() {
	country, err := providers.GetCountry("AR")
	if err != nil {
		fmt.Println("error getting country", err)
	}
	fmt.Printf("country %v", country)
}
