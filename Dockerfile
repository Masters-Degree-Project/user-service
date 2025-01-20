FROM golang:1.23.2-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables and build the API server
RUN GOOS=linux GOARCH=$(go env GOARCH) go build -o main ./cmd

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main .
COPY .env .

CMD ["./main"] 