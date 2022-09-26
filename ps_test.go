package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/suite"
)

type XpsSuite struct {
	suite.Suite
}

func TestXps(t *testing.T) {
	suite.Run(t, new(XpsSuite))
}

func (s *XpsSuite) SetupSuite() {
}

func (s *XpsSuite) TearDownSuite() {
}

func (s *XpsSuite) Test_01() {
	// ps := xps.NewPsMetric()
	// dump.P(xutil.MustStringify(ps))

	dump.P(xutil.HostUUID())
	dump.P(xutil.GetLocalIpAddr())
}
