// package http

// import "github.com/go-resty/resty/v2"

// var client *resty.Client = resty.New()

// func GetClient(name string) *resty.Client {
// 	return client
// }

// package locations
// type Country struct {

// }
package domains

type GeoLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Country struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State        `json:"state"`
}
