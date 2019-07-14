FROM golang:alpine

ENV GO111MODULE=on

RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base

WORKDIR /go/src/go-nfe-repeater

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/go-nfe-repeater

EXPOSE 8080
CMD ["/go/bin/go-nfe-repeater"]