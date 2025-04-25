package domain

type ExchangeTransaction struct {
	Transaction    Transaction    `json:"transaction"`
	CountryCurreny CountryCurreny `json:"country_currency"`
	Rate           float64        `json:"rate"`
	ConvertedValue float64        `json:"converted_value"`
}

// get https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=country,currency,exchange_rate,record_date&filter=record_date:gte:2025-01-01,record_date:lte:2025-03-31,country:eq:Brazil,currency:eq:Real
