package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain"
	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
	rest "github.com/Guilherme-DSGL/vr_exchange_backend/internal/http_adapter"
	"github.com/Guilherme-DSGL/vr_exchange_backend/utils"
)

var (
	baseURL = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"
)

type ExchangeRepository struct {
	httpClient rest.HttpClient
}

func NewExchangeRepository(client rest.HttpClient) *ExchangeRepository {
	return &ExchangeRepository{
		httpClient: client,
	}
}

func (er *ExchangeRepository) GetExchange(ctx context.Context, params *entities.ExchangeRequestParams) (entities.ExchangeTransaction, error) {

	query := fmt.Sprintf(
		"?fields=country,currency,exchange_rate,record_date&filter=record_date:gte:%s,record_date:lte:%s,country:eq:%s,currency:eq:%s&sort=-record_date",
		params.StartDate, params.EndDate, params.CountryCurrency.Country, params.CountryCurrency.Currency,
	)

	url := fmt.Sprintf("%s%s", baseURL, query)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to create request")
	}

	resp, err := er.httpClient.Do(req)
	if err != nil {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to request api")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to get exchange request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to convert api body")
	}

	var apiResp exchangeApiResponse

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to convert json")
	}

	if len(apiResp.Data) == 0 {
		return entities.ExchangeTransaction{}, domain.ErrNotFound
	}

	latest := apiResp.Data[0]

	rate, err := strconv.ParseFloat(latest.ExchangeRate, 64)

	convertedValue := utils.RoundValue2DecimalPlaces(params.Value * rate)
	if err != nil {
		return entities.ExchangeTransaction{}, fmt.Errorf("failed to convert data")
	}
	exchange := entities.ExchangeTransaction{
		CountryCurreny: entities.CountryCurreny{
			Country:  latest.Country,
			Currency: latest.Currency,
		},
		EfectiveDate: utils.ParseDate(latest.RecordDate),
		Rate:         rate,
	}

	exchange.ConvertedValue = convertedValue
	return exchange, nil
}

type exchangeApiResponse struct {
	Data []struct {
		Country      string `json:"country"`
		Currency     string `json:"currency"`
		ExchangeRate string `json:"exchange_rate"`
		RecordDate   string `json:"record_date"`
	} `json:"data"`
}
