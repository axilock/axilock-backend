FROM golang:1.24.0-alpine3.20 AS builder

WORKDIR /app

ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

COPY . .

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -ldflags "-s -w" -o app main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

RUN touch .env

COPY axilock.pem /app/

EXPOSE 8080

CMD [ "/app/app" ]
