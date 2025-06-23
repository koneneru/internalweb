package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// envelope := envelope{
	// 	"status":      "available",
	// 	"environment": app.config.env,
	// }

	resp, err := http.Get(app.config.api + "/healthcheck")
	if err != nil {
		app.logger.PrintError(err, nil)
		return
	}

	defer resp.Body.Close()

	var target struct{}
	json.NewDecoder(resp.Body).Decode(&target)

	tmpl, err := template.ParseFiles("./public/html/healthcheck.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tmpl.ExecuteTemplate(w, "healthcheck", target)
}
