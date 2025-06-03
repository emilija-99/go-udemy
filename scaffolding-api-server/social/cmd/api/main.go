package main

import (
	"log"
	"social/internal/env"
	"social/internal/store"
)

// exectable for api server
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	err := app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
