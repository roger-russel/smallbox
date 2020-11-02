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
var tplTest *template.Template

func init() {
	initFile()
}

func initFile() {
	tplBox = initTemplate("box")
	tplBoxed = initTemplate("boxed")
	tplTest = initTemplate("test_tpl")
}

func initTemplate(name string) *template.Template {

	content, err := box.Get(name)

	if err != nil {
		panic(err)
	}

	return template.Must(template.New(name).Parse(string(content)))

}

func createBoxManagerFile(vf version.FullVersion, force bool) {
	createTemplateFile(vf, force, tplBox, "box", boxManagerFile)
	createTemplateFile(vf, force, tplTest, "test_tpl", boxTestFile)
}

func createTemplateFile(vf version.FullVersion, force bool, tpl *template.Template, tplName string, fileName string) {

	filePath := boxPath + fileName

	_, err := os.Stat(filePath)

	if err != nil || force {
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		} else {
			f, err := os.Create(filePath)

			if err != nil {
				panic(err)
			}

			defer f.Close()

			err = tpl.ExecuteTemplate(f, tplName, struct {
				Version string
				Date    string
			}{
				Version: fmt.Sprintf("%+v", vf),
				Date:    time.Now().Format("2006-01-02 15:04:05.000000"),
			})

			if err != nil {
				panic(fmt.Sprintf("error creating %v %v", filePath, err))
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
