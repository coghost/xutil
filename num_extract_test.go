package xutil_test

import (
	"errors"
	"strconv"

	"github.com/coghost/xutil"

	"github.com/k0kubun/pp/v3"
)

func (s *NumSuite) TestCharToNumNonOpts() {
	tests := []struct {
		txt      string
		custom   string
		expected interface{}
		err      error
	}{
		{txt: "thisisok", expected: "", err: xutil.ErrorNoNumbers},
		{txt: "0this0iso0k", expected: 0, err: nil},

		{txt: "10thisiso0k", expected: 100, err: nil},
		{txt: "10thisiso0k", expected: 100.0, err: nil},
		{txt: "10thisiso0k", expected: "100", err: nil},

		// convert to default int will fail
		{txt: "0thisiso.0k", custom: ".", expected: 0, err: nil},
		{txt: "0thisiso.0k", custom: ".", expected: "0.0", err: nil},
		{txt: "0thisiso.0k", custom: ".", expected: 0.0, err: nil},

		// convert to default int will fail
		{txt: "0thisiso.01k", custom: ".", expected: 0, err: nil},
		{txt: "0thisiso.01k", custom: ".", expected: 0.01, err: nil},

		{txt: "-0this iso.01k", custom: ".-", expected: -0.01, err: nil},
	}

	for _, test := range tests {
		// fmt.Println(test)
		v, err := xutil.CharToNum(test.txt, xutil.Chars(test.custom), xutil.Dft(test.expected))
		if errors.Is(test.err, strconv.ErrSyntax) {
			s.ErrorIs(err, test.err)
			continue
		}
		if err == xutil.ErrorNoNumbers {
			s.Equal(test.err, err)
			continue
		}
		s.Equal(test.expected, v)
		s.Equal(test.err, err)
	}
}

func (s *NumSuite) TestCharToNumWithDefault() {
	tests := []struct {
		txt    string
		custom string
		want   interface{}
		err    error
	}{
		{txt: "0thisiso.1k", custom: ".", want: float32(0.1), err: nil},
		{txt: "0thisiso.1k", custom: ".", want: 0.1, err: nil},
		{txt: "0thisiso.01k", custom: "", want: 1, err: nil},
		{txt: "0thisiso.01k", custom: "", want: int64(1), err: nil},
		{txt: "0thisiso.01k", custom: "", want: "001", err: nil},
		{txt: "thisiso.k", custom: "", want: "", err: xutil.ErrorNoNumbers},
		{
			txt: `Top Rated Seller
(222)`,
			custom: "",
			want:   222,
			err:    nil,
		},
	}
	for _, test := range tests {
		v, err := xutil.CharToNum(test.txt, xutil.Chars(test.custom), xutil.Dft(test.want))
		s.Equal(test.want, v)
		s.Equal(test.err, err)
	}
}

func (s *NumSuite) TestMustCharToNum() {
	raw := "12.1K"
	var dft interface{} = 0.01
	v := xutil.MustCharToNum(raw, xutil.Chars("."), xutil.Dft(dft))
	s.Equal(12.1, v)
}

func (s *NumSuite) Test_01_IntK() {
	raw := "1.8K+"
	raw = "5.3M+"
	v := xutil.MustIntKMFromStr(raw, xutil.Chars(""), xutil.Dft(0.0))
	pp.Println(v)
}

func (s *NumSuite) Test03_NumF64() {
	s1 := "90.6K"
	v := xutil.MustCharToNum(s1)
	pp.Println(v)
	f, b := xutil.F64KMFromStr(s1)
	s.True(b)
	pp.Println(f)

	i := xutil.MustIntKMFromStr(s1)
	s.Equal(90600, i)
}
