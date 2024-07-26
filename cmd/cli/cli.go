package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tweather/cmd/api"

	"gopkg.in/yaml.v3"
)

type Args struct {
	Key      *string
	Coords   *string
	Location *string
	Zip      *string
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

	flag.Parse()
	args := Args{}

	if *key != "" {
		args.Key = key
	} else {
		fmt.Println("No API key was specified! Head to openweathermap.org to create an account and get an API key.")
		os.Exit(1)
	}

	if *coords != "" {
		switch len(strings.Split(*coords, ",")) {
		case 2:
			args.Coords = coords
		default:
			fmt.Println("Incorrect number of comma separated values in 'coords'. Expected 2, got", len(strings.Split(*coords, ",")))
		}
	}

	if *location != "" {
		switch len(strings.Split(*location, ",")) {
		case 3:
			args.Location = location
		default:
			fmt.Println("Incorrect number of comma separated values in 'location'. Expected 3, got", len(strings.Split(*coords, ",")))
		}
	}

	if *zip != "" {
		args.Zip = zip
	}

	if args.Coords == nil && args.Location == nil && args.Zip == nil {
		fmt.Println("No valid location data was passed! Must pass at least one of coords, location, or zip.")
		os.Exit(1)
	}

	return args
}

func GetApiConfig(args Args) (*api.ApiConfig, error) {
	apiConfig := api.ApiConfig{
		Key: *args.Key,
	}

	// TODO this should be refactored
	if args.Coords != nil {
		parts := strings.Split(*args.Coords, ",")
		lat, err1 := strconv.ParseFloat(parts[0], 32)
		lon, err2 := strconv.ParseFloat(parts[1], 32)
		if err1 == nil && err2 == nil {
			coords := api.Coords{Lat: lat, Lon: lon}
			apiConfig.Coords = coords
			return &apiConfig, nil
		}
	}

	if args.Location != nil {
		parts := strings.Split(*args.Location, ",")
		city, state, country := parts[0], parts[1], parts[2]
		geoloc := api.GetCoordsFromLocation(apiConfig.Key, city, state, country)[0] // Take 1st value from list of results
		coords := api.Coords{Lat: geoloc.Lat, Lon: geoloc.Lon}
		apiConfig.Coords = coords
		return &apiConfig, nil
	}

	if args.Zip != nil {
		// TODO
	}

	return nil, &InvalidCoordsError{}
}
