package main

import (
	"authService/internal/data"
	"authService/internal/validator"
	"fmt"
	"net/http"
)

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	data.ValidateUsername(v, input.Username)
	data.ValidatePassword(v, input.Password)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByUsername(input.Username, input.Password)
	if err != nil {
		app.logger.PrintError(err, nil)
		return
	}

	fmt.Println(user.Name)
	fmt.Println(len(user.Password.Hash))
}
