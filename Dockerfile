FROM golang:1.21-alpine AS builder

WORKDIR /app

# Add git and necessary build tools
RUN apk add --no-cache git

# Copy module files first
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"] 