FROM golang:1.25-alpine

WORKDIR /app

# Install 'Air' for hot reloading
RUN go install github.com/air-verse/air@latest

# Copy modules first for faster caching
COPY go.mod go.sum* ./
RUN go mod download

# We don't COPY . . because we will mount it via volume
EXPOSE 8080

# Run Air to watch and compile code for Linux
CMD ["air"]
