package xutil_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/coghost/xutil"

	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
)

type DtmSuite struct {
	suite.Suite
}

func TestDtm(t *testing.T) {
	suite.Run(t, new(DtmSuite))
}

func (s *DtmSuite) SetupSuite() {
}

func (s *DtmSuite) TearDownSuite() {
}

func (s *DtmSuite) Test_01_Now() {
	now := xutil.Now()

	v := now.ToISOString()
	pp.Println(v)

	s.Contains(v, "T")
	_, off := now.ToTime().Zone()
	exp := fmt.Sprintf("+%02d:00", off/3600)
	s.Contains(v, exp, "timezone should match")
}

func (s *DtmSuite) Test_02_StrNow() {
	n := xutil.StrNow()
	s.NotEmpty(n)
	s.True(true)

	n1 := xutil.StrNow("YYYY")
	s.NotEmpty(n)

	s.Equal(len(n1), 4)
}

func (s *DtmSuite) Test_03_PythonTimeTime() {
	// t and t1 should almost equal
	t := xutil.PythonTimeTime(-3600)
	t1 := xutil.PythonTimeTime(0) - 3600
	fmt.Printf("%#v\n", cast.ToString(t))
	fmt.Printf("%#v\n", cast.ToString(t1))
	s.LessOrEqual(t1-t, 1e-5)
}

func (s *DtmSuite) Test_04_UTCNow() {
	g := xutil.UTCNow()
	s.NotNil(g)
	v := xutil.UTCNow().ToISOString()
	pp.Println("got", v)
}

func (s *DtmSuite) Test_05_Unix2Str() {
	n := "1634183927"
	got1 := xutil.Unix2Str(n)
	s.Equal("2021-10-14 11:58:47", got1)

	got2, err := xutil.UnixStr(n)
	s.Equal("2021-10-14T11:58:47+08:00", got2)
	s.Nil(err)

	got11 := xutil.Str2Unix(got1)
	s.NotEqual(cast.ToInt64(n), got11)
	got12 := xutil.Str2UnixWithAutoZone(got1)
	s.Equal(cast.ToInt64(n), got12)

	got21 := xutil.Str2Unix(got2, time.RFC3339)
	s.Equal(cast.ToInt64(n), got21)

	got22 := xutil.Str2UnixWithAutoZone(got2, time.RFC3339)
	s.Equal(cast.ToInt64(n), got22)

	// got = xutil.Unix2Str(n, "20060102150405")
	// s.Equal("20211014115847", got)
}

func (s *DtmSuite) Test_06_PythonTimeTimeAll() {
	offset := -3600

	now := time.Now()
	orig := now.UnixNano()

	now = now.Add(time.Duration(offset) * time.Second)
	m1 := now.UnixMicro()
	ml1 := now.UnixMilli()
	mn1 := now.UnixNano()

	g := xutil.Now()
	g1 := g.Add(time.Duration(offset) * time.Second).ToTime()
	m2 := g1.UnixMicro()
	ml2 := g1.UnixMilli()
	mn2 := g1.UnixNano()

	f1 := cast.ToFloat64(m1) / 1000000.0
	fl1 := cast.ToFloat64(ml1) / 1000.0
	fn1 := cast.ToFloat64(mn1) / 1000000000.0
	f2 := cast.ToFloat64(m2) / 1000000.0
	fl2 := cast.ToFloat64(ml2) / 1000.0
	fn2 := cast.ToFloat64(mn2) / 1000000000.0

	o1 := cast.ToFloat64(orig) / 1000000000.0

	fmt.Printf("%v, %v, %v\n", cast.ToInt64(f1), cast.ToInt64(fl1), cast.ToInt(fn1))
	fmt.Printf("%v, %v, %v\n", cast.ToInt64(f2), cast.ToInt64(fl2), cast.ToInt64(fn2))

	s.Equal(cast.ToInt64(f1), cast.ToInt64(fl1))
	s.Equal(cast.ToInt64(f1), cast.ToInt64(fn1))
	s.Equal(cast.ToInt64(f2), cast.ToInt64(fl2))
	s.Equal(cast.ToInt64(f2), cast.ToInt64(fn2))

	s.Equal(cast.ToInt64(f1), cast.ToInt64(f2))
	s.Equal(cast.ToInt64(fl1), cast.ToInt64(fl2))
	s.Equal(cast.ToInt64(fn1), cast.ToInt64(fn2))

	s.Equal(cast.ToInt64(o1), cast.ToInt64(f1)+3600)
}
