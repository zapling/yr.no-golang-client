# yr.no-golang-client

A unofficial golang client for [Yr.no](https://www.yr.no/) `2.0` API

# Usage

```
$ go get github.com/zapling/yr.no-golang-client
```

```golang
func main() {
    yrClient := client.NewYrClient(http.DefaultClient, "Yr.no golang example")

    forecast, response, err := locationforecast.GetCompact(yrClient, 60.1, 9.58)
	if err != nil {
        fmt.Println(err)
        return
    }

    temperature := forecast.Properties.Timeseries[0].Data.Instant.Details.AirTemperature

    fmt.Printf("%v\n", utils.Float64Value(temperature)) // -3.1
}
```
The client should be pretty easy to use. Returned fields that are optional are defined as `*string` and `*float64` which means that the value can be `nil`. You can use the utility package provided to easily get the pointer value.

# Terms Of Service

Yr.no has a [Terms of service](https://developer.yr.no/doc/TermsOfService/) for their API. You should read and follow these in order to not get blocked from their service. Some of the things you need to do; provide a unique `User-Agent`, rate limit your requests, cache the returned forecast data appropriately.

# Docs

https://developer.yr.no/doc/locationforecast/HowTO/  
https://api.met.no/weatherapi/locationforecast/2.0/documentation  
https://developer.yr.no/doc/TermsOfService/

# Licence

MIT License  
Copyright 2021 Andreas Palm
