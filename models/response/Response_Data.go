package response


type TemperatureData struct {
	CityName       string  `json:"city_name"`
	CelsiusTemp    float64 `json:"celsius_temp"`
	KelvinTemp     float64 `json:"kelvin_temp"`
	FahrenheitTemp float64 `json:"fahrenheit_temp"`
}