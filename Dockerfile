# Build stage
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Copy module files and download dependencies (if any)
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY *.go ./

# Build the Go app as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redis-from-scratch .

# Final minimal stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

# Copy the binary to a system path
COPY --from=builder /app/redis-from-scratch /usr/local/bin/redis-from-scratch

# Set working directory to /data (where database.aof will be created and persisted)
WORKDIR /data

# Expose port 6379 to the host
EXPOSE 6379

# Define a volume mount point for persistence
VOLUME ["/data"]

# Run the binary
CMD ["redis-from-scratch"]
