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
	defer close(input)
	for _, address := range funnel.Addresses {
		input <- address
	}
}
func WORKER(input chan string, output chan string) {
	defer close(output)
	for address := range input {
		output <- funnel.ScrapeTitle(address)
	}
}
func MERGER(workerOutputs []chan string, final chan string) {
	defer close(final)
	var waiter sync.WaitGroup
	defer waiter.Wait()
	waiter.Add(len(workerOutputs))
	for _, out := range workerOutputs {
		go DRAINER(out, final, waiter.Done)
	}
}
func DRAINER(in, out chan string, done func()) {
	defer done()
	for v := range in {
		out <- v
	}
}
