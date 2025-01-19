FROM golang:1.21-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main .
COPY .env .

EXPOSE 8080

CMD ["./main"] 