# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY *.go ./
COPY config.json ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o rest-server .

# Final stage - using alpine (thin OS image ~5MB)
FROM alpine:latest

# Install only CA certificates (minimal dependencies)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary and config from builder
COPY --from=builder /app/rest-server .
COPY --from=builder /app/config.json .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./rest-server"]
