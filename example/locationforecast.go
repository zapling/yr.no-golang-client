package main

import (
	"fmt"
	"net/http"

	"github.com/zapling/yr.no-golang-client/client"
	"github.com/zapling/yr.no-golang-client/locationforecast"
	"github.com/zapling/yr.no-golang-client/utils"
)

func main() {
	cords := utils.Coordinates{Latitude: 60.1, Longitude: 9.58}

	yrClient := client.NewYrClient(http.DefaultClient, "Yr.no go client example")

	forecast, err := locationforecast.GetCompact(yrClient, cords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", forecast)
}
