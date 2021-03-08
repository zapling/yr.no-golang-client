package main

import (
	"fmt"
	"net/http"

	"github.com/zapling/yr.no-golang-client/client"
	"github.com/zapling/yr.no-golang-client/locationforecast"
)

func main() {
	yrClient := client.NewYrClient(http.DefaultClient, "Yr.no go client example")

	forecast, err := locationforecast.GetCompact(yrClient, 60.1, 9.58)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", forecast.Geometry)
}
