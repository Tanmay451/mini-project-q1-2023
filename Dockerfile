# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the project dependencies
RUN go mod download

# Copy the rest of the project files to the working directory
COPY . .

# Build the Go application
RUN go build -o chatapp

# Set the command to run the binary executable
CMD ["./chatapp"]