FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o submanager .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/submanager .
COPY .env .env  

CMD ["./submanager"]