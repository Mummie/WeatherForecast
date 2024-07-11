package main

func characterizeTemperature(temp int, scale string) string {
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
