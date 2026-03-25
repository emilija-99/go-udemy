package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// runs everytime the go program is initialized
func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// send a response and data
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // max 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // show error if there are unknown fields
	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	// error in same format
	type envelop struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelop{Error: message})
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return writeJSON(w, status, &envelope{Data: data})

}
