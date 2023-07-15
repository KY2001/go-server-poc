//go:generate bin/oapi-codegen -generate types -package openapi -o openapi/types.gen.go openapi/openapi.yaml
//go:generate bin/oapi-codegen -generate server -package openapi -o openapi/server.gen.go openapi/openapi.yaml

package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddle "github.com/labstack/echo/v4/middleware"

	"github.com/KY2001/go-server-poc/config"
	"github.com/KY2001/go-server-poc/handler"
	"github.com/KY2001/go-server-poc/openapi"
)

func main() {
	conf := config.NewConfig()

	e := echo.New()

	e.Use(echoMiddle.Recover())
	e.Use(echoMiddle.CORS())
	e.Use(echoMiddle.Logger())
	e.Use(echoMiddle.TimeoutWithConfig(echoMiddle.TimeoutConfig{
		Timeout: conf.Server.Timeout,
	}))

	api := e.Group("")
	handlers := handler.NewHandlers()
	openapi.RegisterHandlers(api, handlers)

	// Start server
	address := getAddress(conf.Server.Port)
	err := e.Start(address)
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func getAddress(portNumber int) string {
	return fmt.Sprintf(":%d", portNumber)
}
