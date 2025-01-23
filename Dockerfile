# Stage 1: Build the Go binary
FROM golang:1.23.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

# Stage 2: Create the production image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server .
ENV ENV production
EXPOSE 8080
CMD ["./server"]
