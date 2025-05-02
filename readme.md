# VR Exchange Backend

API for managing transactions and currency conversion in dollars for other countries through the API [api.fiscaldata.treasury.gov](https://fiscaldata.treasury.gov/datasets/treasury-reporting-rates-exchange/treasury-reporting-rates-of-exchange)

This repository is used by the front-end corresponding to the application.

#### Requirements

- [x] Creating a Transaction API with Go Lang
- [x] Storing a Purchase Transaction
- [x] Retrieving a Purchase Transaction in a Specific Country's Currency
- [x] Converting currencies with the [api.fiscaldata.treasury.gov](https://fiscaldata.treasury.gov/datasets/treasury-reporting-rates-exchange/treasury-reporting-rates-of-exchange) API
- [x] Clean Architecture
- [x] Containerization with Docker
- [x] Documentation with Swagger
- [x] SOLID Principles

---

### Architecture

The application follows the principles of Clean Architecture and SOLID, with well-defined layers:

- **Domain**: responsible for business rules.

- **Internal/Repository**: provides access to data.

- **Rest**: Layer responsible for receiving and handling HTTP calls (API Calls).

#### Folder structure

```
vr_exchange_backend
 ├── app
 │    ├── .env
 │    └── main.go
 ├── docs
 │    ├── docs.go
 │    ├── swagger.json
 │    └── swagger.yaml
 ├── domain
 │    ├── entities
 │    │    ├── <name.go>
 │    ├── repositories
 │    │    ├── mocks
 │    │    ├── <i_name_repository>.go
 │    ├── services
 │    │    ├── <name_service_impl.go>
 │    │    ├── <name_service_test.go>
 │    │    └── <name_service_impl.go>
 │    └── errors.go
 ├── internal
 │    └── repository
 │         ├── <reposiries_impl>
 │         │    └── <repository_name.go>
 │         └── helper.go
 ├── rest
 │    ├── middleware
 │    │    ├── cors.go
 │    │    └── timeout.go
 │    ├── <feature_name>
 │    │    ├── <feature_name.go>
 │    │    └── <feature_name_test.go>
 │    └── utils.go
 ├── utils
 │    └── utils.go
 ├── .gitignore
 ├── Dockerfile
 ├── docker-compose.yaml
 ├── go.mod
 ├── go.sum
 └── vr_exchange.sql
```

---

## Documentation

All endpoints are documented with **Swagger** and can be accessed at the endpoint:

```bash
/swagger/index.html
```

## Tests

Tests were created to ensure the quality and stability of the application.

To run them, use the command:

```bash
go test ./...
```

### How to Run

1. Clone the repository

```bash
git clone https://github.com/Guilherme-DSGL/vr_exchange_backend.git
```

2. Define a .env file with the default inside /app

```.env
# DB
HOST_DATABASE = "postgres"
PORT_DATABASE = "5432"
USER_DATABASE = "postgres"
PASS_DATABASE = "postgres"
NAME_DATABASE = "vr_exchange"
# Server
SERVER_ADDRESS = ":8080"
CONNECTION_TIMEOUT = 2

# Debugging
DEBUG = "TRUE"
```

3. The application has been containerized and just run the command with docker already installed to build the application

```bash
docker compose --env-file ./app/.env up --build
```

## Developer

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Guilherme-DSGL">
        <img src="https://avatars.githubusercontent.com/u/72310683?s=400&u=9f0ec757e6df46288a0bff579b2648b151319db7&v=4" width="100px;" alt="Guilherme de Souza"/><br>
        <sub>
          <b>Guilherme de Souza</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

Thank you for your attention.
