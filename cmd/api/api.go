package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ApiConfig struct {
	Key    string
	Coords Coords
}

type CurrentWeather struct {
	Timestamp int32     `json:"dt"`
	Coords    Coords    `json:"coord"`
	Weather   []Weather `json:"weather"`
	Main      Main      `json:"main"`
	Name      string    `json:"name"`
}

type Forecast struct {
	List []struct {
		Timestamp int32     `json:"dt"`
		Weather   []Weather `json:"weather"`
		Main      Main      `json:"main"`
	} `json:"list"`
}

type Coords struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int32  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int32   `json:"pressure"`
	Humidity  int32   `json:"humidity"`
}

func GetCurrentWeather(apiConfig ApiConfig) CurrentWeather {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s",
		apiConfig.Coords.Lat, apiConfig.Coords.Lon, apiConfig.Key, "imperial")
	body := getApiResponseBody(url)

	var data CurrentWeather
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	return data
}

func GetForecast(apiConfig ApiConfig) Forecast {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=%s",
		0.0, 0.0, apiConfig.Key, "imperial")
	body := getApiResponseBody(url)

	var data Forecast
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	return data
}

type Geolocation []struct {
	Lat float64
	Lon float64
}

func GetCoordsFromLocation(key string, city string, state string, country string) Geolocation {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&limit=%d&appid=%s",
		city, state, country, 1, key)
	body := getApiResponseBody(url)

	var data Geolocation
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	return data
}

func GetCoordsFromZip(key string, zip string, country string) Geolocation {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s",
		zip, country, key)
	body := getApiResponseBody(url)

	// This may not work since the api does not return a list in this case
	var data Geolocation
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	return data
}

func getApiResponseBody(url string) []byte {
	// TODO: check and handle specific error codes
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		os.Exit(1)
	}

	return body
}
