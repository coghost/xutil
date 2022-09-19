package xutil

// refer: https://mp.weixin.qq.com/s/QBZ1dp0XIqMo24vVFYf1fA

type UInts interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}

type Floats interface {
	float32 | float64
}

type Number interface {
	UInts | Ints | Floats
}

type General interface {
	bool | UInts | Ints | Floats | string
}

type Any = interface{}
