package utils

import (
	"math"
	"time"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05.999Z07:00"
	DateFormat     = "2006-01-02"
)

func RoundValue2DecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}
func ParseDate(dateStr string) time.Time {
	t, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
