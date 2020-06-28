package box

import (
	"fmt"
	"os"
)

func handleBoxFolder() {

	err := os.MkdirAll(boxPath, 0755)

	if err != nil {
		if !os.IsExist(err) { // If the error is not that the directory already exists, like: permission denied
			panic(err)
		} else {
			if stat, err := os.Stat(boxPath); !stat.IsDir() {
				if err != nil { // When there is a problem to get the stat from file/folder
					panic(err)
				}
				panic(fmt.Sprintln("box is not a directory:", boxPath)) // if there is a "box" file name it will not be able to create the "box" directory
			}
		}
	}
}