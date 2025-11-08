FROM golang:1.25.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /resto

FROM alpine:3.21

ARG ENVIRONMENT

WORKDIR /app

COPY --from=builder /resto .

COPY db/sql_migrations ./db/sql_migrations

EXPOSE 8080

CMD ["./resto"]
