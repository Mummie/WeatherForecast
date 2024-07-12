package main

func characterizeTemperature(temp int) string {
	hotMin := config.Hot.Min
	hotMax := config.Hot.Max
	modMin := config.Moderate.Min
	modMax := config.Moderate.Max
	coldMin := config.Cold.Min
	coldMax := config.Cold.Max

	switch {
	case temp >= hotMin && temp <= hotMax:
		return "hot"
	case temp >= modMin && temp <= modMax:
		return "moderate"
	case temp >= coldMin && temp <= coldMax:
		return "cold"
	default:
		return "unknown"
	}
}
