# Build stage
FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Final application image
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 8080
CMD ["/app/main"]