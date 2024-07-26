package tweather

import (
	"fmt"
	"os"
)

func Run() {
	cfg, err := GetTweatherConfig("./tweather.yaml")
	if err != nil {
		// Don't panic as config can still be retrieved via args
		fmt.Println(err)
	}

	if cfg == nil {
		cfg = &TweatherConfig{}
	}

	args := GetArgs()

	// Override any config fields with associated argument if that argument is not nil
	if *args.Key != "" {
		cfg.Key = args.Key
	}
	if *args.Coords != "" {
		cfg.Coords = args.Coords
	}
	if *args.Location != "" {
		cfg.Location = args.Location
	}
	if *args.Zip != "" {
		cfg.Zip = args.Zip
	}

	apiConfig, err := GetApiConfig(cfg)
	if err != nil {
		fmt.Println("Error occurred while getting configuring the API request:", err)
		os.Exit(1)
	}

	data := GetCurrentWeather(*apiConfig)
	data.Get(&args)
	// fmt.Println("Current weather data:", data)

	// data := GetForecast(*apiConfig)
	// fmt.Println("Forecast data:", data)
}
