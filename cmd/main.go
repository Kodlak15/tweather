package main

import (
	"fmt"
	"os"
	"tweather/cmd/api"
	"tweather/cmd/cli"
)

func main() {
	cfg, err := cli.GetTweatherConfig("./tweather.yaml")
	if err != nil {
		// Don't panic as config can still be retrieved via args
		fmt.Println(err)
	}

	if cfg == nil {
		cfg = &cli.TweatherConfig{}
	}

	args := cli.GetArgs()

	// Override any config fields with associated argument if that argument is not nil
	if args.Key != nil {
		cfg.Key = args.Key
	}
	if args.Coords != nil {
		cfg.Coords = args.Coords
	}
	if args.Location != nil {
		cfg.Location = args.Location
	}
	if args.Zip != nil {
		cfg.Zip = args.Zip
	}

	apiConfig, err := cli.GetApiConfig(args)
	if err != nil {
		fmt.Println("Error occurred while getting configuring the API request:", err)
		os.Exit(1)
	}

	data := api.GetCurrentWeather(*apiConfig)
	fmt.Println("Current weather data:", data)
	// getForecast(key)
}
