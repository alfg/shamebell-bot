FROM golang:alpine

ADD . /go/src/github.com/alfg/shamebell-bot

RUN go install github.com/alfg/shamebell-bot/cmd/web
RUN go install github.com/alfg/shamebell-bot/cmd/bot

WORKDIR /go/bin

EXPOSE 8080