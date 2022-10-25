package funnel

import (
	"fmt"
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
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
