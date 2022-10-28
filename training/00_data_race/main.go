package main

import (
	"fmt"
)

var state int

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
		go display()
	}
	fmt.Println(state)
}

func increment() {
	state++
}
func decrement() {
	state--
}
func display() {
	fmt.Println(state)
}
