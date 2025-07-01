package main

import (
	"atsinfoService/internal/data"
	"atsinfoService/internal/validator"
	"net/http"
)

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

func (app *application) addGatewayHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"name"`
		Number     string `json:"number"`
		Digit      int    `json:"digit"`
		Trunk      int    `json:"trunk"`
		TrunkGroup int    `json:"trunkgroup"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badReqestResponse(w, r, err)
		return
	}

	gateway := &data.Gateway{
		Name:       input.Name,
		Number:     input.Number,
		Digit:      input.Digit,
		Trunk:      input.Trunk,
		TrunkGroup: input.TrunkGroup,
	}

	v := validator.New()
	if data.ValidateGateway(v, *gateway); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Gateways.Insert(gateway)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	header := make(http.Header)
	header.Set("Location", "/v1/gateways")

	err = app.writeJSON(w, http.StatusCreated, envelope{"gateway": gateway}, header)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
