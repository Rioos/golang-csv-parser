FROM golang:1.11

RUN mkdir -p /go/src/golang-csv-parser
WORKDIR /go/src/golang-csv-parser

ADD . /go/src/golang-csv-parser

RUN go get -v