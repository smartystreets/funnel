```go
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
	for x := 0; x < 10; x++ {
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
func DRAINER(input, output chan string, done func()) {
	for v := range input {
		output <- v
	}
	done()
}
```