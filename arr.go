package xutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// RefineIndex
//
//  1. 0 <= i < length return i
//  2. i < 0 return length+i
//  3. i >= length return length-1
func RefineIndex(i, length int) int {
	if i >= length {
		return length - 1
	} else if i < 0 {
		return length + i
	} else {
		return i
	}
}

// GetStrByIndex
//
// index can be less than 0
func GetStrByIndex(arr []string, index int) string {
	index = RefineIndex(index, len(arr))
	return arr[index]
}

// GetIntByIndex
//
// index can be less than 0
func GetIntByIndex(arr []int, index int) int {
	index = RefineIndex(index, len(arr))
	return arr[index]
}

// AnyJoin:
// join all args with sep, even value in args is zero
// if you don't want zero values, use `AnyJoinNon0` instead
func AnyJoin(sep string, args ...interface{}) string {
	arr := []string{}
	for _, v := range args {
		arr = append(arr, fmt.Sprintf("%v", v))
	}
	return strings.Join(arr, sep)
}

// AnyJoinNon0:
// all zero valid will be skipped
// but " " is allowed
func AnyJoinNon0(sep string, args ...interface{}) string {
	arr := []string{}
	for _, v := range args {
		if !IsZeroVal(v) {
			arr = append(arr, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(arr, sep)
}

// AnyJoinNon0ToUpper
//
//	@return string.ToUpper
func AnyJoinNon0ToUpper(sep string, args ...interface{}) string {
	s := AnyJoinNon0(sep, args...)
	return strings.ToUpper(s)
}

func NewSlice(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

func NewSliceByCount(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

// NewStringSlice
//
// raw: the raw string to be convert to slice
// fixStep: how many chars in each slice
// args: just add any one to enable random step mode
//
// return:
//
//	str slice with each has (~)maxLen chars
func NewStringSlice(raw string, fixStep int, randomStep ...bool) []string {
	rand.Seed(time.Now().Unix())
	step := fixStep
	var ret []string
	s := ""

	isRandStep := FirstOrDefaultArgs(false, randomStep...)

	for _, ch := range raw {
		s += string(ch)
		if len(s) >= step {
			ret = append(ret, s)
			s = ""

			if !isRandStep {
				continue
			}

			c := rand.Float64()
			if c > 0.8 {
				step = fixStep + 2
			} else if c > 0.6 {
				step = fixStep + 1
			} else if c < 0.2 {
				step = fixStep - 2
			} else if c < 0.4 {
				step = fixStep - 1
			} else {
				step = fixStep
			}
		}
	}

	if s != "" {
		ret = append(ret, s)
	}

	return ret
}

// Insert insert value to arr at index
//   - if index >= len(arr), append to arr
//   - else insert at index
func Insert[T General](arr []T, index int, value T) []T {
	// nil or empty slice or after last element
	if len(arr) <= index {
		return append(arr, value)
	}
	// index < len(a)
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value
	return arr
}

// GetStrBySplit: split raw str with separator and join from offset
//
//	example:
//	 raw = "a,b,c,d,e"
//	 v, b := GetStrBySplit(raw, ",", 1)
//	 // v = "bcde", b = true
//
//	 v, b := GetStrBySplit(raw, "_", 1)
//	 // v = "a,b,c,d,e", b = false
//
//	 v, b := GetStrBySplit(raw, ",", -1)
//	 // v = "e", b = false
//
// @return string
// @return bool
func GetStrBySplit(raw string, sep string, offset int) (string, bool) {
	if strings.Contains(raw, sep) {
		arr := strings.Split(raw, sep)
		i := offset
		if n := len(arr) - 1; n < offset {
			i = n
		}
		if offset < 0 {
			i = len(arr) + offset
		}
		return strings.Join(arr[i:], sep), true
	}
	return raw, false
}

// MustGetStrBySplit get str or ""
func MustGetStrBySplit(raw string, sep string, offset int) string {
	s, b := GetStrBySplit(raw, sep, offset)
	if b {
		return s
	}
	return ""
}

// GetStrBySplitAtIndex
// split raw to slice and then return element at index
//
//   - if sep not in raw, returns raw
//   - if index < 0, reset index to len() + index
//   - if index > total length, returns the last one
//   - else returns element at index
func GetStrBySplitAtIndex(raw interface{}, sep string, index int) string {
	str := cast.ToString(raw)
	if sep == "" || !strings.Contains(str, sep) {
		return str
	}

	arr := strings.Split(str, sep)
	if index > len(arr)-1 {
		index = len(arr) - 1
	} else if index < 0 {
		index = len(arr) + index
	}

	return arr[index]
}

func StrToArrWithNonEmpty(raw string, sep string) (arr []string) {
	for _, line := range strings.Split(raw, sep) {
		if v := strings.TrimSpace(line); v != "" {
			arr = append(arr, v)
		}
	}
	return
}

func ConcatSlice[T General](a []T, b []T) []T {
	return append(a, b...)
}
