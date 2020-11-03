package cmd

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/google/go-cmp/cmp"
	v "github.com/roger-russel/smallbox/internal/version"
	"github.com/spf13/cobra"
)

func Test_checkDefaultCommand(t *testing.T) {
	tests := []struct {
		name     string
		run      func()
		wantData []string
		want     bool
	}{
		{
			name: "test control",
			run: func() {
				os.Args = []string{"run"}
			},
			wantData: []string{"run"},
			want:     false,
		},
		{
			name: "add help",
			run: func() {
				os.Args = []string{"run"}
			},
			wantData: []string{"run", "--help"},
			want:     true,
		},
		{
			name: "don't add help",
			run: func() {
				os.Args = []string{"run", "-d", "./templates"}
			},
			wantData: []string{"run", "-d", "./templates"},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.run()
			checkDefaultCommand()

			if tt.want == !cmp.Equal(os.Args, tt.wantData) {
				t.Errorf("checkDefaultCommand() diff %v ", cmp.Diff(os.Args, tt.wantData))
			}

		})
	}
}

func TestRootPanic(t *testing.T) {
	defer monkey.UnpatchAll()
	type args struct {
		vf v.FullVersion
	}
	tests := []struct {
		name       string
		args       args
		run        func()
		wantReturn string
	}{
		{
			name: "panic",
			run: func() {
				var rootCmd *cobra.Command

				monkey.PatchInstanceMethod(reflect.TypeOf(rootCmd), "Execute", func(*cobra.Command) error {
					return fmt.Errorf("some error on TestRootPanic")
				})
			},
			wantReturn: "some error on TestRootPanic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(t *testing.T, wantReturn string) {
				if r := recover(); r != nil {
					if wantReturn != fmt.Sprint(r) {
						t.Errorf("Want \"%v\", Received \"%v\"", wantReturn, r)
					}
				}
			}(t, tt.wantReturn)
			tt.run()
			Root(tt.args.vf)
		})
	}
}
