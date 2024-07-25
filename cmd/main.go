package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO build CLI + configuration system + (maybe) TUI

type CurrentWeather struct {
	Timestamp int32     `json:"dt"`
	Coords    Coords    `json:"coord"`
	Weather   []Weather `json:"weather"`
	Main      Main      `json:"main"`
}

type Forecast struct {
	List []struct {
		Timestamp int32     `json:"dt"`
		Weather   []Weather `json:"weather"`
		Main      Main      `json:"main"`
	} `json:"list"`
}

type Coords struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
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

func getCurrentWeather(key string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s",
		0.0, 0.0, key, "imperial")

	// TODO check and handle specific error codes
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

	var data CurrentWeather
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	fmt.Println(data)
}

func getForecast(key string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=%s",
		0.0, 0.0, key, "imperial")

	// TODO check and handle specific error codes
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

	var data Forecast
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON from response body:", err)
		os.Exit(1)
	}

	fmt.Println(data)
}

func main() {
	key := os.Args[1]

	getCurrentWeather(key)
	getForecast(key)
}
