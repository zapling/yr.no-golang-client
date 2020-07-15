# yr.no-golang-client

A simple unofficial golang client for Yr.no API

# Usage

```golang
package main

import (
    "fmt"
    "github.com/zapling/yr.no-golang-client/pkg/yr"
)

func main() {
    forecast, err := yr.GetLocationForecast(60.1, 9.58, "WeatherApplet/1.0")
    if err != nil {
      panic(err)
    }
    fmt.Printf("%#v", forecast.Geometry)
    // yr.PointGeometry{Type:"Point", Coordinates:[]float64{9.58, 60.1, 496}}
}
```
# Docs

https://api.met.no/weatherapi/locationforecast/2.0/documentation
