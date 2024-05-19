package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	err = decoder.Decode(io.Discard)
	if err != io.EOF {
		return errors.New("request body should only have one json object")
	}
	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	payload := &response{
		Message: err.Error(),
		Error:   true,
	}
	return app.writeJSON(w, statusCode, payload)
}
