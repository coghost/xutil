package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"
	"github.com/stretchr/testify/suite"
)

type ArrSuite struct {
	suite.Suite
}

func TestArr(t *testing.T) {
	suite.Run(t, new(ArrSuite))
}

func (s *ArrSuite) SetupSuite() {
}

func (s *ArrSuite) TearDownSuite() {
}

func (s *ArrSuite) TestAnyJoin() {
	type args struct {
		arr []interface{}
		sep string
	}
	tests := []struct {
		name  string
		args  args
		want1 string
		want2 string
		want3 string
	}{
		{
			name: "all str join with empty and with sep=,",
			args: args{
				arr: []interface{}{"fir", "sec", "", "ddd"},
				sep: ",",
			},
			want1: "fir,sec,,ddd",
			want2: "fir,sec,ddd",
			want3: "FIR,SEC,DDD",
		},
		{
			name:  "all str join without empty and with sep=,",
			args:  args{arr: []interface{}{"fir", "sec", "0", "ddd"}, sep: ","},
			want1: "fir,sec,0,ddd",
			want2: "fir,sec,0,ddd",
			want3: "FIR,SEC,0,DDD",
		},
		{
			name: "any join with 0 value",
			args: args{
				arr: []interface{}{10, "ok", 0, true, "whodoyou", 1.01, false, "false", 0.0, "0.0"},
				sep: "-",
			},
			want1: "10-ok-0-true-whodoyou-1.01-false-false-0-0.0",
			want2: "10-ok-true-whodoyou-1.01-false-0.0",
			want3: "10-OK-TRUE-WHODOYOU-1.01-FALSE-0.0",
		},
	}
	for _, tt := range tests {
		got1 := xutil.AnyJoin(tt.args.sep, tt.args.arr...)
		s.Equal(tt.want1, got1, tt.name)
		got2 := xutil.AnyJoinNon0(tt.args.sep, tt.args.arr...)
		s.Equal(tt.want2, got2, tt.name)
		got3 := xutil.AnyJoinNon0ToUpper(tt.args.sep, tt.args.arr...)
		s.Equal(tt.want3, got3, tt.name)
	}
}

func (s *ArrSuite) TestArrGet() {
	type args struct {
		strArr []string
		intArr []int
		index  int
	}
	arr1 := []string{"first", "second", "third"}
	arr2 := []int{28, 93, 7}
	tests := []struct {
		name    string
		args    args
		wantStr string
		wantInt int
	}{
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  0,
			},
			wantStr: "first",
			wantInt: 28,
		},
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  -1,
			},
			wantStr: "third",
			wantInt: 7,
		},
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  100,
			},
			wantStr: "third",
			wantInt: 7,
		},
	}
	for _, tt := range tests {
		got1 := xutil.GetStrByIndex(tt.args.strArr, tt.args.index)
		s.Equal(tt.wantStr, got1, tt.name)

		got2 := xutil.GetIntByIndex(tt.args.intArr, tt.args.index)
		s.Equal(tt.wantInt, got2, tt.name)
	}
}

func (s *ArrSuite) TestNewSlice() {
	type args struct {
		start int
		end   int
		count int
		step  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				start: 1,
				end:   10,
				count: 10,
				step:  1,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "",
			args: args{
				start: 1,
				end:   10,
				count: 5,
				step:  2,
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "",
			args: args{
				start: 1,
				end:   100,
				count: 10,
				step:  10,
			},
			want: []int{1, 11, 21, 31, 41, 51, 61, 71, 81, 91},
		},
		{
			name: "",
			args: args{
				start: -10,
				end:   -1,
				count: 5,
				step:  2,
			},
			want: []int{-10, -8, -6, -4, -2},
		},
		{
			name: "",
			args: args{
				start: 1,
				end:   10,
				count: 0,
				step:  -1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		got := xutil.NewSlice(tt.args.start, tt.args.end, tt.args.step)
		s.Equal(tt.want, got)

		got1 := xutil.NewSliceByCount(tt.args.start, tt.args.count, tt.args.step)
		s.Equal(tt.want, got1)
	}
}

func (s *ArrSuite) TestStringSlice() {
	raw := "copy editor TestArr/TestStringSlice"
	got := xutil.NewStringSlice(raw, 4)
	s.Equal([]string{"copy", " edi", "tor ", "Test", "Arr/", "Test", "Stri", "ngSl", "ice"}, got)

	raw = `
func (s *ArrSuite) TestStringSlice() {
	raw := "copy editor TestArr/TestStringSlice"
	got := xutil.NewStringSlice(raw, 4)
	s.Equal([]string{"copy", " edi", "tor ", "Test", "Arr/", "Test", "Stri", "ngSl", "ice"}, got)
	`
	got1 := xutil.NewStringSlice(raw, 4, true)
	s.IsType([]string{}, got1)
	// maximum random length is 4(step)+2
	s.Greater(len(got1), len(raw)/6)
}

func (s *ArrSuite) Test_01_Insert() {
	type args struct {
		strArr []string
		intArr []int
		index  int
	}

	arr1 := []string{"first", "second", "third"}
	arr2 := []int{28, 93, 7}
	st1 := "abc"
	int1 := 100

	tests := []struct {
		name  string
		args  args
		want1 []string
		want2 []int
	}{
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  0,
			},
			want1: []string{st1, "first", "second", "third"},
			want2: []int{int1, 28, 93, 7},
		},
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  2,
			},
			want1: []string{"first", "second", st1, "third"},
			want2: []int{28, 93, int1, 7},
		},
		{
			name: "",
			args: args{
				strArr: arr1,
				intArr: arr2,
				index:  20,
			},
			want1: []string{"first", "second", "third", st1},
			want2: []int{28, 93, 7, int1},
		},
	}
	for _, tt := range tests {
		a1 := xutil.Insert(tt.args.strArr, tt.args.index, st1)
		a2 := xutil.Insert(tt.args.intArr, tt.args.index, int1)
		s.Equal(tt.want1, a1)
		s.Equal(tt.want2, a2)
	}
}
