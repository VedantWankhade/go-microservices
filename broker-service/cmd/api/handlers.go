package main

import (
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	res := &response{Message: "Hit the broker", Error: false}
	err := app.writeJSON(w, http.StatusOK, res)
	if err != nil {
		app.errorJSON(w, err)
	}
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "auth":

	default:
		app.errorJSON(w, errors.New("Unknown request action"))
	}
}

func (app *Config) Authenticate(w http.ResponseWriter, a AuthPayload) {
}
