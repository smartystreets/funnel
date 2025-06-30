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

	input := make(chan string)
	go LOADER(input)

	var workerOutputs []chan string
	for range 10 {
		output := make(chan string)
		workerOutputs = append(workerOutputs, output)
		go WORKER(input, output)
	}

	final := make(chan string)
	go MERGER(workerOutputs, final)

	for value := range final {
		fmt.Println(value)
	}
}
func LOADER(input chan string) {
	for _, address := range internet.Addresses {
		input <- address
	}
	close(input)
}
func WORKER(input, output chan string) {
	for address := range input {
		output <- fmt.Sprintf("%s: %s", address, internet.ScrapeTitle(address))
	}
	close(output)
}
func MERGER(workerOutputs []chan string, final chan string) {
	var waiter sync.WaitGroup
	for _, out := range workerOutputs {
		waiter.Add(1)
		go DRAINER(out, final, waiter.Done)
	}
	waiter.Wait()
	close(final)
}
func DRAINER(input, output chan string, done func()) {
	for v := range input {
		output <- v
	}
	done()
}

/*
The draining and closing of every goroutine is accounted for.
This approach does NOT preserve input order.
The number of workers is configurable (scales with more input).
This is the same approach used by the `gitreview` process to execute `git fetch` on all your repos.
It's used in our CLI as well as several other places in actual production code.
10. Another approach to safe concurrency is Go channels, which are FIFO queue structures that are actually safe to
    send to/receive from across multiple goroutines
11. Go Slogan: "Do not communicate by sharing memory; instead, share memory by communicating."
12. Attempting to receive on an empty channel blocks the current goroutine until a value is sent or the channel is
    closed (which results in the 'zero value' being received).
13. Attempting to send on an 'unbuffered' channel blocks until the value is received or the channel is closed (which
    then results in a panic)
14. Concurrency != Parallelism
17. A function that loads a channel should almost always close it when finished (ie. `defer close(ch)`)
*/
