package entities

import "time"

// ExchangeTransaction represents a transaction with exchange rate applied.
// @Description Transaction with exchange rate information.
type ExchangeTransaction struct {
	Transaction    Transaction    `json:"transaction"`
	CountryCurreny CountryCurreny `json:"country_currency"`
	EfectiveDate   time.Time      `json:"efective_date"`
	Rate           float64        `json:"rate"`
	ConvertedValue float64        `json:"converted_value"`
}
