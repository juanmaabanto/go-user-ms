# go-user-ms
user microservice for go

## Installation

  ```sh
// download dependencies
go mod tidy

// Ejecutar
go run cmd/main.go

// run tests
go test ./internal/app/... -v

// Coverage
go test ./internal/app/... -cover
```

Ir a http://localhost:3000 para ver swagger specification

## Swagger

![swagger](https://github.com/juanmaabanto/go-user-ms/blob/main/docs/swagger.png)
