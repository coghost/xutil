package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
	"github.com/thoas/go-funk"
)

type StrSuite struct {
	suite.Suite
}

func TestStr(t *testing.T) {
	suite.Run(t, new(StrSuite))
}

func (s *StrSuite) SetupSuite() {
}

func (s *StrSuite) TearDownSuite() {
}

func (s *StrSuite) Test_01_Truncate() {
	type args struct {
		raw string
		max int
	}
	txt := `How to convert UTC time to unix timestamp`
	txt1 := "he's using python.library.faker"
	tests := []struct {
		name string
		args args

		want     string
		wantWord string
		wantRune string
	}{
		{
			name:     "has 10 chars",
			args:     args{raw: txt, max: 10},
			want:     "How to con",
			wantWord: "How to",
			wantRune: "How to con",
		},
		{
			name:     "same length get all",
			args:     args{raw: txt1, max: len(txt1)},
			want:     txt1,
			wantWord: txt1,
			wantRune: txt1,
		},
		{
			name:     "larger than length get all",
			args:     args{raw: txt1, max: len(txt1) + 1},
			want:     txt1,
			wantWord: txt1,
			wantRune: txt1,
		},
		{
			name:     "0 will got empty",
			args:     args{raw: txt1, max: 0},
			want:     "",
			wantWord: "",
			wantRune: "",
		},
		{
			name:     "less than length get len",
			args:     args{raw: txt1, max: 15},
			want:     "he's using pyth",
			wantWord: "he's using",
			wantRune: "he's using pyth",
		},
		{
			name:     "chinese words",
			args:     args{raw: "is 中文正常吗?", max: 4},
			want:     "is \xe4",
			wantWord: "is",
			wantRune: "is 中",
		},
	}
	for _, tt := range tests {
		got := xutil.TruncateText(tt.args.raw, tt.args.max)
		s.Equal(tt.want, got, tt.name)

		gotw := xutil.TruncateWord(tt.args.raw, tt.args.max)
		s.Equal(tt.wantWord, gotw)

		gotr := xutil.TruncateString(tt.args.raw, tt.args.max)
		s.Equal(tt.wantRune, gotr)
	}
}

func (s *StrSuite) TestFirstOrDefault() {
	type args struct {
		dftI  int
		argsI []int
		dftB  bool
		argsB []bool
		dftS  string
		argsS []string
	}

	tests := []struct {
		name  string
		args  args
		wantI int
		wantB bool
		wantS string
	}{
		{
			name:  "no argsInt return default",
			args:  args{dftI: 1},
			wantI: 1,
		},
		{
			name:  "with argsInt return first",
			args:  args{dftI: 1, argsI: []int{3, 1}},
			wantI: 3,
		},
		{
			name:  "no args return default",
			args:  args{dftB: true},
			wantB: true,
		},
		{
			name:  "with argsBool return first",
			args:  args{dftB: true, argsB: []bool{false}},
			wantB: false,
		},
		{
			name:  "no argsStr return default",
			args:  args{dftS: "default"},
			wantS: "default",
		},
		{
			name:  "with argsStr return first",
			args:  args{dftS: "default", argsS: []string{"first"}},
			wantS: "first",
		},
	}
	for _, tt := range tests {
		if funk.NotEmpty(tt.args.argsI) {
			got := xutil.FirstOrDefaultArgs(tt.args.dftI, tt.args.argsI...)
			s.Equal(tt.wantI, got, tt.name)
		} else {
			got := xutil.FirstOrDefaultArgs(tt.args.dftI)
			s.Equal(tt.wantI, got, tt.name)
		}

		if funk.NotEmpty(tt.args.argsB) {
			got := xutil.FirstOrDefaultArgs(tt.args.dftB, tt.args.argsB...)
			s.Equal(tt.wantB, got, tt.name)
		} else {
			got := xutil.FirstOrDefaultArgs(tt.args.dftB)
			s.Equal(tt.wantB, got, tt.name)
		}

		if funk.NotEmpty(tt.args.argsS) {
			got := xutil.FirstOrDefaultArgs(tt.args.dftS, tt.args.argsS...)
			s.Equal(tt.wantS, got, tt.name)
		} else {
			got := xutil.FirstOrDefaultArgs(tt.args.dftS)
			s.Equal(tt.wantS, got, tt.name)
		}

	}
}
