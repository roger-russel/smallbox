package normalizer

import "github.com/gosimple/slug"

//KeyName normalize a name with the requirements to create a key name
func KeyName(name string) string {

	switch true {
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
