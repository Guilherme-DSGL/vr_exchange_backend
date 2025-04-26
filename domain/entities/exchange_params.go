package entities

import "time"

type ExchangeGetParams struct {
	UIdTranscation  string
	CountryCurrency CountryCurreny
}
type ExchangeRequestParams struct {
	CountryCurrency CountryCurreny
	Value           float64
	Date            time.Time
}
