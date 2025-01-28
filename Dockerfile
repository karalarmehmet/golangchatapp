# Stage 1: Build
FROM golang:1.22.1 AS builder

# Set the working directory
WORKDIR /golangchatapp

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o ./tmp/chat ./cmd/chat

# Stage 2: Run
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /golangchatapp/tmp/chat ./tmp/chat

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./tmp/chat"]
