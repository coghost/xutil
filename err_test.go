package xutil_test

import (
	"errors"
	"testing"
	"xkit/xlog"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
)

type ErrSuite struct {
	suite.Suite
}

func TestErr(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(ErrSuite))
}

func (s *ErrSuite) SetupSuite() {
}

func (s *ErrSuite) TearDownSuite() {
}

func (s *ErrSuite) TestEnsureByRetry() {
	efn := func() error {
		return errors.New("this is error func")
	}

	type args struct {
		fn   func() error
		args []int
	}
	tests := []struct {
		name      string
		args      args
		wantTried int
		wantErr   bool
		wantStr   string
	}{
		{
			name: "want -1, delay 1000, show 1",
			args: args{
				fn:   efn,
				args: []int{3, 1000, 1},
			},
			wantTried: 3,
			wantErr:   true,
			wantStr:   "All attempts fail:\n#1: this is error func\n#2: this is error func\n#3: this is error func",
		},
	}

	for _, tt := range tests {
		got, err := xutil.EnsureByRetry(tt.args.fn, tt.args.args...)
		s.Equal(tt.wantTried, got, "directly call")
		s.EqualError(err, tt.wantStr, "directly call")
	}
}

func (s *ErrSuite) TestCatchPanicAsErr() {
	pfn := func() {
		panic("test of panic")
	}

	e := xutil.CatchPanicAsErr(pfn)
	s.NotNil(e)
	s.ErrorIs(e, &xutil.ErrTry{})

	s.Equal(`error value: "test of panic"`, e.Error())
}

func (s *ErrSuite) TestPanicIfErr() {
	e := errors.New("this is error func")
	s.Panics(func() {
		xutil.PanicIfErr(e)
	})
}
