package main

type TemperatureRanges struct {
	Hot      Range `json:"hot"`
	Moderate Range `json:"moderate"`
	Cold     Range `json:"cold"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

var defaultConfig = TemperatureRanges{
	Hot: Range{
		Min: 85,
		Max: 110,
	},
	Moderate: Range{
		Min: 60,
		Max: 84,
	},
	Cold: Range{
		Min: 0,
		Max: 59,
	},
}

var config = TemperatureRanges{}
