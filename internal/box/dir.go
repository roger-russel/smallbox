package box

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/roger-russel/smallbox/internal/version"
)

func boxDir(vf version.FullVersion, force bool, dirName string, aliasBaseName string) {

	stat, err := os.Stat(dirName)

	if err != nil {
		panic(err)
	}

	if !stat.IsDir() {
		panic(fmt.Sprintf("The dir path: %v, is not a valid path to directory", dirName))
	}

	items, err := ioutil.ReadDir(dirName)

	if err != nil {
		panic(err)
	}

	for _, item := range items {

		if item.IsDir() {
			continue
		}

		var alias string

		if aliasBaseName != "" {
			alias = aliasBaseName + "/" + item.Name()
		}

		boxFile(vf, force, dirName+"/"+item.Name(), alias)

	}

}
