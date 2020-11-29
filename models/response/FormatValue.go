package response

// mapping data form openweathermap.org response

type OpenWeatherMapData struct {
	Current struct {
		KelvinTemp float64 `json:"temp"`
	} `json:"main"`
}

// mapping data from apixu.com response
type ApiXuData struct {
	Current struct {
		CelsiusTemp float64 `json:"temp_c"`
	} `json:"current"`
}


// mapping data from weatherbit.io response

type WeatherBitData struct {
	Current []struct {
		CelsiusTemp float64 `json:"temp"`
	} `json:"data"`
}
