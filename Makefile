## Database DSN from .env.example
DB_URL=postgres://Admin:SuperSecretPassword@localhost:5432/SubManagerDB?sslmode=disable

## Apply all migrations
migrate-up:
	migrate -path=./migrations -database "$(DB_URL)" up

## Rollback all migrations
migrate-down:
	migrate -path=./migrations -database "$(DB_URL)" down

## Show current migration version
migrate-version:
	migrate -path=./migrations -database "$(DB_URL)" version

## Build, run docker-compose with rebuild, and run migrations
up:
	docker-compose up --build -d

## Stop and remove containers and volumes
down:
	docker-compose down -v

## View logs
logs:
	docker-compose logs -f

## Run unit tests
test:
	go test ./... -v

