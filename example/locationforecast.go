package main

import (
	"fmt"
	"net/http"

	"github.com/zapling/yr.no-golang-client/client"
	"github.com/zapling/yr.no-golang-client/locationforecast"
	"github.com/zapling/yr.no-golang-client/utils"
)

func main() {
	yrClient := client.NewYrClient(http.DefaultClient, "Yr.no golang example")

	forecast, response, err := locationforecast.GetCompact(yrClient, 60.1, 9.58)
	if err != nil {
		fmt.Println(err)
		return
	}

	temperature := forecast.Properties.Timeseries[0].Data.Instant.Details.AirTemperature

	fmt.Printf("%v\n", utils.Float64Value(temperature))
	fmt.Printf("Forecast can be re-fetched after: %v", response.Header.Get("expires"))
}
