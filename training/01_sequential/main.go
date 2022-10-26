package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/funnel/training/internet"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	for _, address := range internet.Addresses {
		fmt.Println(internet.Scrape(address))
	}
}
