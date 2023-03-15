package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmaabanto/go-user-ms/internal/ports"
	"github.com/juanmaabanto/go-user-ms/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	router := echo.New()
	ctx := context.Background()
	application, err := service.NewApplication(ctx, os.Getenv("MONGODB_NAME"), os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Panic("Error loading app")
	}

	Handler(ports.NewHttpServer(application), router)
	router.Logger.Fatal(router.Start(":3000"))
}

type ServerInterface interface {
	CreateUser(c echo.Context) error
	// ReadUser(c echo.Context) error
	// UpdateUser(c echo.Context) error
	// DeleteUser(c echo.Context) error
}

func Handler(si ServerInterface, router *echo.Echo) {
	if router == nil {
		router = echo.New()
	}

	api := router.Group("/api/v1")

	//Swagger
	router.GET("/*", echoSwagger.WrapHandler)

	//user
	// api.DELETE("/users/:id", si.DeleteUser)
	// api.GET("/users/:id", si.CreateUser)
	// api.PATCH("/users", si.UpdateUser)
	api.POST("/users", si.CreateUser)
}
