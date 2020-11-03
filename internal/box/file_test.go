package box

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"text/template"

	"bou.ke/monkey"
	"github.com/roger-russel/smallbox/box"
	"github.com/roger-russel/smallbox/internal/version"
)

func Test_initFile(t *testing.T) {

	defer monkey.UnpatchAll()

	t.Run("error getting box", func(t *testing.T) {

		monkey.Patch(box.Get, func(k string) ([]byte, error) {
			return nil, fmt.Errorf("error")
		})

		defer func() {
			if r := recover(); r != nil {
				if fmt.Sprint(r) != "error" {
					t.Errorf("Expecting: %v, got: %v", "error", fmt.Sprint(r))
				}
			}
		}()

		initFile()

	})

	t.Run("error getting boxed", func(t *testing.T) {

		monkey.Patch(box.Get, func(k string) ([]byte, error) {

			if k == "boxed" {
				return nil, fmt.Errorf("boxed")
			}

			return []byte("template"), nil

		})

		defer func() {
			if r := recover(); r != nil {
				if fmt.Sprint(r) != "boxed" {
					t.Errorf("Expecting: %v, got: %v", "boxed", fmt.Sprint(r))
				}
			}
		}()

		initFile()

	})

}

func Test_createBoxManagerFile(t *testing.T) {
	var boxPathOld string = boxPath
	type args struct {
		vf    version.FullVersion
		force bool
	}
	tests := []struct {
		name string
		args args
		pre  func()
		want string
	}{
		{
			name: "unknow error",
			args: args{
				vf:    version.FullVersion{},
				force: false,
			},
			want: "error",
			pre: func() {
				boxPath = "./"
				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					return nil, fmt.Errorf("error")
				})
			},
		},
		{
			name: "fail creating directory",
			args: args{
				vf:    version.FullVersion{},
				force: false,
			},
			want: "permission denied",
			pre: func() {
				boxPath = "./"
				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					return nil, fmt.Errorf("error")
				})
				monkey.Patch(os.IsNotExist, func(e error) bool {
					return true
				})
				monkey.Patch(os.Create, func(string) (*os.File, error) {
					return nil, fmt.Errorf("permission denied")
				})
			},
		},
		{
			name: "fail generating template",
			args: args{
				vf:    version.FullVersion{},
				force: false,
			},
			want: "error creating ./box.go permsission denied",
			pre: func() {
				boxPath = "./"
				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					return nil, fmt.Errorf("error")
				})
				monkey.Patch(os.IsNotExist, func(e error) bool {
					return true
				})
				monkey.Patch(os.Create, func(string) (*os.File, error) {
					return nil, nil
				})

				var tpl *template.Template

				monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
					return fmt.Errorf("permsission denied")
				})
			},
		},
		{
			name: "no error",
			args: args{
				vf:    version.FullVersion{},
				force: false,
			},
			want: "",
			pre: func() {
				boxPath = "./"
				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					return nil, fmt.Errorf("error")
				})
				monkey.Patch(os.IsNotExist, func(e error) bool {
					return true
				})
				monkey.Patch(os.Create, func(string) (*os.File, error) {
					return nil, nil
				})

				var tpl *template.Template

				monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
					return nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if fmt.Sprint(r) != tt.want {
						t.Errorf("Expecting: %v, got: %v", tt.want, fmt.Sprint(r))
					}
				}
			}()
			tt.pre()
			createBoxManagerFile(tt.args.vf, tt.args.force)
			monkey.UnpatchAll()
		})
	}

	boxPath = boxPathOld
}

func Test_boxFile(t *testing.T) {
	type args struct {
		vf        version.FullVersion
		force     bool
		fileName  string
		aliasName string
	}
	tests := []struct {
		name string
		args args
		want string
		pre  func()
	}{
		{
			name: "is a dir",
			args: args{
				vf:        version.FullVersion{},
				force:     false,
				fileName:  "./",
				aliasName: "dir",
			},
			want: "The file name: ./, is a dir but it was expecting a file",
			pre:  func() {},
		},
		{
			name: "permision denied to open a file",
			args: args{
				vf:        version.FullVersion{},
				force:     false,
				fileName:  "./file_test.go",
				aliasName: "dir",
			},
			want: "permission denied opening file",
			pre: func() {
				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					return nil, fmt.Errorf("permission denied opening file")
				})
			},
		},
		{
			name: "Permission denied writing file",
			args: args{
				vf:        version.FullVersion{},
				force:     false,
				fileName:  "./file_test.go",
				aliasName: "write",
			},
			want: "permission denied writing",
			pre: func() {
				c := 0

				swp, err := os.Stat("./file_test.go")

				if err != nil {
					panic(err)
				}

				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					c++

					if c == 2 {
						return nil, fmt.Errorf("permission denied writing")
					}

					return swp, nil

				})
			},
		},
		{
			name: "there is a dir with the same name of file",
			args: args{
				vf:        version.FullVersion{},
				force:     false,
				fileName:  "./file_test.go",
				aliasName: "file",
			},
			want: "there is a dir where it want to create boxed_file: boxed_file.go",
			pre: func() {
				c := 0

				f, err := os.Stat("./file_test.go")

				if err != nil {
					panic(err)
				}

				d, err := os.Stat("./")

				if err != nil {
					panic(err)
				}

				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					c++

					if c == 2 {
						return d, nil
					}

					return f, nil

				})
			},
		},
		{
			name: "error tring to read file",
			args: args{
				vf:        version.FullVersion{},
				force:     true,
				fileName:  "./file_test.go",
				aliasName: "read",
			},
			want: "error reading file",
			pre: func() {
				c := 0

				f, err := os.Stat("./file_test.go")

				if err != nil {
					panic(err)
				}

				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					c++

					if c == 2 {
						return nil, nil
					}

					return f, nil

				})

				monkey.Patch(ioutil.ReadFile, func(n string) ([]byte, error) {
					return []byte{}, fmt.Errorf("error reading file")
				})

			},
		},
		{
			name: "error tring to read file",
			args: args{
				vf:        version.FullVersion{},
				force:     true,
				fileName:  "./file_test.go",
				aliasName: "open",
			},
			want: "create file: error to create/open the boxed_file",
			pre: func() {
				c := 0

				f, err := os.Stat("./file_test.go")

				if err != nil {
					panic(err)
				}

				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					c++

					if c == 2 {
						return nil, nil
					}

					return f, nil

				})

				monkey.Patch(ioutil.ReadFile, func(n string) ([]byte, error) {
					simple := "c2ltcGxlCg==" // simple in b64
					return []byte(simple), nil
				})

				monkey.Patch(os.Create, func(n string) (*os.File, error) {
					return nil, fmt.Errorf("error to create/open the boxed_file")
				})

			},
		},

		{
			name: "error tring to execute template",
			args: args{
				vf:        version.FullVersion{},
				force:     true,
				fileName:  "./file_test.go",
				aliasName: "template",
			},
			want: "create file: error rendering the template",
			pre: func() {
				c := 0

				f, err := os.Stat("./file_test.go")

				if err != nil {
					panic(err)
				}

				monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
					c++

					if c == 2 {
						return nil, nil
					}

					return f, nil

				})

				monkey.Patch(ioutil.ReadFile, func(n string) ([]byte, error) {
					simple := "c2ltcGxlCg==" // simple in b64
					return []byte(simple), nil
				})

				monkey.Patch(os.Create, func(n string) (*os.File, error) {
					return nil, nil
				})

				var tpl *template.Template

				monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
					return fmt.Errorf("error rendering the template")
				})

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if fmt.Sprint(r) != tt.want {
						t.Errorf("Expecting: %v, got: %v", tt.want, fmt.Sprint(r))
					}
				}
				defer monkey.UnpatchAll()
			}()
			tt.pre()
			boxFile(tt.args.vf, tt.args.force, tt.args.fileName, tt.args.aliasName)
		})
	}
}

func Example_boxFileFileAlreadyExists() {

	defer monkey.UnpatchAll()

	gopath := os.Getenv("GOPATH")

	fakeStat, err := os.Stat(gopath + "/src/github.com/roger-russel/smallbox/README.md")

	if err != nil {
		panic(err)
	}

	monkey.Patch(os.Stat, func(name string) (os.FileInfo, error) {
		return fakeStat, nil
	})

	boxFile(
		version.FullVersion{},
		false,
		"./file_test.go",
		"file",
	)
	// output:
	// Boxing: ./file_test.go as file
	// There is a file with same name: "boxed_file.go", to force overwrite add flag --force on command run.
}
