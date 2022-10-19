package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mdwhatcott/funnel"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	var waiter sync.WaitGroup
	for _, address := range funnel.Addresses {
		waiter.Add(1)
		go printTitle(waiter.Done, address)
	}
	waiter.Wait()
}

func printTitle(done func(), address string) {
	defer done()
	fmt.Println(funnel.ScrapeTitle(address))
}
