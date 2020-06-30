package normalizer

import (
	"testing"
)

func TestFileName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				name: "simple",
			},
			want: "boxed_simple.go",
		},
		{
			name: "test with pictografic chars",
			args: args{
				name: "影師",
			},
			want: "boxed_ying-shi.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileName(tt.args.name); got != tt.want {
				t.Errorf("FileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyName(t *testing.T) {
	type args struct {
		name  string
		alias string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple without /",
			args: args{
				name:  "simple",
				alias: "",
			},
			want: "/simple",
		},
		{
			name: "simple with /",
			args: args{
				name:  "/simple",
				alias: "",
			},
			want: "/simple",
		},
		{
			name: "simple with ./",
			args: args{
				name:  "./simple",
				alias: "",
			},
			want: "/simple",
		},
		{
			name: "simple with alias",
			args: args{
				name:  "./simple",
				alias: "s",
			},
			want: "s",
		},
		{
			name: "multiple without /",
			args: args{
				name:  "multiple/foo/boo",
				alias: "",
			},
			want: "/multiple/foo/boo",
		},
		{
			name: "multiple with /",
			args: args{
				name:  "/multiple/foo/boo",
				alias: "",
			},
			want: "/multiple/foo/boo",
		},
		{
			name: "multiple foders with ./",
			args: args{
				name:  "./multiple/foo/boo",
				alias: "",
			},
			want: "/multiple/foo/boo",
		},
		{
			name: "multiple with alias",
			args: args{
				name:  "./simple/foo/boo",
				alias: "s",
			},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyName(tt.args.name, tt.args.alias); got != tt.want {
				t.Errorf("KeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}
