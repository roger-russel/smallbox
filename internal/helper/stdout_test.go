package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"bou.ke/monkey"
)

func TestStdout(t *testing.T) {

	t.Run("simple test", func(t *testing.T) {
		want := "test"
		StdoutCapture()
		fmt.Print(want)
		got := StdoutFree()
		if want != got {
			t.Errorf("Error expect \"%v\" but got \"%v\"", want, got)
		}
	})

	t.Run("something went wrong opening pipe", func(t *testing.T) {

		want := "something went wrong"

		monkey.Patch(os.Pipe, func() (*os.File, *os.File, error) {
			return nil, nil, fmt.Errorf(want)
		})
		defer monkey.UnpatchAll()
		defer func() {

			if r := recover(); r != nil {
				got := fmt.Sprint(r)

				if want != got {
					t.Errorf("Error expect \"%v\" but got \"%v\"", want, got)
				}

			}

		}()

		StdoutCapture()
		fmt.Print(want)
	})

	t.Run("something went wrong when tried to free stdout", func(t *testing.T) {

		want := "something went wrong"

		monkey.Patch(ioutil.ReadAll, func(r io.Reader) ([]byte, error) {
			return []byte{}, fmt.Errorf(want)
		})
		defer monkey.UnpatchAll()

		defer func() {
			if r := recover(); r != nil {
				got := fmt.Sprint(r)
				if want != got {
					t.Errorf("Error expect \"%v\" but got \"%v\"", want, got)
				}
			}
		}()

		StdoutFree()
	})

}
