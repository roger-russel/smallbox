package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

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
