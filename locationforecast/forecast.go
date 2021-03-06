package locationforecast

import (
	"encoding/json"
	"fmt"

	"github.com/zapling/yr.no-golang-client/client"
	"github.com/zapling/yr.no-golang-client/utils"
)

func GetCompact(client *client.YrClient, cords utils.Coordinates) (*GeoJson, error) {
	var forecast GeoJson

	apiUrl := fmt.Sprintf(
		"%s/locationforecast/%s/compact?lat=%.2f&lon=%.2f",
		"https://api.met.no/weatherapi",
		"2.0",
		cords.Latitude,
		cords.Longitude,
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
