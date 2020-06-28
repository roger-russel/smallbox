package box

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/roger-russel/smallbox/box"
	"github.com/roger-russel/smallbox/internal/normalizer"
	"github.com/roger-russel/smallbox/internal/version"
)

var tplBox *template.Template
var tplBoxed *template.Template

func init() {

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

	managerPath := boxPath + "/" + boxManagerFile

	_, err := os.Stat(managerPath)

	if err != nil || force {
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		} else {
			f, err := os.Create(managerPath)
			defer f.Close()

			if err != nil {
				panic(err)
			}

			err = tplBox.ExecuteTemplate(f, "box", struct {
				Version string
			}{
				Version: fmt.Sprintf("%+v", vf),
			})

			if err != nil {
				log.Println("create box.go:", err)
				return
			}

		}
	}
}

func boxFile(vf version.FullVersion, fileName string, aliasName string) {

	fmt.Println(aliasName, fileName)

	stat, err := os.Stat(fileName)

	if err != nil {
		panic(err)
	}

	if stat.IsDir() {
		panic(fmt.Sprintln("A dir is given but smallbox was expecting a file:", fileName))
	}

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	sFile := base64.StdEncoding.EncodeToString(fileBytes)

	f, err := os.Create(boxPath + "/" + normalizer.FileName(aliasName))

	if err != nil {
		log.Println("create file:", err)
		return
	}

	err = tplBoxed.ExecuteTemplate(f, "boxed", struct {
		Content string
		Name    string
		Version string
	}{
		Content: sFile,
		Name:    normalizer.KeyName(fileName, aliasName),
		Version: fmt.Sprintf("%+v", vf),
	})

	if err != nil {
		log.Println("create file:", err)
		return
	}

}
