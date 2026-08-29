FROM golang:1.26.5 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /notes_build ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder build/.env .env

COPY --from=builder build/pkg/migrations ./pkg/migrations

COPY --from=builder /notes_build ./notes_server

CMD [ "./notes_server" ]