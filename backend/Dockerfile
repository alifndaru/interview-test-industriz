# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git and build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy migration files
COPY --from=builder /app/database/migrations ./database/migrations
COPY --from=builder /app/database/seed ./database/seed

# Copy .env file (optional, can use environment variables instead)
COPY --from=builder /app/.env* ./

# Expose port
EXPOSE 3030

# Command to run the executable
CMD ["./main"]