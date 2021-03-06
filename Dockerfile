# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Agus Wibawantara"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

WORKDIR /app/cmd

# Download all the dependencies
RUN go mod tidy
RUN go get -d -v ./...

WORKDIR /app/internal

# Run tests
RUN go mod tidy
RUN go test -v ./...

WORKDIR /app/cmd

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]
