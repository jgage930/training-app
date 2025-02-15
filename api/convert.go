package api

import (
	"errors"
	"log"
)

type Unit int

const (
	Meter Unit = iota
	Mile
	MetersPerSecond
	MilesPerHour
)

func ConvertValue(value float64, current Unit, target Unit) float64 {
	rate, err := rateLookup(current, target)
	if err != nil {
		log.Fatal("Unsupported conversion.")
	}

	return rate * value
}

func rateLookup(current Unit, target Unit) (float64, error) {
	if current == Meter && target == Mile {
		return 0.000621371, nil
	}

	if current == MetersPerSecond && target == MilesPerHour {
		return 2.23694, nil
	}

	return 0, errors.New("Unsupported Conversion")
}
