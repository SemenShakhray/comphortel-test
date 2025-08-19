FROM golang:1.23.4-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gettext musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY utils ./utils
COPY migrations ./migrations
COPY .env .env

RUN CGO_ENABLED=0 go build -o /app/comphortel-test ./cmd/main.go

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /app/comphortel-test .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

CMD ["/app/comphortel-test"]
