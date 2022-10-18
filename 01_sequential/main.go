package main

import (
	"fmt"

	"github.com/mdwhatcott/funnel"
)

func main() {
	for _, address := range funnel.Addresses {
		fmt.Println(funnel.ScrapeTitle(address))
	}
}
