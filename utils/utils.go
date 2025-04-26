package utils

import "math"

func RoundValue2DecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}
