package main

import (
	"fmt"

	"github.com/roger-russel/smallbox/internal/cmd"
	v "github.com/roger-russel/smallbox/internal/version"
)

var version string
var commit string
var date string

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some thing went wrong:", r)
		}
	}()

	cmd.Root(v.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}
