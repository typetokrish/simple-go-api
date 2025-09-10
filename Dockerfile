# Use official Go image
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the Go binary
RUN go build -o server .

# Final lightweight image

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/server /server
CMD ["/server"]
