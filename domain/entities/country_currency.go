package entities

// CountryCurreny represents a country and its currency.
// @Description Object containing country and corresponding currency.
type CountryCurreny struct {
	Country  string `json:"country" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}
