package rest

import (
	"net/http"
	"strconv"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	ts "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/transaction"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/rest"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

type TransactionHandler struct {
	Ts ts.ITransactionService
}

const defaultNum = 10

func NewTransactionHandler(e *echo.Echo, ts ts.ITransactionService) {
	handler := &TransactionHandler{
		Ts: ts,
	}
	e.GET("/transaction", handler.FetchTransaction)
	e.POST("/transaction", handler.Save)
	e.GET("/transaction/:id", handler.GetById)
	e.DELETE("/transaction/:id", handler.Delete)
}

func (th *TransactionHandler) FetchTransaction(context echo.Context) error {
	numS := context.QueryParam("num")
	num, err := strconv.Atoi(numS)
	if err != nil || num == 0 {
		num = defaultNum
	}

	cursor := context.QueryParam("cursor")
	ctx := context.Request().Context()

	listAr, nextCursor, err := th.Ts.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	context.Response().Header().Set(`X-Cursor`, nextCursor)
	return context.JSON(http.StatusOK, listAr)
}

func (th *TransactionHandler) GetById(context echo.Context) error {
	id := context.Param("id")
	ctx := context.Request().Context()

	art, err := th.Ts.GetById(ctx, id)
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, art)
}

func isRequestValid(transaction *entities.Transaction) (bool, error) {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return false, err
	}
	return true, nil
}

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

func (th *TransactionHandler) Delete(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	err := th.Ts.Delete(ctx, id)
	if err != nil {
		return context.JSON(rest.GetStatusCode(err), rest.ResponseError{Message: err.Error()})
	}

	return context.NoContent(http.StatusNoContent)
}
