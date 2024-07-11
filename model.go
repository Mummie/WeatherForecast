package main

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
