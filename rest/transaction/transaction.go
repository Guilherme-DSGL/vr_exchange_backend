package rest

import (
	"net/http"
	"strconv"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
	es "github.com/Guilherme-DSGL/vr_exchange_backend/domain/services/exchange"
	ts "github.com/Guilherme-DSGL/vr_exchange_backend/domain/services/transaction"
	"github.com/Guilherme-DSGL/vr_exchange_backend/rest"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

type TransactionHandler struct {
	Ts ts.ITransactionService
	Es es.IExchangeService
}

const defaultLimit = 30

func NewTransactionHandler(e *echo.Group, ts ts.ITransactionService, es es.IExchangeService) {
	handler := &TransactionHandler{
		Ts: ts,
		Es: es,
	}
	e.GET("/transaction", handler.FetchTransaction)
	e.POST("/transaction", handler.Save)
	e.GET("/transaction/:id", handler.GetById)
	e.DELETE("/transaction/:id", handler.Delete)
	e.POST("/exchange", handler.Exchange)
}

// FetchTransaction godoc
// @Summary      List transactions
// @Description  Retrieves a paginated list of transactions based on cursor
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        limit     query    int   false  "Number of items to return"
// @Param        cursor  query     string  false  "Cursor for pagination"
// @Success      200     {array}   entities.Transaction
// @Failure      400,404,500  {object}  rest.ResponseError
// @Router       /transaction [get]
func (th *TransactionHandler) FetchTransaction(context echo.Context) error {
	limitS := context.QueryParam("limit")
	limit, err := strconv.Atoi(limitS)
	if err != nil || limit == 0 {
		limit = defaultLimit
	}

	cursor := context.QueryParam("cursor")
	ctx := context.Request().Context()

	listAr, nextCursor, err := th.Ts.Fetch(ctx, cursor, int64(limit))
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	context.Response().Header().Set(`X-Cursor`, nextCursor)
	return context.JSON(http.StatusOK, listAr)
}

// GetById godoc
// @Summary      Get transaction by ID
// @Description  Retrieves a transaction by unique ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Success      200  {object}  entities.Transaction
// @Failure      400,404,500  {object}  rest.ResponseError
// @Router       /transaction/{id} [get]
func (th *TransactionHandler) GetById(context echo.Context) error {
	id := context.Param("id")
	ctx := context.Request().Context()

	if id == "" {
		return context.JSON(http.StatusBadRequest, rest.ResponseError{Message: "id must be a valid uuid"})
	}

	art, err := th.Ts.GetById(ctx, id)
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, art)
}

func isRequestValid(i interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Save godoc
// @Summary      Create a new transaction
// @Description  Saves a new transaction to the system
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        transaction  body      entities.Transaction  true  "Transaction object"
// @Success      201          {object}  entities.Transaction
// @Failure      400,409,422,500  {object}  rest.ResponseError
// @Router       /transaction [post]
func (th *TransactionHandler) Save(context echo.Context) (err error) {
	var transaction entities.Transaction
	err = context.Bind(&transaction)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&transaction); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Request().Context()
	err = th.Ts.Save(ctx, &transaction)

	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, transaction)
}

// Delete godoc
// @Summary      Delete transaction
// @Description  Deletes a transaction by unique ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Success      204
// @Failure      400,404,500  {object}  rest.ResponseError
// @Router       /transaction/{id} [delete]
func (th *TransactionHandler) Delete(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	if id == "" {
		return context.JSON(http.StatusBadRequest, rest.ResponseError{Message: "id must be a valid uuid"})
	}

	err := th.Ts.Delete(ctx, id)
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return context.NoContent(http.StatusNoContent)
}

// Exchange godoc
// @Summary      Get exchange transaction
// @Description  Retrieves exchange transaction based on a given transaction ID and country currency
// @Tags         exchange
// @Accept       json
// @Produce      json
// @Param        exchange  body      entities.ExchangeGetParams  true  "Exchange request"
// @Success      200       {object}  entities.ExchangeTransaction
// @Failure      400,404,422,500  {object}  rest.ResponseError
// @Router       /exchange [post]
func (th TransactionHandler) Exchange(context echo.Context) (err error) {
	var exchange entities.ExchangeGetParams
	err = context.Bind(&exchange)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, rest.ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&exchange); !ok {
		return context.JSON(http.StatusBadRequest, rest.ResponseError{Message: err.Error()})
	}

	ctx := context.Request().Context()
	ex, err := th.Es.GetExchange(ctx, &exchange)

	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ex)
}
