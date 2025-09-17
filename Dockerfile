FROM golang:1.25.1-alpine AS builder

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o main ./cmd/server

# ---- Final image ----
FROM alpine:latest

WORKDIR /app

# Copy built binary and .env
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]
