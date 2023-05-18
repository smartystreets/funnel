package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/smartystreets/funnel/training/internet"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	var waiter sync.WaitGroup
	for a, address := range internet.Addresses {
		waiter.Add(1)
		go printTitle(waiter.Done, a, address)
	}
	waiter.Wait()
}

func printTitle(done func(), a int, address string) {
	defer done()
	fmt.Println(a, address, internet.ScrapeTitle(address))
}

/*
What happened to the input order?
What are the implications of this approach as the number of inputs increases?
*/
