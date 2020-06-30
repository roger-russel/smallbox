package normalizer

import "github.com/gosimple/slug"

//KeyName normalize a name with the requirements to create a key name
func KeyName(name string, alias string) string {

	switch true {
	case alias != "":
		name = alias
	case name[0:2] == "./":
		name = name[1:]
	case name[0:1] != "/":
		name = "/" + name
	}

	return name
}

//FileName normalize a name to be used as filename
func FileName(name string) string {
	return "boxed_" + slug.Make(name) + ".go"
}

func BoxPath(prefix, file string) (boxPath string) {

	boxPath = prefix

	switch true {
	case prefix == "./":
		boxPath += file
	case prefix == "/":
		boxPath += file

		boxPath += "/" + file
	}

	return boxPath
}
