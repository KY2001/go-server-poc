package middleware

import (
	"context"
	"log"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"

	"github.com/KY2001/go-server-poc/infrastructure/firebase"
	"github.com/KY2001/go-server-poc/openapi"
)

func RequestValidator() echo.MiddlewareFunc {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalf("RequestValidator: Failed to load swagger: %v\n", err)
	}

	// Skip validating the servers array in the swagger spec
	// See: https://github.com/deepmap/oapi-codegen/issues/882
	swagger.Servers = nil

	return middleware.OapiRequestValidatorWithOptions(
		swagger,
		NewOapiRequestValidatorOptions(),
	)
}

func NewOapiRequestValidatorOptions() *middleware.Options {
	return &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: authenticationFunc(),
		},
	}
}

// Is called only when "security: bearerAuth: []"" is declared in openapi.yaml
func authenticationFunc() func(c context.Context, input *openapi3filter.AuthenticationInput) error {
	return func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		authClient := firebase.GetClient()

		idToken := getTokenFromRequestHeader(input)
		_, err := authClient.VerifyIDToken(c, idToken)
		if err != nil {
			return err
		}

		return nil
	}
}

func getTokenFromRequestHeader(input *openapi3filter.AuthenticationInput) string {
	return input.RequestValidationInput.Request.Header.Get(echo.HeaderAuthorization)
}
