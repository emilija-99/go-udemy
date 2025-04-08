package main

import (
	"net/http"
)

type server struct {
	addr string
}

type api struct {
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list.."))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user.."))
}

/*
Implementing interface in GO must have same
method name like interface is defined.
*/
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("welcome from the server sd"))
	switch r.Method {
	// use predefined http methods
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
			// curl http://localhost:8080/users
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	default:
		w.Write([]byte("404 page not found"))
	}
}

func (api *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// use predefined http methods
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
			// curl http://localhost:8080/users
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	default:
		w.Write([]byte("404 page not found"))
	}
}

func main() {
	// initialize the Server MUX  - router will direct incoming req to the appropriate handlers
	mux := http.NewServeMux()

	// implementing the httpHandler interface - listen on port 8080
	api := &api{addr: ":8080"}

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	// optional, path with method
	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	/*
		When an HTTP req handling with No maching handler,
		default behaviour:
		- 404 - req comes and doesn't match any reqistered patter
		response from the mux is 404 "Not Found";

		or

		- 405 - path maches but the method doesn't the mux will
		return 405 "Method Not Allowed";

	*/
	server.ListenAndServe()
}

/*
	2025/04/08 17:35:02 listen tcp :8080: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
	exit status 1

	explaied:
	This error message indicates that your application is trying to bind to port 8080, but the port
	is already being used by another process on your system.const

	net -ano | findstr: 8080
	taskkill /PID <PID> /F


	How to run this program:
	- first run in command line:
	go run main.go - bind address and write message
	- open new terminal tab and run:
	curl http://localhost:8080 - listen what happen and read message

*/
