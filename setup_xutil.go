package xutil

type XOpts struct {
	debug   bool
	recover bool

	dummyLog     bool
	pauseInPanic bool
}

type XOptFunc func(o *XOpts)

func bindXOpts(opt *XOpts, opts ...XOptFunc) {
	for _, f := range opts {
		f(opt)
	}
}

func SetDebug(b bool) XOptFunc {
	return func(o *XOpts) {
		o.debug = b
	}
}

func SetRecover(b bool) XOptFunc {
	return func(o *XOpts) {
		o.recover = b
	}
}

func SetDummyLog(b bool) XOptFunc {
	return func(o *XOpts) {
		o.dummyLog = b
	}
}

func SetPauseInPanic(b bool) XOptFunc {
	return func(o *XOpts) {
		o.pauseInPanic = b
	}
}

var uc = &XOpts{}

// InitializeXOpts setups
//   - debug: used in `PauseToDebug`
//   - recover: used in `RecoverAndDumpOnly`
//   - dummyLog: used in `DummyLog`
//   - pauseInPanic: use in `PanicIfErr`
func InitializeXOpts(opts ...XOptFunc) {
	opt := XOpts{}
	bindXOpts(&opt, opts...)
	uc = &opt
}
