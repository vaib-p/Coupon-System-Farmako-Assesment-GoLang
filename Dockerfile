# Build stage
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Install git (required if your dependencies are from git)
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o coupon-api ./main.go  # adjust main.go path if needed

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/coupon-api .

# Copy any required files (e.g., configs) here if needed
# Copy .env file
COPY .env . 
# Expose the port your app runs on
EXPOSE 8080

# Command to run the app
CMD ["./coupon-api"]
