package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mdwhatcott/funnel/training/stuff"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	var waiter sync.WaitGroup
	for _, address := range stuff.Addresses {
		waiter.Add(1)
		go printTitle(waiter.Done, address)
	}
	waiter.Wait()
}

func printTitle(done func(), address string) {
	defer done()
	fmt.Println(stuff.ScrapeTitle(address))
}
