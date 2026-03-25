DB_URL=postgres://postgres:root@localhost:5433/social?sslmode=disable
MIGRATIONS_PATH=./cmd/migrate.migrations

migrate-up:
	migrate -path ./cmd/migrate.migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force 1

migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version
