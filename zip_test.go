package zip

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	ignoreFilepath := filepath.Join("testdata", ".ignore")
	excludeFilepaths, _ := ExcludeFilepaths(ignoreFilepath)
	fmt.Println("excludeFilepaths:", excludeFilepaths)

	type args struct {
		src              string
		dst              string
		excludeFilepaths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "excludes some files from the specified directory, compresses it, and completes normally.",
			args: args{
				src:              "testdata",
				dst:              "testdata.zip",
				excludeFilepaths: excludeFilepaths,
			},
			wantErr: false,
		},
		{
			name: "specifies a directory that does not exist and return error",
			args: args{
				src:              "not_exist_dir",
				dst:              "testdata.zip",
				excludeFilepaths: excludeFilepaths,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compress(tt.args.src, tt.args.dst, tt.args.excludeFilepaths); (err != nil) != tt.wantErr {
				t.Errorf("Compress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExcludeFilepaths(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "specify the path and extract the excluded file path by considering the character string specified with wildcards.",
			args: args{
				path: filepath.Join("testdata", ".ignore"),
			},
			want:    []string{"test/3.txt", "2.txt"},
			wantErr: false,
		},
		{
			name: "returns an empty slice of string if no path is specified",
			args: args{
				path: "",
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "specify a file path that does not exist and return error",
			args: args{
				path: "not_exist_filepath",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExcludeFilepaths(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExcludeFilepaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExcludeFilepaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
