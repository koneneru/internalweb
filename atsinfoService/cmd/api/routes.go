package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/calls", app.listCallsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/gateways", app.listGatewaysHandler)
	router.HandlerFunc(http.MethodPost, "/v1/gateways", app.addGatewayHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/gateways/:id", app.editGatewayHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/gateways/:id", app.deleteGatewayHandler)

	return router
}
