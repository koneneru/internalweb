package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc("GET", "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc("GET", "/phonebook", app.listPhonebookHandler)

	return router
}
