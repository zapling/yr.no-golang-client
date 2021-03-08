package locationforecast

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zapling/yr.no-golang-client/client"
)

// GetCompact get weather forecast for a specific place specified by the latitude and longitude
// coordinates.
// On success you should cache the forecast for one hour beforing re-fetching it.
//
// https://api.met.no/weatherapi/locationforecast/2.0/documentation#!/data/get_compact
func GetCompact(
	client *client.YrClient,
	latitude float64,
	longitude float64,
) (*GeoJson, *http.Response, error) {
	var forecast GeoJson

	apiUrl := fmt.Sprintf(
		"%s/locationforecast/%s/compact?lat=%.2f&lon=%.2f",
		"https://api.met.no/weatherapi",
		"2.0",
		latitude,
		longitude,
	)

	response, err := client.GetRequest(apiUrl)
	if err != nil {
		if response == nil {
			return nil, nil, err
		}

		return nil, response, err
	}

	if err = json.NewDecoder(response.Body).Decode(&forecast); err != nil {
		return nil, response, err
	}

	return &forecast, response, nil
}
