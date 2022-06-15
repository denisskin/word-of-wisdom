FROM golang:1.18.3-alpine

ADD . /sources
WORKDIR /sources

RUN go build -o /usr/local/bin/wow-server ./cmd/server/...

EXPOSE 8080

# Start wow-server
#   - set limit: 10 req/sec;
#   - set min PoW difficulty: 20000 hashes/sec
CMD wow-server -r 10  -d 20000
