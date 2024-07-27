package tweather

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Failed to load the user config directory")
	}

	// TODO allow user to specify config file from another location
	cfgPath := filepath.Join(userCfgDir, "tweather", "tweather.yaml")
	cfg, err := GetTweatherConfig(cfgPath)
	if err != nil {
		// Don't panic as config can still be retrieved via args
		fmt.Println(err)
	}

	if cfg == nil {
		cfg = &TweatherConfig{}
	}

	args := GetArgs()
	if *args.Opts == "" {
		flag.Usage()
	}

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

	// TODO allow user to choose api (current vs forecast)
	data := GetCurrentWeather(*apiConfig)
	data.Get(&args)
	// data := GetForecast(*apiConfig)
	// fmt.Println("Forecast data:", data)
}
