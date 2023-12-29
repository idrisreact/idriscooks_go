# Install required Go version
FROM golang:1.18-alpine AS build

# Set working directory
WORKDIR /app

# Download dependencies
RUN go get github.com/cosmtrek/air@latest
RUN go mod download

# Copy source code
COPY . . 

# Build application
RUN go build -o /go/bin/app

# Use minimal base image 
FROM alpine:latest

# Copy built binary
COPY --from=build /go/bin/app /app/

# Run app
CMD ["/app/app"]
