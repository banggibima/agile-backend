# Use the official Golang image as the builder
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:3.20 AS production

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the environment file
COPY --from=builder /app/.env .

# Copy the migration folder
COPY --from=builder /app/migration ./migration

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
