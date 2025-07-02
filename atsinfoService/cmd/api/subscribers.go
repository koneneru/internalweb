package main

import (
	"fmt"
	"net/http"
)

func (app *application) listSubscribersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id   string
		Name string
	}

	qs := r.URL.Query()

	input.Id = app.readString(qs, "id", "")
	input.Name = app.readString(qs, "name", "")
	fmt.Println(input.Name)

	subscribers, err := app.models.Subscribers.GetAll(input.Id, input.Name)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"subscribers": subscribers, "total": len(subscribers)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
