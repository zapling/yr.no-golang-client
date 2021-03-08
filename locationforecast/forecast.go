package locationforecast

import (
	"encoding/json"
	"fmt"

	"github.com/zapling/yr.no-golang-client/client"
)

func GetCompact(client *client.YrClient, latitude float64, longitude float64) (*GeoJson, error) {
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
		return nil, err
	}

	if err = json.NewDecoder(response.Body).Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}
