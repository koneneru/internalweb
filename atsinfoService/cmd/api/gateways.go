package main

import "net/http"

func (app *application) listGatewaysHandler(w http.ResponseWriter, r *http.Request) {
	gateways, err := app.models.Gateways.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"gateways": gateways, "total": len(gateways)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
