package domain

type ExchangeTransaction struct {
	Transaction    Transaction `json:"transaction"`
	Rate           float64     `json:"rate"`
	ConvertedValue float64     `json:"converted_value"`
}
