package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	res := &response{Message: "Hit the broker", Error: false}
	err := app.writeJSON(w, http.StatusOK, res)
	if err != nil {
		app.errorJSON(w, err)
	}
}
