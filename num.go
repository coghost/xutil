package xutil

import (
	"math"
	"math/rand"

	"github.com/spf13/cast"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Max[T Number](args ...T) (max T) {
	max = args[0]
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func Min[T Number](args ...T) (min T) {
	min = args[0]
	for _, v := range args {
		if min > v {
			min = v
		}
	}
	return min
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision ...int) float64 {
	p := FirstOrDefaultArgs(2, precision...)
	output := math.Pow(10, float64(p))
	return float64(Round(num*output)) / output
}

func ToFixedStr(num float64, precision ...int) string {
	output := ToFixed(num, precision...)
	return cast.ToString(output)
}

// CeilInt:
//
//	@return int(Ceil(a / b))
func CeilInt(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

// DiffRate
//
//	@return float(a-b) / b
func DiffRate(a, b int) float64 {
	return math.Abs(float64(a-b)) / float64(b)
}

// RandFloatX1k:
// return a value between (min, max) * 1000
func RandFloatX1k(min, max float64) int {
	minI, maxI := int(math.Round(min*1000.0)), int(math.Round(max*1000.0))
	slept := minI + rand.Intn(maxI-minI)
	return slept
}
