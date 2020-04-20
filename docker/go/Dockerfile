ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
ADD ./src /src
WORKDIR /src/cmd

RUN pwd
RUN ls -alh
RUN go build -o app

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /src/cmd/app .

EXPOSE 8080

ENTRYPOINT ["./app"]