FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o kvcache ./cmd/kvcache

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kvcache .

EXPOSE 7171

CMD ["./kvcache"]
