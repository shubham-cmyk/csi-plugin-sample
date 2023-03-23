# Build stage
FROM golang:1.19 AS build
# Set working directory
WORKDIR /app
# Copy Go module files
COPY go.mod go.mod
COPY go.sum go.sum
# Download module dependencies
RUN go mod download
# Copy the go source
COPY main.go main.go
COPY pkg/ pkg/
# Build the binary
RUN CGO_ENABLED=0 go build -o myapp main.go


# Final stage
FROM alpine:latest
# Copy the binary from the build stage
COPY --from=build /app/myapp /usr/local/bin/myapp
# Set entrypoint
ENTRYPOINT ["/usr/local/bin/myapp"]
