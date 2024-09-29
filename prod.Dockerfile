# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest

WORKDIR /usr/local/bin

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server /app/.env ./
EXPOSE 8080

CMD ["server"]
