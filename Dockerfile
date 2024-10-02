# Use the official Golang image as a build stage
FROM golang:1.23-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app

# Create a smaller image to run the built application
FROM alpine:latest

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Go app will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
