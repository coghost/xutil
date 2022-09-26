package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
)

type ChaosSuite struct {
	suite.Suite
}

func TestChaos(t *testing.T) {
	suite.Run(t, new(ChaosSuite))
}

func (s *ChaosSuite) SetupSuite() {
}

func (s *ChaosSuite) TearDownSuite() {
}

func (s *ChaosSuite) Test_01_enc() {
	key := "aaaaaaaaaaaaaaaa"
	iv := "bbbbbbbbbbbbbbbb"
	iv2 := "cccccccccccccccc"
	chaos := xutil.NewChaos(key, iv)
	raw := "plaintext"
	v := chaos.Encrypt(raw)
	s.Equal("8a67020d988f9b009318d5c1dd429b4c", v)
	v1 := chaos.Decrypt(v)
	s.Equal(raw, v1)

	c2 := xutil.NewChaos(key, iv2)
	v2 := c2.Decrypt(v)
	s.NotEqual(raw, v2)
}
