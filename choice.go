package xutil

// IfThen evaluates a condition, if true returns the parameters otherwise nil
func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElseStr(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

// DefaultIfNil checks if the value is nil, if true returns the default value otherwise the original
func DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

// FirstNonNil returns the first non nil parameter
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// AorB returns the first non zero value
func AorB[T General](a, b T) T {
	if !IsZeroVal(a) {
		return a
	}
	return b
}

// FirstNonZero returns the first non zero values,
// and if all are zero vlaues, return the first value in args
func FirstNonZero[T General](args ...T) T {
	for _, v := range args {
		if !IsZeroVal(v) {
			return v
		}
	}
	return args[0]
}

func IfaceAorB(a, b interface{}) interface{} {
	return IfThenElse(!IsZeroVal(a), a, b)
}

func FirstIface(args ...interface{}) interface{} {
	for _, v := range args {
		if !IsZeroVal(v) {
			return v
		}
	}
	return args[0]
}

// FirstOrDefaultArgs
//
// return the first args value, if args not empty
// else return default value
func FirstOrDefaultArgs[T General](dft T, args ...T) (val T) {
	val = dft
	if len(args) > 0 {
		val = args[0]
	}
	return val
}
