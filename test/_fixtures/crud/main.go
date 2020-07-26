package main

import (
	"fmt"

	"github.com/roger-russel/smallbox/test/fixtures/crud/box"
)

func main() {

	bSimple, err := box.Get("simple")

	if err != nil {
		panic(err)
	}

	fmt.Print(string(bSimple))

}
