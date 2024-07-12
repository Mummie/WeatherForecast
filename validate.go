package main

import (
	"strconv"
)

// ValidateCoordinates takes latitude and longitude as strings and returns true if they are legitimate geo coordinates.
func validateCoordinates(latitude, longitude string) bool {
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return false
	}

	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return false
	}

	if lat < -90 || lat > 90 {
		return false
	}

	if lon < -180 || lon > 180 {
		return false
	}

	return true
}
