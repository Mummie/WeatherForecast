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

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	if config.Cold.Max == 0 {
		config = defaultConfig
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

	weatherData, err := GetWeatherData(latitude, longitude)
	if err != nil {
		log.Println(errors.Unwrap(err))
		sendErrorResponse(w, err)
		return
	}

	shortForecast := weatherData.Properties.Periods[0].ShortForecast
	temperature := weatherData.Properties.Periods[0].Temperature
	temperatureUnit := "F"

	// Normally wouldn't do this and would have a separate file/package to handle responses but for simplicity sake
	response := map[string]interface{}{
		"shortForecast":    shortForecast,
		"temperature":      temperature,
		"temperatureUnit":  temperatureUnit,
		"characterization": characterizeTemperature(temperature),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
