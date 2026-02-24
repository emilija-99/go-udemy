# ---- Build stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

# ---- Run stage ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/cmd/migrate.migrations ./cmd/migrate.migrations

EXPOSE 3000

CMD ["./main"]
