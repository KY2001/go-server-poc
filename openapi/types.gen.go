// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

// Error defines model for Error.
type Error struct {
	// Message Human-readable error message.
	Message string `json:"message"`

	// StatusCode HTTP status code.
	StatusCode int `json:"statusCode"`

	// Title Short error code or identifier.
	Title *string `json:"title,omitempty"`
}

// GetHealthResponse defines model for GetHealthResponse.
type GetHealthResponse struct {
	// Status Server status.
	Status string `json:"status"`
}
