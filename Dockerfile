# Build stage
FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service main.go

# Final stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/auth-service .

EXPOSE 8080

ENTRYPOINT ["/app/auth-service"] 