## Database DSN from .env.example
DB_URL=postgres://Admin:SuperSecretPassword@localhost:5432/SubManagerDB?sslmode=disable

## Creates a new migration
migrate-create:
	@echo "Creating new migration: $(name)"
	migrate create -seq -ext=.sql -dir=./migrations $(name)

## Apply all migrations
migrate-up:
	migrate -path=./migrations -database "$(DB_URL)" up

## Rollback all migrations
migrate-down:
	migrate -path=./migrations -database "$(DB_URL)" down

## View the current version of migrations
migrate-version:
	migrate -path=./migrations -database "$(DB_URL)" version
