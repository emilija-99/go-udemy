package main

import (
	"net/http"
)

type api struct {
	addr string
}

func main() {

	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("GET /users", api.createUsersHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
