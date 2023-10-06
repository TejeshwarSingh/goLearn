# Use the official Golang image with Alpine as the base image for building
FROM golang:1.19-alpine AS build-env

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

# Initialize a new module or migrate old ones
RUN go mod init app || echo "Module already initialized"

# Fetch and tidy up the dependencies
RUN go mod tidy

# Build the Go app
#RUN go build -o main
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a smaller, Alpine image for running the app
FROM alpine:latest

WORKDIR /app

# Add ca-certificates in case you need them
RUN apk --no-cache add ca-certificates

# Copy the Go binary from the build stage
COPY --from=build-env /app/main /app/

# Command to run when starting the container
CMD ["/app//main"]
