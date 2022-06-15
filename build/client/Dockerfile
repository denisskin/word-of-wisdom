FROM golang:1.18.3-alpine

ADD . /sources
WORKDIR /sources

RUN go build -o /usr/local/bin/wow-client ./cmd/client/...

CMD wow-client >/var/log/wow-client.log
