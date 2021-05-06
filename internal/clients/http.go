package clients

import "github.com/go-resty/resty/v2"

var restyClient *resty.Client

func GetHTTPClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}
