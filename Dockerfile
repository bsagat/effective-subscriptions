FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

COPY docs/swagger.json docs/swagger.json

RUN go mod download

RUN go build -o submanager cmd/main.go

CMD ["./submanager"]

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/submanager .
COPY .env .env
COPY docs/swagger.json docs/swagger.json

CMD ["./submanager"]
