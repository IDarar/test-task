FROM golang:1.17-alpine3.13 AS builder
WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main .

EXPOSE 6464
CMD [ "/app/main" ]

