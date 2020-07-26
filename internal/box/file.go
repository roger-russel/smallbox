package box

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/roger-russel/smallbox/box"
	"github.com/roger-russel/smallbox/internal/normalizer"
	"github.com/roger-russel/smallbox/internal/version"
)

var tplBox *template.Template
var tplBoxed *template.Template

func init() {
	initFile()
}

func initFile() {
	content, err := box.Get("box")

	if err != nil {
		panic(err)
	}

	tplBox = template.Must(template.New("box").Parse(string(content)))

	content, err = box.Get("boxed")

	if err != nil {
		panic(err)
	}

	tplBoxed = template.Must(template.New("boxed").Parse(string(content)))
}

func createBoxManagerFile(vf version.FullVersion, force bool) {

	managerPath := boxPath + boxManagerFile

	_, err := os.Stat(managerPath)

	if err != nil || force {
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		} else {
			f, err := os.Create(managerPath)

			if err != nil {
				panic(err)
			}

			defer f.Close()

			err = tplBox.ExecuteTemplate(f, "box", struct {
				Version string
				Date    string
			}{
				Version: fmt.Sprintf("%+v", vf),
				Date:    time.Now().Format("2006-01-02 15:04:05.000000"),
			})

			if err != nil {
				panic(fmt.Sprintf("error creating %v %v", managerPath, err))
			}

		}
	}
}

func boxFile(vf version.FullVersion, force bool, fileName string, aliasName string) {

	fmt.Printf("Boxing: %v", fileName)

	if aliasName != "" {
		fmt.Printf(" as %v", aliasName)
	}

	fmt.Printf("\n")

	stat, err := os.Stat(fileName)

	if err != nil {
		panic(err)
	}

	if stat.IsDir() {
		panic(fmt.Sprintf("The file name: %v, is a dir but it was expecting a file", fileName))
	}

	boxedFileName := boxPath + normalizer.FileName(aliasName)

	stat, err = os.Stat(boxedFileName)

	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if stat != nil && stat.IsDir() {
		panic(fmt.Sprintf("there is a dir where it want to create boxed_file: %v", boxedFileName))
	}

	if stat != nil && !force {
		fmt.Printf("There is a file with same name: \"%v\", to force overwrite add flag --force on command run.", boxedFileName)
		return // continue to next if there is one
	}

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	sFile := base64.StdEncoding.EncodeToString(fileBytes)

	f, err := os.Create(boxedFileName)

	if err != nil {
		panic(fmt.Sprintf("create file: %v", err))
	}

	err = tplBoxed.ExecuteTemplate(f, "boxed", struct {
		Content string
		Name    string
		Version string
		Date    string
	}{
		Content: sFile,
		Name:    normalizer.KeyName(fileName, aliasName),
		Version: fmt.Sprintf("%+v", vf),
		Date:    time.Now().Format("2006-01-02 15:04:05.000000"),
	})

	if err != nil {
		panic(fmt.Sprintf("create file: %v", err))
	}

}
