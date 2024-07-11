package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	TemperatureRanges struct {
		Hot struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"hot"`
		Moderate struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"moderate"`
		Cold struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"cold"`
	} `json:"temperature_ranges"`
}

var config Config

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/forecast", weatherHandler).Methods("GET")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	scale := r.URL.Query().Get("scale")

	if latitude == "" || longitude == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Latitude and Longitude are required parameters")
		return
	}

	if scale == "" {
		scale = "fahrenheit"
	}

	weatherData, err := GetWeatherData(latitude, longitude)
	if err != nil {
		log.Println(errors.Unwrap(err))
		sendErrorResponse(w, http.StatusInternalServerError, "Error fetching weather data")
		return
	}

	shortForecast := weatherData.Data.Weather
	temperature := weatherData.CurrentObservation.Temp
	temperatureUnit := "F"

	// tempCharacterization := characterizeTemperature(temperature, scale)

	response := map[string]interface{}{
		"shortForecast":       shortForecast,
		"temperature":         temperature,
		"temperatureUnit":     scale,
		"temperatureCategory": temperatureUnit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func characterizeTemperature(temp float64, scale string) string {
	var hotMin, hotMax, modMin, modMax, coldMin, coldMax int

	if scale == "celsius" {
		hotMin = (config.TemperatureRanges.Hot.Min - 32) * 5 / 9
		hotMax = (config.TemperatureRanges.Hot.Max - 32) * 5 / 9
		modMin = (config.TemperatureRanges.Moderate.Min - 32) * 5 / 9
		modMax = (config.TemperatureRanges.Moderate.Max - 32) * 5 / 9
		coldMin = (config.TemperatureRanges.Cold.Min - 32) * 5 / 9
		coldMax = (config.TemperatureRanges.Cold.Max - 32) * 5 / 9
	} else {
		hotMin = config.TemperatureRanges.Hot.Min
		hotMax = config.TemperatureRanges.Hot.Max
		modMin = config.TemperatureRanges.Moderate.Min
		modMax = config.TemperatureRanges.Moderate.Max
		coldMin = config.TemperatureRanges.Cold.Min
		coldMax = config.TemperatureRanges.Cold.Max
	}

	switch {
	case temp >= float64(hotMin) && temp <= float64(hotMax):
		return "hot"
	case temp >= float64(modMin) && temp <= float64(modMax):
		return "moderate"
	case temp >= float64(coldMin) && temp <= float64(coldMax):
		return "cold"
	default:
		return "unknown"
	}
}
