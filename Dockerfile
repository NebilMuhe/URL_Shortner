FROM golang:1.24-alpine AS builder
# Set the working directory inside the container
WORKDIR /
ADD . .
# Build the application
RUN go build -o ./server /cmd/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./server"]