# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o cdn-simulator ./main.go

# Stage 2: Run the application in a lightweight container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Create directories for file storage
RUN mkdir -p /root/server1_files /root/server2_files

# Copy the Go binary from the builder
COPY --from=builder /app/cdn-simulator .

# Expose the ports used by the CDN and load balancer
EXPOSE 8080 8081 8082

# Run the binary
CMD ["./cdn-simulator"]
