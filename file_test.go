package xutil_test

import (
	"fmt"
	"testing"

	"github.com/coghost/xdtm"
	"github.com/coghost/xutil"

	"github.com/gookit/goutil/fsutil"
	"github.com/stretchr/testify/suite"
)

type FileSuite struct {
	suite.Suite
}

func TestFile(t *testing.T) {
	suite.Run(t, new(FileSuite))
}

func (s *FileSuite) SetupSuite() {
}

func (s *FileSuite) TearDownSuite() {
}

func (s *FileSuite) TestRefineFileName() {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// `\/:*?"<>|`
		{
			name: `with /`,
			args: args{name: "datafile/first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with \`,
			args: args{name: "datafile\\first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with :`,
			args: args{name: "datafile:first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with *`,
			args: args{name: "datafile*first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with ?`,
			args: args{name: "datafile?first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with "`,
			args: args{name: "datafile\"first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with <`,
			args: args{name: "datafile<first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with >`,
			args: args{name: "datafile>first.txt"},
			want: "datafile_first.txt",
		},
		{
			name: `with |`,
			args: args{name: "datafile|first.txt"},
			want: "datafile_first.txt",
		},
	}
	for _, tt := range tests {
		got := xutil.RefineString(tt.args.name)
		s.Equal(tt.want, got, tt.name)
	}
}

func (s *FileSuite) TestRefineWinFileName() {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `default`,
			args: args{name: "datafile_first.txt"},
			want: "datafile_first.txt",
		},
	}
	for _, tt := range tests {
		got := xutil.RefineWinFileName(tt.args.name)
		s.Equal(tt.want, got, tt.name)
	}
}

func (s *FileSuite) TestMustWriteFile() {
	name := fmt.Sprintf("/tmp/filetest.%s.001.log", xdtm.Now().Layout("2006010215"))
	type args struct {
		name string
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: name, args: args{name: name, data: "this is a fake data"}},
	}
	for _, tt := range tests {
		b := fsutil.IsFile(tt.args.name)
		s.False(b, tt.name)
		xutil.MustWriteFile(tt.args.name, tt.args.data)
		b1 := fsutil.IsFile(tt.args.name)
		s.True(b1, tt.name)
		fsutil.MustRm(tt.args.name)
		b = fsutil.IsFile(tt.args.name)
		s.False(b, tt.name)
	}
}

func (s *FileSuite) TestWriteFile() {
	type args struct {
		name string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "/tmp/a/b/c/d/001.txt",
			args: args{
				name: "/tmp/a/b/c/d/001.txt",
				data: "this is all",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		err := xutil.WriteFile(tt.args.name, tt.args.data)
		if tt.wantErr {
			s.Error(err, tt.name)
		}
	}
}
