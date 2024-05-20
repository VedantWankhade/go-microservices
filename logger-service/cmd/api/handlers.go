package main

import (
	"net/http"

	"github.com/vedantwankhade/go-microservices/logger-service/data"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	_ = app.readJSON(w, r, &requestPayload)
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	res := response{
		Error:   false,
		Message: "Loged",
	}
	app.writeJSON(w, http.StatusAccepted, res)
}
