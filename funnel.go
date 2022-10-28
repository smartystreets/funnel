package funnel

import "sync"

func FanOut[T any](input chan T, workers int, work func(T) T) chan T {
	if workers <= minWorkerCount {
		panic("worker count must be positive")
	}
	if workers > maxWorkerCount {
		panic("are you sure you want that many workers?")
	}
	if input == nil {
		panic("input chan is nil")
	}
	if work == nil {
		panic("work callback is nil")
	}
	var outputs []chan T
	for x := 0; x < workers; x++ {
		output := make(chan T)
		outputs = append(outputs, output)
		go worker(input, output, work)
	}
	merged := make(chan T)
	go merge(outputs, merged)
	return merged
}
func worker[T any](input, output chan T, work func(T) T) {
	defer close(output)
	for item := range input {
		output <- work(item)
	}
}
func merge[T any](outputs []chan T, merged chan T) {
	defer close(merged)
	var waiter sync.WaitGroup
	defer waiter.Wait()
	waiter.Add(len(outputs))
	for _, output := range outputs {
		go drain(output, merged, waiter.Done)
	}
}
func drain[T any](output, merged chan T, done func()) {
	defer done()
	for item := range output {
		merged <- item
	}
}

const maxWorkerCount = 1024 * 10
const minWorkerCount = 0
