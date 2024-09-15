package main

import (
	"fmt"
	"time"

	"github.com/nadedan/sigint"
	"github.com/nadedan/sigint/cmd/test/internal/some"
)

func main() {
	sigint.Defer(func() {
		fmt.Println("First deferred")
	})

	_ = some.NewThing()

	fmt.Println("Entering forever loop. You can kill me with ctrl-c")
	for {
		time.Sleep(1 * time.Second)
	}
}
