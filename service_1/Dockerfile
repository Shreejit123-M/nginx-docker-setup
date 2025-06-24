# Use the official Go image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy source code into container
COPY . .

# Initialize Go module (inside Docker)
RUN go mod init service1

# Build the Go app
RUN go build -o main .

# Expose the app's port
EXPOSE 8001

# Run the app
CMD ["./main"]
