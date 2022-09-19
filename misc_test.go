package xutil_test

import (
	"testing"
	"time"
	"xkit/xlog"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
)

type MiscSuite struct {
	suite.Suite
}

func TestMisc(t *testing.T) {
	xlog.InitLog()
	suite.Run(t, new(MiscSuite))
}

func (s *MiscSuite) SetupSuite() {
}

func (s *MiscSuite) TearDownSuite() {
}

func (s *MiscSuite) TestRandFloatX1k() {
	v := xutil.RandFloatX1k(0.1, 0.101)
	s.Equal(100, v)
	s.LessOrEqual(100, v)
	s.GreaterOrEqual(101, v)
}

func (s *MiscSuite) TestRandSleep() {
	ts := time.Now()

	xutil.RandSleep(0.3, 1.1)

	end := int(time.Since(ts).Milliseconds())

	s.LessOrEqual(end, 1100)
	s.GreaterOrEqual(end, 300)
}

func (s *MiscSuite) TestIsJSON() {
	tests := []struct {
		expected bool
		err      error
		raw      string
	}{
		{raw: ` [{"a":1}, {"a": 2}] `, expected: true},
		{raw: ` {"a":{"b": 2}} `, expected: true},
		{raw: ` [] `, expected: true},
		{raw: ` {} `, expected: true},
		{raw: ` {{"a": 1}} `, expected: false},
	}

	for _, test := range tests {
		v := xutil.IsJSON(test.raw)
		s.Equal(test.expected, v)
	}
}

func (s *MiscSuite) TestGetIp() {
	xutil.GetPublicIp()
}

func (s *MiscSuite) TestGetHostIp() {
	s.Panics(func() {
		xutil.GetHostPublicIp("")
	}, "empty url should panic")
}

func (s *MiscSuite) Test_01_PTD() {
	xutil.InitializeXOpts(xutil.SetDebug(true))
	xutil.PauseToDebug("whoami")
}
