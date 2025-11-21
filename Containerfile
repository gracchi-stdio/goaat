# Build stage
FROM docker.io/library/golang:1.24-alpine AS builder

WORKDIR /app

# Install templ for template generation
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ templates
RUN templ generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# Runtime stage
FROM docker.io/library/alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
