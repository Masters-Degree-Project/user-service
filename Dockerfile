FROM golang:1.23.2-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables and build the API server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o cmd/main.go .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main .
COPY .env .

EXPOSE 8080

CMD ["./main"] 