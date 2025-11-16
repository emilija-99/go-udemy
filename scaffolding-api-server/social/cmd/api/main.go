package main

import (
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/store"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

const version = "0.0.1"

// exectable for api server
func main() {
	dsn := env.GetString("DB_ADDR",
		"postgres://postgres:root@127.0.0.1:5433/social?sslmode=disable")
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{
			addr:         dsn,
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	m, err := migrate.New(
		"file://../migrate.migrations",
		dsn)

	if err != nil {
		log.Fatalf(" migrate.New failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf(" migration failed: %v", err)
	}

	log.Println("Migrations ran successfully.")

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		log.Fatal(err)
	}

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	err = app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
