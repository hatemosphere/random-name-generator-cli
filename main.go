package main

import (
	"flag"
	"fmt"

	"github.com/moby/moby/pkg/namesgenerator"
)

func main() {
	count := flag.Int("count", 1, "Number of random names to generate")
	flag.Parse()

	for i := 0; i < *count; i++ {
		name := namesgenerator.GetRandomName(0)
		fmt.Println(name)
	}
}
