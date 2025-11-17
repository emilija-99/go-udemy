package main

import (
	"log"
	"net/http"
)

func (app *application) statusInternlServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Interval Server Error: %s Path: %s, Errors: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "The Server encountered a probled.")
}

func (app *application) statusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request Error: %s Path: %s, Errors: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) statusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found Error: %s Path: %s, Errors: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, err.Error())
}
