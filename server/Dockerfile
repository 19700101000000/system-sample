FROM golang
MAINTAINER asuka

WORKDIR /go/src/github.com/ShikinamiAsuka/system-sample/server

RUN go get -u github.com/golang/dep/...


CMD dep ensure && go build && ./server
