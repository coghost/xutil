package xutil

import (
	"time"

	"github.com/gookit/goutil/timex"
	"github.com/spf13/cast"
)

func ElapsedSeconds(start time.Time, args ...int) float64 {
	p := FirstOrDefaultArgs(0, args...)
	s := timex.ElapsedNow(start)
	f := cast.ToFloat64(s) / 1000

	if p == 0 {
		return f
	}
	return ToFixed(f, p)
}

// RandSleep random sleep some seconds
//
// @returns milliseconds
func RandSleep(min, max float64, msg ...string) int {
	slept := RandFloatX1k(min, max)
	time.Sleep(time.Duration(slept) * time.Millisecond)
	return slept
}
