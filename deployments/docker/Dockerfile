FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o redira ./cmd/server


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/redira .

EXPOSE 8080

CMD ["./redira"]