package irepo

import (
	"context"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
)

//go:generate mockery --name IExchangeRepository
type IExchangeRepository interface {
	GetExchange(ctx context.Context, exchangedParams *entities.ExchangeRequestParams) (entities.ExchangeTransaction, error)
}
