/*Package box {Version: Commit: Date:} */
package box

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLen(t *testing.T) {

	box = make(map[string]string)
	box["a"] = "1"
	box["b"] = "2"
	box["c"] = "3"

	tests := []struct {
		name         string
		want         bool
		wantResponse int
	}{
		{
			name:         "simple",
			want:         true,
			wantResponse: 3,
		},
		{
			name:         "different",
			want:         false,
			wantResponse: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Len()
			if tt.want == (got != tt.wantResponse) {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {

	box = make(map[string]string)
	box["a"] = "1"
	box["b"] = "2"
	box["c"] = "3"

	tests := []struct {
		name     string
		want     bool
		wantList []string
	}{
		{
			name: "equals",
			want: true,
			wantList: []string{
				"a",
				"b",
				"c",
			},
		},
		{
			name: "different",
			want: false,
			wantList: []string{
				"z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotList := List()
			sort.Strings(tt.wantList)
			sort.Strings(gotList)

			if tt.want == !cmp.Equal(gotList, tt.wantList) {
				t.Errorf("List() diff %v ", cmp.Diff(gotList, tt.wantList))
			}
		})
	}
}

func TestGet(t *testing.T) {

	box["test/simple"] = "c2ltcGxl"
	box["test/fail"] = "c2ltcGxl="

	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		want       bool
		wantReturn []byte
		wantErr    bool
	}{
		{
			name: "simple",
			args: args{
				name: "test/simple",
			},
			want:       true,
			wantReturn: []byte("simple"),
			wantErr:    false,
		},
		{
			name: "different",
			args: args{
				name: "test/simple",
			},
			want:       false,
			wantReturn: []byte("not simple"),
			wantErr:    false,
		},
		{
			name: "not found",
			args: args{
				name: "test/not-found",
			},
			want:       false,
			wantReturn: []byte("simple"),
			wantErr:    true,
		},
		{
			name: "parse fail",
			args: args{
				name: "test/fail",
			},
			want:       true,
			wantReturn: []byte("simple"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sGot := string(got)
			sWant := string(tt.wantReturn)

			if tt.want == (sGot != sWant) {
				t.Errorf("Get() = %v, want %v", sGot, sWant)
			}
		})
	}
}
