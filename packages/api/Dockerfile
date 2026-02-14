FROM golang:1.25-alpine

WORKDIR /app
RUN apk add --no-cache git build-base

# Install 'Air' for hot reloading
RUN go install github.com/air-verse/air@latest

# Install 'sqlc' for code generation
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Copy modules first for faster caching
COPY go.mod go.sum* ./
RUN go mod download

# We don't COPY . . because we will mount it via volume
EXPOSE 8080
CMD ["sh", "-c", "go mod tidy && air"]
