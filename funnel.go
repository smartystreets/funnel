package funnel

import "sync"

func GoFanOut[T any](input chan T, workers int, doWork func(T) T) chan T {
	if workers <= minWorkerCount {
		panic("worker count must be positive")
	}
	if workers > maxWorkerCount {
		panic("are you sure you want that many workers?")
	}
	if input == nil {
		panic("input chan is nil")
	}
	if doWork == nil {
		panic("doWork callback is nil")
	}

	merged := make(chan T)
	go coordinate(workers, doWork, input, merged)
	return merged
}
func coordinate[T any](workers int, doWork func(T) T, initial, final chan T) {
	defer close(final)
	var waiter sync.WaitGroup
	defer waiter.Wait()
	waiter.Add(workers)

	for w := 0; w < workers; w++ {
		intermediate := make(chan T)
		go process(initial, intermediate, doWork, func() { close(intermediate) })
		go process(intermediate, final, func(t T) T { return t }, waiter.Done)
	}
}
func process[T any](input, output chan T, workFunc func(T) T, done func()) {
	defer done()
	for item := range input {
		output <- workFunc(item)
	}
}

const maxWorkerCount = 1024 * 10
const minWorkerCount = 0
