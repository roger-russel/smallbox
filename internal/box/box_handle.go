package box

import (
	"fmt"
	"os"

	"github.com/roger-russel/smallbox/internal/flags"
	"github.com/roger-russel/smallbox/internal/version"
	"github.com/spf13/cobra"
)

const boxTestFile = "box_test.go"
const boxManagerFile = "box.go"
const boxFolderName = "box"

var boxPath string

//Handle Box Generate Command
func Handle(cmd *cobra.Command, args []string, flags flags.Flags, vf version.FullVersion) {

	if flags.Path[len(flags.Path)-1:] == "/" {
		boxPath = flags.Path
	} else {
		boxPath = flags.Path + "/"
	}

	boxPath += boxFolderName + "/"

	handleBoxFolder()
	createBoxManagerFile(vf, flags.Force)

	if flags.File != "" {
		boxFile(vf, flags.Force, flags.File, flags.Name)
	}

	if flags.Dir != "" {
		boxDir(vf, flags.Force, flags.Dir, flags.Name)
	}

}

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
