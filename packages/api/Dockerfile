# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# This now only copies files from packages/api 
# because that is our build context
COPY . .

# Now 'go build .' works because main.go is in the root of the context
RUN go mod init myapi && go mod tidy
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
