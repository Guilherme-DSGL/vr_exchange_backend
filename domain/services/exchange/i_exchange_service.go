package services

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
)

type IExchangeService interface {
	GetExchange(ctx context.Context, exchangeParams *entities.ExchangeGetParams)
}
