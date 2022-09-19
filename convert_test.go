package xutil_test

import (
	"testing"
	"xkit/xlog"

	"github.com/coghost/xutil"
	"github.com/stretchr/testify/suite"
)

type ConvertSuite struct {
	suite.Suite
}

func TestConvert(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(ConvertSuite))
}

func (s *ConvertSuite) SetupSuite() {
}

func (s *ConvertSuite) TearDownSuite() {
}

func (s *ConvertSuite) Test_01_AnB() {
	type args struct {
		src string
		dft interface{}
	}
	tests := []struct {
		name  string
		args  args
		want1 interface{}
		want2 interface{}
	}{
		{
			name:  "is str 12",
			args:  args{src: "12", dft: ""},
			want1: "12",
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is int 12",
			args:  args{src: "12", dft: 0},
			want1: 12,
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is int64(12)",
			args:  args{src: "12", dft: int64(0)},
			want1: int64(12),
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is float64(12)",
			args:  args{src: "12", dft: 0.0},
			want1: 12.0,
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is float32(12)",
			args:  args{src: "12", dft: float32(0)},
			want1: float32(12),
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is byte(12)",
			args:  args{src: "12", dft: []byte{0}},
			want1: []byte{0x31, 0x32},
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is int(-12)",
			args:  args{src: "-12", dft: -1},
			want1: -12,
			want2: []byte{0x2d, 0x31, 0x32},
		},
		{
			name:  "is str(12)",
			args:  args{src: "12", dft: int16(0)},
			want1: "12",
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is str(12)",
			args:  args{src: "12", dft: nil},
			want1: "12",
			want2: []byte{0x31, 0x32},
		},
		{
			name:  "is true",
			args:  args{src: "12", dft: true},
			want1: true,
			want2: []byte{0x74, 0x72, 0x75, 0x65},
		},
	}
	for _, tt := range tests {
		got1 := xutil.B2A([]byte(tt.args.src), tt.args.dft)
		s.Equal(tt.want1, got1, tt.name)

		got2 := xutil.A2B(got1)
		s.Equal(tt.want2, got2)
	}

	got := xutil.B2A([]byte("12"))
	s.Equal("12", got)
}

func (s *ConvertSuite) Test_02_ArrAnB() {
	type args struct {
		arr []string
	}
	tests := []struct {
		name   string
		args   args
		wantSA []interface{}
		wantBA [][]byte
	}{
		{
			name: "string byte array convertion",
			args: args{
				arr: []string{"first", "second", "191", "true"},
			},
			wantSA: []interface{}{"first", "second", "191", "true"},
			wantBA: [][]uint8{{0x66, 0x69, 0x72, 0x73, 0x74}, {0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64}, {0x31, 0x39, 0x31}, {0x74, 0x72, 0x75, 0x65}},
		},
	}
	for _, tt := range tests {
		got := xutil.ArrS2B(tt.args.arr...)
		s.Equal(tt.wantBA, got)

		got1 := xutil.ArrB2A(got...)
		s.Equal(tt.wantSA, got1)
	}

	b1 := [][]byte{{}, nil}
	got := xutil.ArrB2A(b1...)
	s.Equal([]interface{}{"", interface{}(nil)}, got)
}

type Person struct {
	Name  string
	Age   int
	Other map[string]interface{} `structs:",remain"`
}

func (s *ConvertSuite) Test_03_Map2Struct() {
	input := map[string]interface{}{
		"name":  "Mitchell",
		"age":   91,
		"email": "mitchell@example.com",
	}

	var p Person
	xutil.Map2StructureByDefault(input, &p)

	s.Equal(91, p.Age)
	s.Equal("Mitchell", p.Name)
	s.Equal("mitchell@example.com", p.Other["email"])
}

func (s *ConvertSuite) Test_04_Structify() {
	dat := `{"Name":"Rick","Age":100}`
	p := &Person{}
	xutil.MustStructify(dat, p)
	p1 := &Person{"Rick", 100, nil}
	s.Equal(p1, p)
}

func (s *ConvertSuite) Test_05_Stringify() {
	p := &Person{
		Name: "Rick",
		Age:  100,
	}
	str := xutil.MustStringify(p)
	st1 := `{"Name":"Rick","Age":100,"Other":null}`
	s.Equal(st1, str)
}
