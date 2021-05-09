package main

import (
	"fmt"
	sequel "unit-integration-and-functional-testing-in-golang/internal/sql"
)

// "unit-integration-and-functional-testing-in-golang/internal/api/app"

func main() {
	user, err := sequel.GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	// app.StartApp()
}
