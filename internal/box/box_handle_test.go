package box

import (
	"fmt"
	"os"
	"testing"

	"bou.ke/monkey"
)

func Test_handleBoxFolder(t *testing.T) {

	boxPathOld := boxPath

	defer monkey.UnpatchAll()

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return fmt.Errorf("force error")
	})

	monkey.Patch(os.IsExist, func(err error) bool {
		return true
	})

	tests := []struct {
		name     string
		before   func()
		after    func()
		wantBool bool
		wantOut  string
	}{
		{
			name: "not a directory",
			after: func() {

			},
			before: func() {
				boxPath = "./dir_test.go"
			},
			wantBool: true,
			wantOut:  "box is not a directory: ./dir_test.go\n",
		},
		{
			name: "unknow error",
			after: func() {

			},
			before: func() {
				boxPath = "/boxpath-doesnexist-123/"
			},
			wantBool: true,
			wantOut:  "stat /boxpath-doesnexist-123/: no such file or directory",
		},
		{
			name: "unknow error",
			after: func() {

			},
			before: func() {
				boxPath = "/boxpath-doesnexist-123/"
			},
			wantBool: false,
			wantOut:  "stat /boxpath/: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if (fmt.Sprint(r) != tt.wantOut) == tt.wantBool {
						t.Errorf("Expecting: %v, got: %v", tt.wantOut, fmt.Sprint(r))
					}
				}
			}()
			tt.before()
			handleBoxFolder()
			tt.after()
		})
	}

	boxPath = boxPathOld
}
