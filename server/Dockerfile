FROM golang
MAINTAINER asuka

WORKDIR /go/src/github.com/19700101000000/system-sample/server

RUN go get -u github.com/golang/dep/... && \
	dep init

CMD dep ensure && go build && ./server
