package http

import "github.com/go-resty/resty/v2"

var restyClient *resty.Client

func GetClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}
