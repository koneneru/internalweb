package main

import "net/http"

func (app *application) listPhonebookHandler(w http.ResponseWriter, r *http.Request) {
	records, err := app.models.PhonebookRecords.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"phonebookRecords": records}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
