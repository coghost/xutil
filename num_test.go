package xutil_test

import (
	"testing"
	"time"
	"xkit/xlog"

	"github.com/coghost/xutil"

	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/suite"
)

type NumSuite struct {
	suite.Suite
}

func TestNum(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(NumSuite))
}

func (s *NumSuite) SetupSuite() {
}

func (s *NumSuite) TearDownSuite() {
}

func (s *NumSuite) Test_ConstValue() {
	s.Equal(9223372036854775807, xutil.MaxInt)
	s.Equal(-9223372036854775808, xutil.MinInt)
	var b uint = 18446744073709551615
	s.Equal(b, xutil.MaxUint)
	s.Equal(0, xutil.MinUint)
}

func (s *NumSuite) TestElapsedSeconds() {
	st := time.Now()
	time.Sleep(time.Duration(1000) * time.Millisecond)
	v := xutil.ElapsedSeconds(st)
	dump.P(v)
	v1 := xutil.ToFixed(v, 0)
	s.Equal(1.0, v1)
	s.Greater(v, 1.0)
	s.Less(v, 1.1)
}

func (s *NumSuite) TestElapsedSecondsV1() {
	type args struct {
		dur  int
		prec int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "sleep 1 s",
			args: args{
				dur:  1020,
				prec: 2,
			},
			want: 1.02,
		},
	}
	for _, tt := range tests {
		st := time.Now()
		time.Sleep(time.Duration(tt.args.dur) * time.Millisecond)
		got := xutil.ElapsedSeconds(st, tt.args.prec)

		s.Equal(tt.want, got, tt.name)
	}
}

func (s *NumSuite) TestCeilInt() {
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
			name:    "99 / 25 is 4",
			args:    args{a: 99, b: 25},
			wantInt: 4,
		},
		{
			name:    "100 / 25 is 4",
			args:    args{a: 100, b: 25},
			wantInt: 4,
		},
		{
			name:    "101 / 25 is 5",
			args:    args{a: 101, b: 25},
			wantInt: 5,
		},
		{
			name:    "-1 / 25 is 0",
			args:    args{a: -1, b: 25},
			wantInt: 0,
		},
		{
			name:    "1 / 25 is 1",
			args:    args{a: 1, b: 25},
			wantInt: 1,
		},
	}
	for _, tt := range tests {
		got := xutil.CeilInt(tt.args.a, tt.args.b)
		s.Equal(tt.wantInt, got, tt.name)
	}
}

func (s *NumSuite) TestRound() {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "round 1.499 should be 1",
			args: args{
				num: 1.499,
			},
			want: 1,
		},
		{
			name: "round 1.5000001 should be 2",
			args: args{
				num: 1.5000001,
			},
			want: 2,
		},
		{
			name: "round -1.5000001 should be -2",
			args: args{
				num: -1.5000001,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		got := xutil.Round(tt.args.num)
		s.Equal(tt.want, got, tt.name)
	}
}

func (s *NumSuite) TestToFixed() {
	type args struct {
		num  float64
		prec int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "fixed 1.212 should be 1.21",
			args: args{
				num:  1.212,
				prec: 2,
			},
			want: 1.21,
		},
	}
	for _, tt := range tests {
		got := xutil.ToFixed(tt.args.num, tt.args.prec)
		s.Equal(tt.want, got, tt.name)
	}
}

func (s *NumSuite) Test_01_MaxMin() {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
		wantMin int
	}{
		{
			name: "min(1,21)=1, max(1,21)=21",
			args: args{
				arr: []int{1, 21},
			},
			wantMax: 21,
			wantMin: 1,
		},
		{
			name: "",
			args: args{
				arr: []int{21, 3},
			},
			wantMax: 21,
			wantMin: 3,
		},
	}
	for _, tt := range tests {
		got1 := xutil.Max(tt.args.arr[0], tt.args.arr[1])
		s.Equal(tt.wantMax, got1)

		got2 := xutil.Min(tt.args.arr[0], tt.args.arr[1])
		s.Equal(tt.wantMin, got2)
	}
}
