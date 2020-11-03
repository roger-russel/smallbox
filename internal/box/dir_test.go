package box

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"bou.ke/monkey"

	"github.com/roger-russel/smallbox/internal/version"
)

func Test_boxDir(t *testing.T) {
	monkey.UnpatchAll()
	defer monkey.UnpatchAll()
	basePath := os.Getenv("GOPATH") + "/src/github.com/roger-russel/smallbox"
	testPath := basePath + "/test/_fixtures/dir/assets"
	type args struct {
		vf            version.FullVersion
		force         bool
		dirName       string
		aliasBaseName string
	}
	tests := []struct {
		name     string
		args     args
		run      func(t *testing.T)
		wantData string
		want     bool
	}{

		{
			name: "Could not stat the path",
			args: args{
				vf:            version.FullVersion{},
				force:         true,
				dirName:       basePath,
				aliasBaseName: "tpm",
			},
			wantData: "Could not stat the path",
			want:     true,
			run: func(t *testing.T) {
				monkey.Patch(os.Stat, func(path string) (of os.FileInfo, err error) {
					return nil, fmt.Errorf("Could not stat the path")
				})
			},
		},

		{
			name: "Is not a Dir",
			args: args{
				vf:            version.FullVersion{},
				force:         true,
				dirName:       basePath,
				aliasBaseName: "tpm",
			},
			wantData: fmt.Sprintf("The dir path: %v, is not a valid path to directory", basePath),
			want:     true,
			run: func(t *testing.T) {

				stat, err := os.Stat(basePath + "/README.md")

				monkey.Patch(os.Stat, func(path string) (os.FileInfo, error) {
					return stat, err
				})

			},
		},

		{
			name: "Could not read the dir given",
			args: args{
				vf:            version.FullVersion{},
				force:         true,
				dirName:       basePath,
				aliasBaseName: "tpm",
			},
			wantData: "Could not read the dir given",
			want:     true,
			run: func(t *testing.T) {
				monkey.Patch(ioutil.ReadDir, func(dirName string) ([]os.FileInfo, error) {
					return nil, fmt.Errorf("Could not read the dir given")
				})

			},
		},
		{
			name: "Reading files on Directory",
			args: args{
				vf:            version.FullVersion{},
				force:         true,
				dirName:       testPath,
				aliasBaseName: "tpm",
			},
			wantData: "Could not read the dir given",
			want:     true,
			run: func(t *testing.T) {

				files := map[string]bool{
					"tpm/bomb-char.txt":        true,
					"tpm/lorem-ipsum.txt":      true,
					"tpm/square-root-of-2.txt": true,
				}

				monkey.Patch(boxFile, func(vf version.FullVersion, force bool, tplName string, fileName string) {
					_, ok := files[fileName]

					if !ok {
						t.Errorf("Found unexpected file: %v\n ", fileName)
					}
				})

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{
				defer monkey.UnpatchAll()

				defer func() {
					if r := recover(); r != nil {
						if (fmt.Sprint(r) != tt.wantData) == tt.want {
							t.Errorf("Expecting: %v, got: %v", tt.wantData, fmt.Sprint(r))
						}
					}
				}()

				tt.run(t)
				boxDir(tt.args.vf, tt.args.force, tt.args.dirName, tt.args.aliasBaseName)
			}
		})
	}
}
