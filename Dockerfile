FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Build the executables
RUN go build -o /app/bin/api ./cmd/api/main.go
RUN go build -o /app/bin/cron ./cmd/cron/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/api /app/bin/api
COPY --from=builder /app/bin/cron /app/bin/cron

# Set the default command (will be overridden by fly.toml for each app)
CMD ["/app/bin/api"]

