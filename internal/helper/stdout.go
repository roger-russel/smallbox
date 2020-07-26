package helper

import (
	"io/ioutil"
	"os"
)

var stdOutReader *os.File
var stdOutWriter *os.File
var stdOut *os.File
var err error

//StdoutCapture the stdout
func StdoutCapture() {

	stdOut = os.Stdout
	stdOutReader, stdOutWriter, err = os.Pipe()

	if err != nil {
		panic(err)
	}

	os.Stdout = stdOutWriter

}

//StdoutFree the captured stdout
func StdoutFree() string {

	if err := stdOutWriter.Close(); err != nil {
		panic(err)
	}

	os.Stdout = stdOut

	out, err := ioutil.ReadAll(stdOutReader)

	if err != nil {
		panic(err)
	}

	return string(out)

}
