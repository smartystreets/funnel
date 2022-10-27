package main

import (
	"fmt"
	"time"
)

// [ ] data race: sync.Mutex
// [ ] data race: atomic.AddInt64
// [ ] main shut down: sync.WaitGroup
// [ ] clear
// [ ] sequential, blocking I/O
// [ ] concurrent, per job I/O
// [ ] fan-out

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()
}
