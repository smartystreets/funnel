package main

import (
	"fmt"
	"sync"
	"time"
)

var state int
var lock sync.Mutex

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
	}
	time.Sleep(time.Second)

	lock.Lock()
	defer lock.Unlock()
	fmt.Println(state)
}

func increment() {
	lock.Lock()
	defer lock.Unlock()
	state++
}
func decrement() {
	lock.Lock()
	defer lock.Unlock()
	state--
}
