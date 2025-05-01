package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/Guilherme-DSGL/purchase_transaction_backend/docs"
	exchangeServ "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/exchange"
	transactionServ "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/transaction"
	exchangeRepo "github.com/Guilherme-DSGL/purchase_transaction_backend/internal/repository/exchange_api"
	postgressRepo "github.com/Guilherme-DSGL/purchase_transaction_backend/internal/repository/postgress"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/rest/middleware"
	transactionRest "github.com/Guilherme-DSGL/purchase_transaction_backend/rest/transaction"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title VR Exchange API
// @version 1.0
// @description API to persist transactions and convert the value to currencies of other countries
// @host localhost:8080
// @BasePath /api/v1
func main() {
	dbConn := initDb()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("error when closing DB", err)
		}
	}()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.CORS)
	connectionTimeout := os.Getenv("CONNECTION_TIMEOUT")

	timeout, err := strconv.Atoi(connectionTimeout)
	if err != nil {
		log.Fatal("timout failed to parse as int")
		return
	}
	timeoutDuration := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextTimeout(timeoutDuration))

	apiGroup := e.Group("/api/v1")
	initDependencies(dbConn, apiGroup)

	address := os.Getenv("SERVER_ADDRESS")
	log.Fatal(e.Start(address))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env file not found")
	}
}

func initDb() *sql.DB {
	dbHost := os.Getenv("HOST_DATABASE")
	dbPort := os.Getenv("PORT_DATABASE")
	dbUser := os.Getenv("USER_DATABASE")
	dbPass := os.Getenv("PASS_DATABASE")
	dbName := os.Getenv("NAME_DATABASE")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open(`postgres`, connection)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}
	return dbConn
}

func initDependencies(dbConn *sql.DB, e *echo.Group) {

	httpClient := &http.Client{Timeout: 10 * time.Second}
	//  Repositories
	transactioRepo := postgressRepo.NewTransactionRepository(dbConn)
	exchangeRepo := exchangeRepo.NewExchangeRepository(httpClient)
	// Services
	tServ := transactionServ.NewTransactionService(transactioRepo)
	eServ := exchangeServ.NewExchangeService(exchangeRepo, transactioRepo)
	transactionRest.NewTransactionHandler(e, tServ, eServ)
}
