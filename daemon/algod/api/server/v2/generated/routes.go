// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
