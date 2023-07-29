# Use a Go base image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o genhttpserv

RUN chmod +x genhttpserv


EXPOSE 8888

ENTRYPOINT ["./genhttpserv"]