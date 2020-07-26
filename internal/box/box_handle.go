package box

import (
	"github.com/roger-russel/smallbox/internal/flags"
	"github.com/roger-russel/smallbox/internal/version"
	"github.com/spf13/cobra"
)

var boxFolderName string = "box"
var boxManagerFile string = "box.go"
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

	switch {
	case flags.File != "":
		boxFile(vf, flags.Force, flags.File, flags.Name)
	}

}
