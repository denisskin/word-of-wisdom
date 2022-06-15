FROM golang:1.18.3-alpine

ADD . /sources
WORKDIR /sources

RUN go build -o /usr/local/bin/wow-server ./cmd/server/...

EXPOSE 8080

CMD wow-server -r=10 >/var/log/wow-server.log
