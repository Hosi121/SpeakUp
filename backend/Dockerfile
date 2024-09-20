FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Attempt to build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["./main"]