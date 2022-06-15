package main

import (
	"flag"
	"github.com/denisskin/word-of-wisdom"
)

var (
	port       = flag.Uint("p", 8080, "TCP port of incoming connections")
	difficulty = flag.Uint64("d", 20000, "Minimal PoW Difficulty. (Number of hashes per request)")
	limit      = flag.Uint64("r", 100, "Income Requests Limit. (Allowed number of requests per second)")
)

func main() {
	flag.Parse()

	wow.StartServer(*port, *difficulty, *limit)
}
