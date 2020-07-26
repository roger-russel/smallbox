
/*Package box {Version: Commit: Date:}*/
package box

import (
	"encoding/base64"
	"fmt"
)

//Box struct
type Box struct {
	Content map[string]string
}

var box map[string]string

//initializeBox check and Initialize files on it
func initializeBox() {
	if box == nil {
		box = make(map[string]string)
	}
}

//Len of items on Box
func Len() int {
	return len(box)
}

//Get box content
func Get(name string) ([]byte, error) {

	content, ok := box[name]

	if !ok {
		return []byte{}, fmt.Errorf("Content not found: %v", name)
	}

	return base64.StdEncoding.DecodeString(content)
}

//List return the list of keys
func List() (list []string) {

	for i := range box {
		list = append(list, i)
	}

	return list
}
