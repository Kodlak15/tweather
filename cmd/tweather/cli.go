package tweather

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Args struct {
	Key      *string
	Coords   *string
	Location *string
	Zip      *string
	Opts     *string
	Verbose  bool
}

type TweatherConfig struct {
	Key      *string `yaml:"key"`
	Coords   *string `yaml:"coords"`
	Location *string `yaml:"location"`
	Zip      *string `yaml:"zip"`
}

type InvalidCoordsError struct{}

func (e *InvalidCoordsError) Error() string {
	return "Unable to parse the provided location data"
}

type MissingApiKeyError struct{}

func (e *MissingApiKeyError) Error() string {
	return "Missing API key! Make an account at openweathermap.org to generate an API key."
}

func GetTweatherConfig(cfgPath string) (*TweatherConfig, error) {
	file, err := os.ReadFile(cfgPath)
	if err != nil {
		fmt.Println("Unable to read file")
		return nil, err
	}

	var cfg TweatherConfig
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		fmt.Println("Failed to unmarshall yaml")
		return nil, err
	}

	return &cfg, nil
}

func GetArgs() Args {
	key := flag.String("key", "", "Specify your openweathermap API key")
	coords := flag.String("coords", "", "Specify the (comma-seperated, ex: 0,0) coordinates of the location to retrieve data for")
	location := flag.String("location", "", "Specify the city, state code, and country code (ex: Boston,MA,USA) to retrieve data for")
	zip := flag.String("zip", "", "Specify the zip code to retrieve data for")
	opts := flag.String("opts", "", "Specify the data to fetch from the API")
	verbose := flag.Bool("verbose", false, "Specify whether to print with extra verbosity")
	flag.Parse()

	return Args{
		Key:      key,
		Coords:   coords,
		Location: location,
		Zip:      zip,
		Opts:     opts,
		Verbose:  *verbose,
	}
}

func GetApiConfig(cfg *TweatherConfig) (*ApiConfig, error) {
	apiConfig := ApiConfig{
		Key: *cfg.Key,
	}

	// fmt.Println("Config:", *cfg.Coords)

	// TODO this should be refactored
	if cfg.Coords != nil {
		parts := strings.Split(*cfg.Coords, ",")
		lat, err1 := strconv.ParseFloat(parts[0], 32)
		lon, err2 := strconv.ParseFloat(parts[1], 32)
		if err1 == nil && err2 == nil {
			coords := Coords{Lat: lat, Lon: lon}
			apiConfig.Coords = coords
			return &apiConfig, nil
		}
	}

	if cfg.Location != nil {
		parts := strings.Split(*cfg.Location, ",")
		city, state, country := parts[0], parts[1], parts[2]
		geoloc := GetCoordsFromLocation(apiConfig.Key, city, state, country)[0] // Take 1st value from list of results
		coords := Coords{Lat: geoloc.Lat, Lon: geoloc.Lon}
		apiConfig.Coords = coords
		return &apiConfig, nil
	}

	if cfg.Zip != nil {
		// TODO
	}

	return nil, &InvalidCoordsError{}
}
