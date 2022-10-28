package funnel

import (
	"fmt"
	"testing"
	"time"
)

func TestFanOut_ZeroWorkers_Panics(t *testing.T) {
	defer func() { recover() }()
	FanOut(make(chan int), 0, func(int) int { return 0 })
	t.Error("should have panicked...")
}
func TestFanOut_TooManyWorkers_Panics(t *testing.T) {
	defer func() { recover() }()
	FanOut(make(chan int), maxWorkerCount+1, func(int) int { return 0 })
	t.Error("should have panicked...")
}
func TestFanOut_NilInput_Panics(t *testing.T) {
	defer func() { recover() }()
	FanOut(nil, 1, func(int) int { return 0 })
	t.Error("should have panicked...")
}
func TestFanOut_NilCallback_Panics(t *testing.T) {
	defer func() { recover() }()
	FanOut(make(chan int), 1, nil)
	t.Error("should have panicked...")
}
func TestFanOut(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
	}
	var (
		workItemCount             = 100
		workerCount               = 10
		expectedDurationInSeconds = workItemCount / workerCount
	)

	input := make(chan string)
	go func() {
		defer close(input)
		for x := 0; x < workItemCount; x++ {
			input <- fmt.Sprint(time.Now().Second())
		}
	}()

	started := time.Now()
	output := FanOut(input, workerCount, func(s string) string {
		time.Sleep(time.Second) // simulate long-running process
		return s + " " + fmt.Sprint(time.Now().Second())
	})

	for item := range output {
		t.Log(item)
	}

	actualDurationInSeconds := int(time.Since(started).Seconds())
	if actualDurationInSeconds != expectedDurationInSeconds {
		t.Errorf("\n"+
			"got: %d\n"+
			"want: %d",
			expectedDurationInSeconds,
			actualDurationInSeconds)
	}
}
