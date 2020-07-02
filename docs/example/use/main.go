package main

import (
	//The box folder that you want to import
	"fmt"

	"github.com/roger-russel/smallbox/docs/example/use/box"
)

func main() {
	out := foo("/assets/simple.txt")
	fmt.Print(out)
}

func foo(name string) string {

	bContent, err := box.Get(name)

	if err != nil {
		panic(err)
	}

	return string(bContent)

}
