# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o file-storage .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/file-storage .
COPY templates/ ./templates/
EXPOSE 8080
COPY .env ./
CMD ["./file-storage"]
