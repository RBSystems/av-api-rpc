FROM golang:1.7.1-alpine

RUN apk update && apk upgrade && apk add git

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/av-api-rpc

WORKDIR /go/src/github.com/byuoitav/av-api-rpc
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/av-api-rpc"]

EXPOSE 8100
