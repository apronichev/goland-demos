# Use the official Golang image to build the app
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go application code
COPY . .

# Build the Go application
RUN go build -o hello-world .

# Use a minimal image to run the application
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/hello-world /hello-world

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["/hello-world"]
