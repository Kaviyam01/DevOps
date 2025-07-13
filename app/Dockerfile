# Stage 1: Build the Go app
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o simpletimeservice main.go

# Stage 2: Create minimal container
FROM alpine:latest

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser
COPY --from=builder /app/simpletimeservice .

EXPOSE 8080
CMD ["./simpletimeservice"]