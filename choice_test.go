package xutil_test

import (
	"testing"
	"xkit/xlog"

	"github.com/coghost/xutil"
	"github.com/stretchr/testify/suite"
)

type ChoiceSuite struct {
	suite.Suite
}

func TestChoice(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(ChoiceSuite))
}

func (s *ChoiceSuite) SetupSuite() {
}

func (s *ChoiceSuite) TearDownSuite() {
}

func (s *ChoiceSuite) Test_02_IfaceAorB() {
	s.Equal(0.1, xutil.FirstNonZero(0.00, 0.1))
	s.Equal(1, xutil.FirstNonZero(1, 10))

	s.Equal(0.1, xutil.IfaceAorB(0.00, 0.1).(float64))
	s.Equal(1, xutil.IfaceAorB(1, 10).(int))
	s.Equal(10, xutil.IfaceAorB("", 10).(int))

	s.Equal(10, xutil.FirstIface(0, 10).(int))
}

func (s *ChoiceSuite) Test_0301_FirstNonZero() {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantInt int
	}{
		{
			name:    "IntAorB(0, 3) = 3",
			args:    args{a: 0, b: 3},
			wantInt: 3,
		},
		{
			name:    "IntAorB(7, 3) = 7",
			args:    args{a: 7, b: 3},
			wantInt: 7,
		},
	}
	for _, tt := range tests {
		got := xutil.FirstNonZero(tt.args.a, tt.args.b)
		s.Equal(tt.wantInt, got, tt.name)
	}
}

func (s *ChoiceSuite) Test_0302_FirstNonZero() {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `FirstNonZero("", "none") = none`,
			args: args{a: "", b: "none"},
			want: "none",
		},
		{
			name: `FirstNonZero("ok", "none") = none`,
			args: args{a: "ok", b: "none"},
			want: "ok",
		},
	}
	for _, tt := range tests {
		got := xutil.FirstNonZero(tt.args.a, tt.args.b)
		s.Equal(tt.want, got, tt.name)
	}
}

func (s *ChoiceSuite) TestFirstNonEmptyInt() {
	got := xutil.FirstNonZero(0, 0, 2, 4)
	s.Equal(2, got)
	got1 := xutil.FirstNonZero(0, 0, 0, 0, 0)
	s.Equal(0, got1)
}
