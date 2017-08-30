FROM golang:alpine

RUN apk update && apk add git yarn

ADD . /go/src/github.com/alfg/shamebell-bot
# ADD ./assets /go/bin/assets

RUN go get -u github.com/golang/dep/cmd/dep

RUN cd /go/src/github.com/alfg/shamebell-bot && dep ensure 

RUN cd /go/src/github.com/alfg/shamebell-bot/static && yarn

RUN go install github.com/alfg/shamebell-bot/cmd/web
RUN go install github.com/alfg/shamebell-bot/cmd/bot

WORKDIR /go/src/github.com/alfg/shamebell-bot

EXPOSE 4000