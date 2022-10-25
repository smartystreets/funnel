package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/funnel/training/stuff"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	for _, address := range stuff.Addresses {
		fmt.Println(stuff.ScrapeTitle(address))
	}
}
