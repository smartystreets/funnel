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

	var workers []chan string
	for x := 0; x < 10; x++ {
		output := make(chan string)
		workers = append(workers, output)
		go WORKER(input, output)
	}

	final := make(chan string)
	go MERGER(workers, final)

	for a := range final {
		fmt.Println(a)
	}
}
func LOADER(input chan string) {
	for _, address := range internet.Addresses {
		input <- address
	}
	close(input)
}
func WORKER(input chan string, output chan string) {
	for address := range input {
		output <- internet.Scrape(address)
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
func DRAINER(in, out chan string, done func()) {
	for v := range in {
		out <- v
	}
	done()
}
