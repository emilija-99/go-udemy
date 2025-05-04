package main

import (
	"log"
	"social/internal/env"
)

// exectable for api server
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	err := app.run(mux)

	if err != nil {
		log.Fatal(err)
	}
}
