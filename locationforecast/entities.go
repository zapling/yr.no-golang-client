package locationforecast

type GeoJson struct {
	Type       string     `json:"type"`
	Geometry   GeoPoint   `json:"geometry"`
	Properties Properties `json:"properties"`
}

// GeoPoint represents a geographical location.
// Coordinates are [longitude, latitude, altitude]
type GeoPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Properties holds all forecast data and meta data related to those.
type Properties struct {
	Meta       ForecastMeta `json:"meta"`
	Timeseries []Forecast   `json:"timeseries"`
}

// ForecastMeta holds information about when the forecast data was last updated by Yr.
// It also holds units of all forecast parameters.
// The datetime format for the UpdatedAt field is "2020-06-10T13:04:26Z"
type ForecastMeta struct {
	UpdatedAt string `json:"updated_at"`
	Units     struct {
		AirPressureAtSeaLevel       string `json:"air_pressure_at_sea_level"`
		AirTemperature              string `json:"air_temperature"`
		AirTemperatureMax           string `json:"air_temperature_max"`
		AirTemperatureMin           string `json:"air_temperature_min"`
		CloudAreaFraction           string `json:"cloud_area_fraction"`
		CloudAreaFractionHigh       string `json:"cloud_area_fraction_high"`
		CloudAreaFractionMedium     string `json:"cloud_area_fraction_medium"`
		CloudAreaFractionLow        string `json:"cloud_area_fraction_low"`
		DewPointTemperature         string `json:"dew_point_temperature"`
		FogAreaFraction             string `json:"fog_area_fraction"`
		PrecipitationAmount         string `json:"precipitation_amount"`
		PrecipitationAmountMin      string `json:"precipitation_amount_min"`
		PrecipitationAmountMax      string `json:"precipitation_amount_max"`
		ProbabilityOfPrecipitation  string `json:"probability_of_precipitation"`
		ProbabilityOfThunder        string `json:"probability_of_thunder"`
		RelativeHumidity            string `json:"relative_humidity"`
		UltravioletIndexClearSkyMax string `json:"ultraviolet_index_clear_sky_max"`
		WindFromDirection           string `json:"wind_from_direction"`
		WindSpeed                   string `json:"wind_speed"`
		WindSpeedOfGust             string `json:"wind_speed_of_gust"`
	} `json:"units"`
}

// Forecast holds a weather forecast.
// The available forecast data may vary depending on how far in the future
// the forecast is for. Forecasts nearer in time holds more information.
type Forecast struct {
	Time string `json:"time"`
	Data struct {
		Instant     *Parameters `json:"instant"`
		Next1Hours  *Parameters `json:"next_1_hours"`
		Next6Hours  *Parameters `json:"next_6_hours"`
		Next12Hours *Parameters `json:"next_12_hours"`
	} `json:"data"`
}

type Parameters struct {
	Summary *struct {
		SymbolCode       *string `json:"symbol_code"`
		SymbolConfidence *string `json:"symbol_confidence"`
	} `json:"summary"`
	Details struct {
		AirPressureAtSeaLevel       *float64 `json:"air_pressure_at_sea_level"`
		AirTemperature              *float64 `json:"air_temperature"`
		AirTemperatureMin           *float64 `json:"air_temperature_min"`
		AirTemperatureMax           *float64 `json:"air_temperature_max"`
		CloudAreaFraction           *float64 `json:"cloud_area_fraction"`
		CloudAreaFractionHigh       *float64 `json:"cloud_area_fraction_hight"`
		CloudAreaFractionMedium     *float64 `json:"cloud_area_fraction_medium"`
		CloudAreaFractionLow        *float64 `json:"cloud_area_fraction_low"`
		DewPointTemperature         *float64 `json:"dew_point_temperature"`
		FogAreaFraction             *float64 `json:"fog_area_fraction"`
		PrecipitationAmount         *float64 `json:"precipitation_amount"`
		PrecipitationAmountMin      *float64 `json:"precipitation_amount_min"`
		PrecipitationAmountMax      *float64 `json:"precipitation_amount_max"`
		ProbabilityOfPrecipitation  *float64 `json:"probability_of_precipitation"`
		ProbabilityOfThunder        *float64 `json:"probability_of_thunder"`
		RelativeHumidity            *float64 `json:"relative_humidity"`
		UltravioletIndexClearSkyMax *float64 `json:"ultraviolet_index_clear_sky_max"`
		WindFromDirection           *float64 `json:"wind_from_direction"`
		WindSpeed                   *float64 `json:"wind_speed"`
		WindSpeedOfGust             *float64 `json:"wind_speed_of_gust"`
	} `json:"details"`
}
