# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/order/main.go

# Final stage
FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /app/main .

USER nonroot
EXPOSE 8081
ENTRYPOINT ["/app/main"]
