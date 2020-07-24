package yr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var api_url = "https://api.met.no/weatherapi"
var api_version = "2.0"

// LocationForecast => METJSONForecast
type LocationForecast struct {
	Type       string        `json:"type"`
	Geometry   PointGeometry `json:"geometry"`
	Properties Properties    `json:"properties"`
}

type PointGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"` // [longitude, latitude, altitude]
}

// Properties => Forecast
type Properties struct {
	Meta       Meta       `json:"meta"`
	Timeseries []Timestep `json:"timeseries"`
}

// Meta => inline_model_0
type Meta struct {
	UpdatedAt string        `json:"updated_at"`
	Units     ForecastUnits `json:"units"`
}

type ForecastUnits struct {
	PrecipitationAmount    string `json:"precipitation_amount,omitempty"`
	PrecipitationAmountMin string `json:"precipitation_amount_min"`
	PrecipitationAmountMax string `json:"precipitation_amount_max"`

	CloudAreaFractionLow    string `json:"cloud_area_fraction_low"`
	CloudAreaFractionMedium string `json:"cloud_area_fraction_medium"`
	CloudAreaFractionHigh   string `json:"cloud_area_fraction_high"`

	FogAreaFraction   string `json:"fog_area_fraction"`
	CloudAreaFraction string `json:"cloud_area_fraction"`

	WindSpeed         string `json:"wind_speed"`
	WindSpeedOfGust   string `json:"wind_speed_of_gust"`
	WindFromDirection string `json:"wind_from_direction"`

	AirTemperature        string `json:"air_temperature"`
	AirTemperatureMin     string `json:"air_temperature_min"`
	AirTemperatureMax     string `json:"air_temperature_max"`
	AirPressureAtSeaLevel string `json:"air_pressure_at_sea_level"`

	ProbabillityOfPrecipitation string `json:"probability_of_precipitation"`
	ProbabillityOfThunder       string `json:"probability_of_thunder"`

	RelativeHumidity            string `json:"relative_humidity"`
	DewPointTemperature         string `json:"dew_point_temperature"`
	UltravioletIndexClearSkyMax string `json:"ultraviolet_index_clear_sky_max"`
}

// Timestep => ForecastTimestep
type Timestep struct {
	Time string       `json:"time"`
	Data TimestepData `json:"data"`
}

// TimestepData => inline_model
type TimestepData struct {
	Instant     InstantData    `json:"instant"`
	Next1Hours  Next1HoursData `json:"next_1_hours"`
	Next6Hours  Next1HoursData `json:"next_6_hours"`
	Next12Hours Next1HoursData `json:"next_12_hours"`
}

// InstantData => Inline Model 2
type InstantData struct {
	Details ForecastTimeInstant `json:"details"`
}

type ForecastTimeInstant struct {
	AirTemperature          float64 `json:"air_temperature"`
	AirPressureAtSeaLevel   float64 `json:"air_pressure_at_sea_level"`
	CloudAreaFraction       float64 `json:"cloud_area_fraction"`
	WindSpeed               float64 `json:"wind_speed"`
	RelativeHumidity        float64 `json:"relative_humidity"`
	DewPointTemperature     float64 `json:"dew_point_temperature"`
	WindFromDirection       float64 `json:"wind_from_direction"`
	FogAreaFraction         float64 `json:"fog_area_fraction"`
	CloudAreaFractionHigh   float64 `json:"cloud_area_fraction_hight"`
	WindSpeedOfGust         float64 `json:"wind_speed_of_gust"`
	CloudAreaFractionMedium float64 `json:"cloud_area_fraction_medium"`
	CloudAreaFractionLow    float64 `json:"cloud_area_fraction_low"`
}

// Next1HoursData => Inline Model 3
type Next1HoursData struct {
	Details ForecastTimePeriod `json:"details"`
	Summary ForecastSummary    `json:"summary"`
}

type ForecastTimePeriod struct {
	ProbabillityOfPrecipitation float64 `json:"probability_of_precipitation"`
	AirTemperatureMin           float64 `json:"air_temperature_min"`
	UltravioletIndexClearSkyMax float64 `json:"ultraviolet_index_clear_sky_max"`
	AirTemperatureMax           float64 `json:"air_temperature_max"`
	PrecipitationAmountMin      float64 `json:"precipitation_amount_min"`
	ProbabillityOfThunder       float64 `json:"probability_of_thunder"`
	PrecipitationAmountMax      float64 `json:"precipitation_amount_max"`
	PrecipitationAmount         float64 `json:"precipitation_amount"`
}

type ForecastSummary struct {
	SymbolCode string `json:"symbol_code"`
}

func request(url string, queryParams map[string]string, userAgent string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	// Add query params
	reqQuery := req.URL.Query()
	for k, v := range queryParams {
		reqQuery.Add(k, v)
	}
	req.URL.RawQuery = reqQuery.Encode()

	// Set User-Agent
	req.Header.Set("User-Agent", userAgent)

	// Send request
	res, err := http.DefaultClient.Do(req)
	return res, err
}

func GetLocationForecast(lat float64, lon float64, userAgent string) (LocationForecast, error) {
	var Forecast LocationForecast
	url := fmt.Sprintf("%s/locationforecast/%s/compact", api_url, api_version)

	queryParams := map[string]string{
		"lat": fmt.Sprintf("%.2f", lat),
		"lon": fmt.Sprintf("%.2f", lon),
	}

	res, err := request(url, queryParams, userAgent)

	if res.StatusCode != 200 && res.StatusCode != 203 {
		return Forecast, errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Forecast, err
	}

	json.Unmarshal(body, &Forecast)

	return Forecast, nil
}
