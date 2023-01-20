package xutil

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

var ErrorNoNumbers = errors.New("no number found")

type NumOpts struct {
	chars string
	dft   interface{}
}

type NumOptFunc func(o *NumOpts)

func Chars(s string) NumOptFunc {
	return func(o *NumOpts) {
		o.chars = s
	}
}

func Dft(i interface{}) NumOptFunc {
	return func(o *NumOpts) {
		o.dft = i
	}
}

func bindOpts(opt *NumOpts, opts ...NumOptFunc) {
	for _, f := range opts {
		f(opt)
	}
}

// CharToNum: substract `number+Chars` from source str
// returns int by default
func CharToNum(s string, opts ...NumOptFunc) (v interface{}, e error) {
	opt := NumOpts{chars: ".", dft: 1.0}
	bindOpts(&opt, opts...)

	a := "[0-9" + opt.chars + "]+"
	re := regexp.MustCompile(a)
	c := re.FindAllString(s, -1)
	r := strings.Join(c, "")

	if r == "" {
		return r, ErrorNoNumbers
	}

	switch opt.dft.(type) {
	case int:
		// v could be float
		v, e := cast.ToFloat64E(r)
		if e != nil {
			return nil, e
		}
		return cast.ToIntE(v)
	case int64:
		// v could be float
		v, e := cast.ToFloat64E(r)
		if e != nil {
			return nil, e
		}
		return cast.ToInt64E(v)
	case float32:
		return cast.ToFloat32E(r)
	case float64:
		return cast.ToFloat64E(r)
	default:
		return cast.ToStringE(r)
	}
}

func MustCharToNum(s string, opts ...NumOptFunc) (v interface{}) {
	v, e := CharToNum(s, opts...)
	PanicIfErr(e)
	return v
}

// GetCharNumOr returns parsed number or keep original
func GetCharNumOr(s string, opts ...NumOptFunc) (v interface{}) {
	v, e := CharToNum(s, opts...)
	if e == nil {
		return v
	}
	return s
}

func F64KMFromStr(str string, opts ...NumOptFunc) (i float64, err error) {
	unit := 1.0

	if strings.Contains(strings.ToUpper(str), "K") {
		unit = 1000.0
	}
	if strings.Contains(strings.ToUpper(str), "M") {
		unit = 1000000.0
	}

	opt := NumOpts{chars: ".", dft: 1.0}
	bindOpts(&opt, opts...)

	if !strings.Contains(opt.chars, ".") {
		opt.chars += "."
	}

	v, err := CharToNum(str, Chars(opt.chars), Dft(opt.dft))
	if err != nil {
		return
	}
	return cast.ToFloat64(v) * unit, err
}

func MustF64KMFromStr(str string, opts ...NumOptFunc) float64 {
	if v, err := F64KMFromStr(str, opts...); err != nil {
		panic(fmt.Sprintf("no number found in %s, %v", str, err))
	} else {
		return v
	}
}

func MustIntKMFromStr(str string, opts ...NumOptFunc) int {
	v := MustF64KMFromStr(str, opts...)
	return cast.ToInt(cast.ToFloat64(v))
}

func IntKMFromStr(str string, opts ...NumOptFunc) (int, bool) {
	if v, err := F64KMFromStr(str, opts...); err != nil {
		return 0, false
	} else {
		return cast.ToInt(cast.ToFloat64(v)), true
	}
}
