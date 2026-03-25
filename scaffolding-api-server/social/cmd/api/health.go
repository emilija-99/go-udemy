package main

import (
	"log"
	"net/http"
)

/*
  - type ResponseWriter interface {
    Header() http.Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
    }

    w http.ReponseWritet is the OUTPUT of the channel back to the client
    r *http.Request is INPUT (url, headers, body, method) from the client

  - http localhost:3000/v1/health

# Becomes a TCP request that GO's HTTP server receives

GO math the route:
r.Get("/v1/health", app.healthCheckHandler)

Request object:
// read is a pointer handlers can modify mutable data inside context
r.Method        // "GET"
r.URL.Path      // "/v1/health"
r.Header        // HTTP headers
r.Body          // request body stream (for POST, PUT, etc.)
r.Context()     // request-scoped cancellation & deadline
*/
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("w: %+v", w)
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	log.Printf("w: %+v", data)
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.statusInternlServerError(w, r, err)
	}
}
