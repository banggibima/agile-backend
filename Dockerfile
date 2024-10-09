FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:3.20 AS production

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

COPY --from=builder /app/migration ./migration

EXPOSE 8080

CMD ["./main"]
