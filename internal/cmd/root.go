package cmd

import (
	"os"

	"github.com/roger-russel/smallbox/internal/box"
	"github.com/roger-russel/smallbox/internal/flags"
	v "github.com/roger-russel/smallbox/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root command
func Root(vf v.FullVersion) {

	checkDefaultCommand()

	var flags flags.Flags

	rootCmd = &cobra.Command{
		Use:   "smallbox",
		Short: "smallbox",
		Run: func(cmd *cobra.Command, args []string) {
			box.Handle(cmd, args, flags, vf)
		},
	}

	rootCmd.AddCommand(version(vf))

	rootCmd.Flags().StringVarP(
		&flags.File, "file", "f", "",
		"The file to be created/updated on box folder eg: -f ./assets/index.html",
	)

	rootCmd.Flags().StringVarP(
		&flags.Path, "path", "p", "./",
		"The path to box folder: -p ./autogenerate",
	)

	rootCmd.Flags().StringVarP(
		&flags.Name, "name", "n", "",
		"Use a different name from file to store it content: -n easyname",
	)

	flags.Force = *rootCmd.Flags().Bool(
		"force", false,
		"Force will make smallbox rewrite all files again: --force",
	)

	rootCmd.Execute()

}

func checkDefaultCommand() {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}
}
