package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	type ServiceInfo struct {
		Status      string
		Environment string
	}

	data := map[string]ServiceInfo{
		"frontend-server": {
			Status:      "available",
			Environment: app.config.env,
		},
	}

	resp, err := http.Get(app.config.api + "/healthcheck")
	if err != nil {
		app.logger.PrintError(err, nil)
		data["backend-server"] = ServiceInfo{
			Status: "unavailable",
		}
	} else {
		defer resp.Body.Close()

		var target ServiceInfo
		err = json.NewDecoder(resp.Body).Decode(&target)
		if err != nil {
			app.logger.PrintError(err, nil)
		}

		data["backend-server"] = target
	}

	tmpl, err := template.ParseFiles("./public/html/healthcheck.html")
	if err != nil {
		app.logger.PrintError(err, nil)
		return
	}

	tmpl.ExecuteTemplate(w, "healthcheck", data)
}
