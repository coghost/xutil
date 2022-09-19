package xutil_test

import (
	"reflect"
	"testing"
	"xkit/xlog"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
)

type ReflectSuite struct {
	suite.Suite
}

func TestReflect(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(ReflectSuite))
}

func (s *ReflectSuite) SetupSuite() {
}

func (s *ReflectSuite) TearDownSuite() {
}

func (s *ReflectSuite) TestCaller() {
	type args struct {
		lvl int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"default lvl 0", args{lvl: 0}, "xutil.Caller", 10},
		{"default lvl 1", args{lvl: 1}, "TestCaller", 10},
		{"default lvl 2", args{lvl: 5}, "testing.tRunner", 10},
	}
	for _, tt := range tests {
		got, got1 := xutil.Caller(tt.args.lvl)
		s.Contains(got, tt.want, tt.name)
		s.GreaterOrEqual(got1, tt.want1, tt.name)
	}

	xutil.DumpCallerStack()
}

var zeroTests = []struct {
	name string
	arg  interface{}
	want bool
}{
	{"byte", byte(0), true},
	{"int", 0, true},
	{"float", 0.0, true},
	{"string", "", true},
	{"bool", false, true},
}

func (s *ReflectSuite) TestIsZeroVal() {
	for _, tt := range zeroTests {
		b := xutil.IsZeroVal(tt.arg)
		s.Equal(tt.want, b, tt.name)
	}
}

func (s *ReflectSuite) TestIsDefaultVal() {
	for _, tt := range zeroTests {
		b := xutil.IsDefaultVal(tt.arg)
		s.Equal(tt.want, b, tt.name)
	}
}

type mockStruct struct {
	// basic
	Male   bool
	Age    int
	Height float64
	Weight float64
	Name   string

	Dow  []byte
	Ui64 uint64
	I64  int64

	Favor    []string
	LuckyNum []int
}

func (s *ReflectSuite) TestGetAttr() {
	ms := mockStruct{
		Male:     false,
		Age:      0,
		Height:   0,
		Weight:   0,
		Name:     "",
		Dow:      []byte{},
		Ui64:     0,
		I64:      0,
		Favor:    []string{},
		LuckyNum: []int{},
	}

	type args struct {
		o   interface{}
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{"default", args{o: &ms, key: "Age"}},
	}

	v := reflect.Value{}

	for _, tt := range tests {
		v1 := xutil.GetAttr(tt.args.o, tt.args.key)
		s.IsType(v, v1, tt.name)
	}
}

func (s *ReflectSuite) TestSetKV() {
	ms := mockStruct{
		Male:     false,
		Age:      0,
		Height:   0.0,
		Weight:   0.0,
		Name:     "",
		Dow:      []byte{},
		Ui64:     0,
		I64:      0,
		Favor:    []string{},
		LuckyNum: []int{},
	}

	type args struct {
		o   interface{}
		key string
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Male", args{o: &ms, key: "Male", val: true}, true},
		{"Age", args{o: &ms, key: "Age", val: 10}, 10},
		{"Height", args{o: &ms, key: "Height", val: 10.01}, 10.01},
		{"Weight", args{o: &ms, key: "Weight", val: 50.4}, 50.4},
		{"Name", args{o: &ms, key: "Name", val: "FakeName"}, "FakeName"},
		{"Dow", args{o: &ms, key: "Dow", val: []byte{1}}, []byte{1}},
		{"Ui64", args{o: &ms, key: "Ui64", val: uint64(10)}, uint64(10)},
		{"I64", args{o: &ms, key: "I64", val: int64(10)}, int64(10)},
		{"Favor", args{o: &ms, key: "Favor", val: []string{"001", "002"}}, []string{"001", "002"}},
	}

	for _, tt := range tests[:len(tests)-1] {
		xutil.SetAttr(tt.args.o, tt.args.key, tt.args.val)
		switch tt.args.key {
		case "Male":
			s.Equal(tt.want, ms.Male, tt.name)
		case "Age":
			s.Equal(tt.want, ms.Age, tt.name)
		case "Height":
			s.Equal(tt.want, ms.Height, tt.name)
		case "Weight":
			s.Equal(tt.want, ms.Weight, tt.name)
		case "Name":
			s.Equal(tt.want, ms.Name, tt.name)
		case "Dow":
			s.Equal(tt.want, ms.Dow, tt.name)
		case "Ui64":
			s.Equal(tt.want, ms.Ui64, tt.name)
		case "I64":
			s.Equal(tt.want, ms.I64, tt.name)
		}
	}

	for _, tt := range tests[len(tests)-1:] {
		s.Panics(func() {
			xutil.SetAttr(tt.args.o, tt.args.key, tt.args.val)
		})
	}
}
