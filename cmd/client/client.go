package main

import (
	"github.com/denisskin/word-of-wisdom"
	"log"
	"runtime"
	"sync"
)

func main() {
	// test client
	var wg sync.WaitGroup
	routines := runtime.NumCPU()
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := wow.NewClient("127.0.0.1:8080")
			for {
				if msg, err := client.Get(); err != nil {
					log.Printf("ERROR: %v", err)
				} else {
					log.Printf("Wisdom: %s", msg)
				}
			}
		}()
	}
	wg.Wait()
}
