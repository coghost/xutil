package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
)

type DummySuite struct {
	suite.Suite
}

func TestDummy(t *testing.T) {
	suite.Run(t, new(DummySuite))
}

func (s *DummySuite) SetupSuite() {
}

func (s *DummySuite) TearDownSuite() {
}

func (s *DummySuite) TestDumbLog() {
	type args struct {
		msg []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"empty", args{msg: []interface{}{}}},
		{"empty str", args{msg: []interface{}{""}}},
		{"str int", args{msg: []interface{}{"abc", 123}}},
	}
	for _, tt := range tests {
		xutil.DummyLog(tt.args.msg...)
		s.True(true)
	}

	for _, tt := range tests {
		xutil.DummyLog(tt.args.msg...)
		s.True(true)
	}
}
