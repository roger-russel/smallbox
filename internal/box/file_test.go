package box

import (
	"fmt"
	"io"
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
}
