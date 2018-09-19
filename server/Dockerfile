FROM golang
MAINTAINER asuka

WORKDIR /go/src/github.com/ShikinamiAsuka/ih13/server

RUN go get -u github.com/golang/dep/...


CMD dep ensure && go build && ./server
