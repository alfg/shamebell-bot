FROM golang:alpine

RUN apk update && apk add git

ADD . /go/src/github.com/alfg/shamebell-bot

RUN go get -u github.com/golang/dep/cmd/dep

RUN cd /go/src/github.com/alfg/shamebell-bot && dep ensure 

RUN go install github.com/alfg/shamebell-bot/cmd/web
RUN go install github.com/alfg/shamebell-bot/cmd/bot

WORKDIR /go/bin

EXPOSE 8080