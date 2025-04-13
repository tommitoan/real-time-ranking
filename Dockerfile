# 1. Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (needed for go get) and build tools
RUN apk add --no-cache git

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the binary
RUN go build -o real-time-ranking ./cmd/main.go

# 2. Run stage (small image)
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/real-time-ranking .

COPY --from=builder /app/config.yaml .
COPY --from=builder /app/docs ./docs

# Expose the port your app listens on (update if needed)
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["./real-time-ranking"]
