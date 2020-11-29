package impl
import (
	"encoding/json"
	"fmt"
	"net/http"
	//req"Weather_api/models/Request"
	"Weather_api/models/response"

)



// Provider cho openweathermap.org
type OpenWeatherMapProvider struct {
	APIKey string
	URL    string
}


// Implement hàm GetTemperature của WeatherProvider Interface
func (p OpenWeatherMapProvider) GetTemperature(city string) (float64, error) {
	res, err := http.Get(p.URL + p.APIKey + "&q=" + city)

	if err != nil || res.StatusCode != 200 {
		return 0, err
	}

	defer res.Body.Close()

	data := response.OpenWeatherMapData{}

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return 0, err
	}
	// Tính lại theo độ C
	tempC := data.Current.KelvinTemp - 273.15
	fmt.Println("openweathermap: ", tempC)

	return tempC, err
}


// Provider cho weatherbit.io
type WeatherBitProvider struct {
	APIKey string
	URL    string
}


// Implement hàm GetTemperature của Weather Interface
func (p WeatherBitProvider) GetTemperature(city string) (float64, error) {
	res, err := http.Get(p.URL + p.APIKey + "&city=" + city)

	if err != nil || res.StatusCode != 200 {
		return 0, err
	}

	defer res.Body.Close()

	data := response.WeatherBitData{}

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return 0, err
	}

	fmt.Println("weatherbit: ", data.Current[0].CelsiusTemp)
	return data.Current[0].CelsiusTemp, err
}

type ApiXuProvider struct {
	APIKey string
	URL    string
}



// Implement hàm GetTemperature của WeatherProvider Interface
func (p ApiXuProvider) GetTemperature(city string) (float64, error) {
	res, err := http.Get(p.URL + p.APIKey + "&q=" + city)

	if err != nil || res.StatusCode != 200 {
		return 0, err
	}

	defer res.Body.Close()

	data := response.ApiXuData{}

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return 0, err
	}

	fmt.Println("apixu: ", data.Current.CelsiusTemp)
	return data.Current.CelsiusTemp, err
}