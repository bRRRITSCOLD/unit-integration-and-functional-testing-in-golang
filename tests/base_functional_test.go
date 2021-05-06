package tests

import (
	"os"
	"testing"
	"unit-integration-and-functional-testing-in-golang/internal/api/app"
	"unit-integration-and-functional-testing-in-golang/internal/clients"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	httpClient := clients.GetHTTPClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	go app.StartApp()

	os.Exit(m.Run())
}
