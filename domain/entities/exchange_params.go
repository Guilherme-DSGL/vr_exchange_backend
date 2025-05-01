package entities

// ExchangeGetParams represents the input for requesting an exchange.
// @Description Parameters to retrieve an exchange rate for a transaction.
type ExchangeGetParams struct {
	IdTranscation   string         `json:"id" validate:"required"`
	CountryCurrency CountryCurreny `json:"country_currency" validate:"required"`
}
type ExchangeRequestParams struct {
	CountryCurrency CountryCurreny
	Value           float64
	StartDate       string
	EndDate         string
}
