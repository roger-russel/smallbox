package box

import (
	"fmt"
	"os"
)

func handleBoxFolder() {

	err := os.MkdirAll(boxPath, 0740)

	if err != nil {
		if !os.IsExist(err) { // If the error is not that the directory already exists, like: permission denied
			panic(err)
		} else {

			stat, err := os.Stat(boxPath)

			if err != nil { // When there is a problem to get the stat from file/folder
				panic(err)
			}

			if !stat.IsDir() {
				panic(fmt.Sprintln("box is not a directory:", boxPath)) // if there is a "box" file name it will not be able to create the "box" directory
			}
		}
	}
}
