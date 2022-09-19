package xutil

import (
	"math"
	"time"

	"github.com/nleeper/goment"
	"github.com/spf13/cast"
)

// WARN: this will be replaced by xdtm in the future

const (
	DFmt  = "YYYY-MM-DD HH:mm:ss"
	DFmt1 = "YYYYMMDDHHmmss"
	DFmt2 = "HH:mm:ss"
	DFDay = "YYYYMMDD"

	// fmt with zone
	ZoneFmt = "YYYY-MM-DDTHH:mm:ssZ"

	DFYMDhms = "2006-01-02 15:04:05"
)

// StrNow returns datetime with format `YYYY-MM-DD HH:mm:ss`
func StrNow(opts ...string) string {
	fmt := DFmt
	if len(opts) > 0 {
		fmt = opts[0]
	}

	return Now().Format(fmt)
}

// UTCStrNow
func UTCStrNow(opts ...string) string {
	fmt := ZoneFmt
	if len(opts) > 0 {
		fmt = opts[0]
	}
	return UTCNow().Format(fmt)
}

func Now() *goment.Goment {
	g, _ := goment.New()
	return g
}

func UTCNow() *goment.Goment {
	return Now().UTC()
}

// PythonTimeTime
//
// return same format with python's time.time() `1234567890.123456`
func PythonTimeTime(offsetSeconds int64) float64 {
	now := time.Now()
	now = now.Add(time.Duration(offsetSeconds) * time.Second)
	m := now.UnixNano()
	sameAsPyTime := cast.ToFloat64(m) / math.Pow10(9)
	return sameAsPyTime
}

// Unix2Str
//
// returns "2006-01-02 15:04:05" by default
func Unix2Str(dtm interface{}, fmtArgs ...string) string {
	fmt := DFYMDhms
	if len(fmtArgs) > 0 {
		fmt = fmtArgs[0]
	}
	g, e := goment.Unix(cast.ToInt64(dtm))
	PanicIfErr(e)
	return g.ToTime().Format(fmt)
}

func UnixStr(dtm interface{}, layoutArgs ...string) (str string, err error) {
	it, err := cast.ToInt64E(dtm)
	if err != nil {
		return "", err
	}

	ut := time.Unix(it, 0)

	layout := FirstOrDefaultArgs(time.RFC3339, layoutArgs...)
	str = ut.Format(layout)
	return
}

// Str2UnixWithAutoZone converts value to unix timestamp
//
// if there is no zone in value, will auto add/sub the time zone
// use with caution
func Str2UnixWithAutoZone(value string, layoutArgs ...string) int64 {
	layout := FirstOrDefaultArgs(DFYMDhms, layoutArgs...)
	t, err := time.Parse(layout, value)
	PanicIfErr(err)

	name, offset := t.Zone()
	tz := offset / 3600
	if name == "UTC" && tz == 0 {
		_, offset := time.Now().Zone()
		lz := offset / 3600
		t = t.Add(time.Hour * -time.Duration(lz))
	}

	return t.Unix()
}

// Str2Unix
func Str2Unix(value string, layoutArgs ...string) int64 {
	layout := FirstOrDefaultArgs(DFYMDhms, layoutArgs...)
	t, err := time.Parse(layout, value)
	PanicIfErr(err)
	return t.Unix()
}
