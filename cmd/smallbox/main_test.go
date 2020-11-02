package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"text/template"

	"bou.ke/monkey"

	"github.com/roger-russel/smallbox/internal/helper"
)

func Test_mainHelp(t *testing.T) {

	tests := []struct {
		name    string
		pre     func()
		want    bool
		wantOut string
	}{
		{
			name: "help with --help",
			pre: func() {
				os.Args = []string{
					os.Args[0],
					"--help",
				}
			},
			want:    true,
			wantOut: "smallbox",
		},
		{
			name: "help without --help",
			pre: func() {
				os.Args = []string{
					os.Args[0],
				}
			},
			want:    true,
			wantOut: "smallbox",
		},
		{
			name: "test control",
			pre: func() {
				os.Args = []string{
					os.Args[0],
				}
			},
			want:    false,
			wantOut: "boxmalls",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pre()
			helper.StdoutCapture()
			main()
			out := helper.StdoutFree()
			header := string(out)[:8]
			if (header != tt.wantOut) == tt.want {
				t.Errorf("Expecting helper message starting with \"%v\" but got: \"%v\"", tt.wantOut, header)
			}
		})
	}

}

func Test_mainVersion(t *testing.T) {

	tests := []struct {
		name    string
		pre     func()
		want    bool
		wantOut string
	}{
		{
			name: "version",
			pre: func() {
				os.Args = []string{
					os.Args[0],
					"version",
				}
			},
			want:    true,
			wantOut: "version: \nbuilded at: \ncommit hash: \n",
		},
		{
			name: "text control",
			pre: func() {
				os.Args = []string{
					os.Args[0],
				}
			},
			want:    false,
			wantOut: "no version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pre()
			helper.StdoutCapture()
			main()
			out := helper.StdoutFree()
			if (out != tt.wantOut) == tt.want {
				t.Errorf("Expecting \"%v\" but got: \"%v\"", tt.wantOut, out)
			}
		})
	}

}

func Test_mainPanic(t *testing.T) {

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return fmt.Errorf("permission denied")
	})

	defer monkey.UnpatchAll()

	tests := []struct {
		name    string
		pre     func()
		want    bool
		wantOut string
	}{
		{
			name: "permission denied",
			pre: func() {
				os.Args = []string{
					os.Args[0],
					"-f",
					"some file",
				}
			},
			want:    true,
			wantOut: "Some thing went wrong: permission denied",
		},
		{
			name: "control",
			pre: func() {
				os.Args = []string{
					os.Args[0],
					"-f",
					"some file",
				}
			},
			want:    false,
			wantOut: "everything is ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pre()
			helper.StdoutCapture()
			main()
			out := strings.TrimSpace(helper.StdoutFree())
			if (out != tt.wantOut) == tt.want {
				t.Errorf("Expecting \"%v\" but got: \"%v\"", tt.wantOut, out)
			}
		})
	}

}

func Example_main() {

	defer monkey.UnpatchAll()

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return nil
	})

	c := 0
	monkey.Patch(os.Stat, func(path string) (of os.FileInfo, err error) {
		c++
		if c == 3 {
			return nil, fmt.Errorf("some error")
		}
		return nil, nil
	})

	monkey.Patch(os.Create, func(path string) (*os.File, error) {
		return nil, nil
	})

	f := &os.File{}

	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		return nil
	})

	var tpl *template.Template

	monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
		return nil
	})

	os.Args = []string{
		os.Args[0],
		"-f",
		"/some/file.txt",
	}

	main()
	// output:
	// Boxing: /some/file.txt
	// Some thing went wrong: some error

}

func Example_mainPath() {

	defer monkey.UnpatchAll()

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return nil
	})

	c := 0
	monkey.Patch(os.Stat, func(path string) (of os.FileInfo, err error) {
		c++
		if c == 3 {
			return nil, fmt.Errorf("some error")
		}
		return nil, nil
	})

	monkey.Patch(os.Create, func(path string) (*os.File, error) {
		return nil, nil
	})

	f := &os.File{}

	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		return nil
	})

	var tpl *template.Template

	monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
		return nil
	})

	os.Args = []string{
		os.Args[0],
		"-f",
		"/some/file.txt",
		"-p", "/",
		"-n", "file",
	}

	main()
	// output:
	// Boxing: /some/file.txt as file
	// Some thing went wrong: some error

}

func Example_mainDir() {

	defer monkey.UnpatchAll()

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return nil
	})

	c := 0
	monkey.Patch(os.Stat, func(path string) (of os.FileInfo, err error) {
		c++
		if c == 3 {
			return nil, fmt.Errorf("some error")
		}
		return nil, nil
	})

	monkey.Patch(os.Create, func(path string) (*os.File, error) {
		return nil, nil
	})

	f := &os.File{}

	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		return nil
	})

	var tpl *template.Template

	monkey.PatchInstanceMethod(reflect.TypeOf(tpl), "ExecuteTemplate", func(a *template.Template, b io.Writer, c string, d interface{}) error {
		return nil
	})

	os.Args = []string{
		os.Args[0],
		"-d",
		"./foo",
		"-p", "/",
		"-n", "boo",
	}

	main()
	// output:
	// Some thing went wrong: some error

}
