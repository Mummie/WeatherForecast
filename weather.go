package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Period struct {
	Temperature     int    `json:"temperature"`
	TemperatureUnit string `json:"temperatureUnit"`
	ShortForecast   string `json:"shortForecast"`
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
	ErrInvalidGeoPoints = errors.New("invalid geo point data")
	ErrWeatherAPIPoints = errors.New("error fetching point data")
	ErrInvalidResponse  = errors.New("invalid response from weather service")
	ErrForecastAPI      = errors.New("error fetching forecast data")
	geoPointUrl         = `https://api.weather.gov/points/%s,%s`
)

const userAgent = "(myweatherapp.com, contact@myweatherapp.com)"

func GetWeatherData(latitude, longitude string) (*ForecastResponse, error) {
	client := &http.Client{}

	geoUrl, err := GetGeoPoints(latitude, longitude)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidGeoPoints, err.Error())
	}

	log.Println("making request to ", geoUrl.Properties.Forecast)
	req, err := http.NewRequest("GET", geoUrl.Properties.Forecast, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching point data: %w", err)
	}
	defer resp.Body.Close()

	log.Println(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(ErrInvalidResponse, fmt.Sprintf("status code: %d", resp.StatusCode))
	}

	var forecastData ForecastResponse
	err = json.NewDecoder(resp.Body).Decode(&forecastData)
	if err != nil {
		return nil, fmt.Errorf("error decoding point data: %w", err)
	}

	log.Println(forecastData)

	return &forecastData, nil
}

func GetGeoPoints(latitude, longitude string) (*GeoResponse, error) {
	client := &http.Client{}

	url := fmt.Sprintf(geoPointUrl, latitude, longitude)

	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching point data: %w", err)
	}
	defer resp.Body.Close()

	log.Println(url)
	log.Println(resp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(ErrInvalidResponse, fmt.Sprintf("status code: %d", resp.StatusCode))
	}

	var r GeoResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("error decoding point data: %w", err)
	}

	log.Println(r.Properties.Forecast)
	return &r, nil
}
