package main

import (
	"net/http"
	"time"
)

func (app *application) listCallsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Direction       string
		FromDate        time.Time
		ToDate          time.Time
		CallPhone       string
		SubscriberPhone string
		Gateway         string
		Intercity       string
	}

	qs := r.URL.Query()

	input.Direction = app.readString(qs, "direction", "")
	input.FromDate = app.readDate(qs, "fromdate", time.Now())
	input.ToDate = app.readDate(qs, "todate", time.Now())
	input.CallPhone = app.readString(qs, "callphone", "")
	input.SubscriberPhone = app.readString(qs, "subscriber", "")
	input.Gateway = app.readString(qs, "gateway", "")
	input.Intercity = app.readString(qs, "intercity", "false")

	input.FromDate = time.Date(input.FromDate.Year(), input.FromDate.Month(), input.FromDate.Day(), 0, 0, 0, 0, input.FromDate.Location())
	input.ToDate = time.Date(input.ToDate.Year(), input.ToDate.Month(), input.ToDate.Day(), 23, 59, 59, 0, input.ToDate.Location())

	calls, err := app.models.Calls.GetAll(input.Direction, input.SubscriberPhone, input.Gateway, input.CallPhone, input.Intercity, input.FromDate, input.ToDate)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"calls": calls}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
