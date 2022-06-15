FROM golang:1.18.3-alpine

ADD . /sources
WORKDIR /sources

RUN go build -o /usr/local/bin/wow-client ./cmd/client/...

# Start test client, send 10K requests
CMD wow-client -a 'host.docker.internal:8080' -n 10000
