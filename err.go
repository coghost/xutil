package xutil

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/avast/retry-go"
	"github.com/ungerik/go-dry"
)

func PanicIfErr(args ...interface{}) {
	paused := false
	for _, v := range args {
		if err, _ := v.(error); err != nil {
			paused = true
		}
	}
	if paused && uc.pauseInPanic {
		Pause("press any key to quit...")
	}
	dry.PanicIfErr(args...)
}

// ErrTry error
//
// which is extracted from go-rod
type ErrTry struct {
	Value interface{}
}

func (e *ErrTry) Error() string {
	return fmt.Sprintf("error value: %#v", e.Value)
}

// Is interface
func (e *ErrTry) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}

// CatchPanicAsErr
//
// try fn with recover, return the panic as error
func CatchPanicAsErr(fn func()) (err error) {
	defer func() {
		if val := recover(); val != nil {
			var ok bool
			err, ok = val.(error)
			if !ok {
				err = &ErrTry{val}
			}
		}
	}()
	fn()
	return err
}

// EnsureByRetry
//
// Params:
//   - fn
//   - args: tries, delay, showLogOrNot
//
// wrapper of retry.Do with default
// - total try 3 times
//   - overwritten by args[0]
//   - WARN: if args[0] == 0, the "fn" will be skipped
//   - if args[0] < 0, will use default 3
//
// - retry delay of 0 millisecond
//   - overwritten by args[1]
//
// - on retry not print logs
//   - overwritten by args[2]
func EnsureByRetry(fn func() error, args ...int) (tried int, err error) {
	var tries uint = 3
	if v := FirstOrDefaultArgs(3, args...); v >= 0 {
		tries = uint(v)
	}

	delay := 0
	if len(args) > 1 {
		delay = args[1]
	}

	show := 0
	if len(args) > 2 {
		show = args[2]
	}

	if show > 0 {
		log.Printf("[retry args]: tries, delay, show = %d, %d, %d", tries, delay, show)
	}

	err = retry.Do(
		func() error {
			tried += 1
			return fn()
		},
		retry.Attempts(tries),
		retry.Delay(time.Millisecond*time.Duration(delay)),
		retry.OnRetry(func(n uint, err error) {
			if show > 0 {
				log.Printf("#%d: %s", n, err)
			}
		}),
	)
	return
}
