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

