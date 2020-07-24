# yr.no-golang-client

A simple unofficial golang client for Yr.no `2.0` API

# Usage

```
$ go get github.com/zapling/yr.no-golang-client
```

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

# Todo

* Tests
* Rate limiting?
* ~~Status `203` should be ok as it's used as a version [deprecation warning](https://api.met.no/doc/TermsOfService#generalinformationabouttheservice) by YR.~~
* Better handling of optional fields, should maybe be `nil` instead of empty value `string => ""`

# Contribution

Feel free to open an PR for improvments.

# Docs

https://api.met.no/weatherapi/locationforecast/2.0/documentation

# Licence

MIT License

Copyright 2020 Andreas Palm
