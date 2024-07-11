package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Period struct {
	Temperature     int    `json:"temperature"`
	TemperatureUnit string `json:"temperatureUnit"`
	ShortForecast   string `json:"shortForecast"`
	Characterize    string `json:"characterization"`
}

type ForecastProperties struct {
	Periods []Period `json:"periods"`
}

type ForecastResponse struct {
	Properties ForecastProperties `json:"properties"`
}

type Properties struct {
	Forecast string `json:"forecast"`
}

type GeoResponse struct {
	Properties Properties `json:"properties"`
}

var (
	ErrInvalidGeoPoints   = errors.New("got nil or empty URL from GeoPoints for coordinates")
	ErrInvalidResponse    = errors.New("invalid response from weather service")
	ErrInvalidCoordinates = errors.New("invalid coordinates")
	ErrForecastAPI        = errors.New("error fetching forecast data")
	geoPointUrl           = `https://api.weather.gov/points/%s,%s`
)

const userAgent = "(myweatherapp.com, contact@myweatherapp.com)"

func GetWeatherData(latitude, longitude string) (f *ForecastResponse, err error) {
	if !ValidateCoordinates(latitude, longitude) {
		return nil, ErrInvalidCoordinates
	}
	client := &http.Client{}

	geoUrl, err := GetGeoPoints(latitude, longitude)
	if err != nil {
		return nil, err
	}

	// Check if geoUrl is nil
	if geoUrl == nil || geoUrl.Properties.Forecast == "" {
		return nil, errors.Wrap(ErrInvalidGeoPoints, fmt.Sprintf("got err for invalid coordinates %v, %v", latitude, longitude))
	}

	req, err := http.NewRequest("GET", geoUrl.Properties.Forecast, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("got error creating request to %v", geoUrl.Properties.Forecast))
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("got error sending request to %v", geoUrl.Properties.Forecast))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(ErrInvalidResponse, fmt.Sprintf("status code: %d", resp.StatusCode))
	}

	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return nil, fmt.Errorf("error decoding point data: %w", err)
	}

	return f, nil
}

func GetGeoPoints(latitude, longitude string) (r *GeoResponse, err error) {
	client := &http.Client{}

	url := fmt.Sprintf(geoPointUrl, latitude, longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "got error creating request to GeoPoints URL")
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "got error sending client request to GeoPoints URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "got non 200 status code from GeoPoints URL")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "got error decoding GeoPoints response to struct")
	}

	return r, nil
}
