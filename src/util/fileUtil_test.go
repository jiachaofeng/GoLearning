package util

import (
	"testing"
	"time"
)

func TestPathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test Case1", args{path: "../../bin/"}, true, false},
		{"Test Case2", args{path: "./notexit"}, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkDirAllIfPathNotExit(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test Case 1", args{path: "../../bin/NewDir"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MkDirAllIfPathNotExit(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("MkDirAllIfPathNotExit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		filepathOrigin string
		filepathDest   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test Case1",
			args{filepathOrigin: "../../bin/testCopy.txt", filepathDest: "../../bin/NewDir/testCopied.txt"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.filepathOrigin, tt.args.filepathDest); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyFiles(t *testing.T) {
	layout := "2006-01"
	timeStart, _ := time.ParseInLocation(layout, "2017-11", time.Local)
	timeEnd, _ := time.ParseInLocation(layout, "2017-12", time.Local)

	type args struct {
		filepathOrigin string
		filepathDest   string
		st             time.Time
		et             time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test Case1",
			args{filepathOrigin: "H:/Invoice file/",
				filepathDest: "H:/Invoice file/",
				st:           timeStart,
				et:           timeEnd},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFiles(tt.args.filepathOrigin, tt.args.filepathDest, tt.args.st, tt.args.et); (err != nil) != tt.wantErr {
				t.Errorf("CopyFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test Case1",
			args{filepath: "../../bin/NewDir/testCopied.txt"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFile(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLink(t *testing.T) {
	type args struct {
		filepathOrigin string
		filepathDest   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test Case1",
			args{filepathOrigin: "../../bin/testCopy.txt", filepathDest: "../../bin/NewDir/testCopied.txt"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Link(tt.args.filepathOrigin, tt.args.filepathDest); (err != nil) != tt.wantErr {
				t.Errorf("Link() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestZipFile(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test Case1",
			args{source: "H:/Invoice file/2017/",
				target: "H:/Invoice file/2017.zip"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ZipFile(tt.args.source, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ZipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
