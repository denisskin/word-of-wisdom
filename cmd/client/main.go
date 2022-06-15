package main

import (
	"flag"
	"github.com/denisskin/word-of-wisdom/wow_client"
	"log"
)

var (
	address   = flag.String("a", ":8080", "Address of TCP-server")
	nRequests = flag.Int("n", 10000, "Count of requests")
)

func main() {
	flag.Parse()

	client := wow_client.New(*address)
	for i := 0; i < *nRequests; i++ {
		if msg, err := client.Get(); err != nil {
			log.Printf("ERROR: %v", err)
		} else {
			log.Printf("Wisdom: %s", msg)
		}
	}
}
