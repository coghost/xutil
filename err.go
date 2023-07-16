package xutil

import (
	"fmt"
	"reflect"
	"time"

	"github.com/avast/retry-go"
	"github.com/rs/zerolog/log"
	"github.com/ungerik/go-dry"
)

// LogFatalIfErr logs error with level fatal if `err!=nil`, else do nothing
func LogFatalIfErr(err error, msg string) {
	if err != nil {
		log.Fatal().CallerSkipFrame(1).Err(err).Msg(msg)
	}
}

// LogIfErr logs error with level error, and return true if err!=nil, else false
//
//	@return bool
func LogIfErr(err error, msg string) bool {
	if err != nil {
		log.Error().CallerSkipFrame(1).Err(err).Msg(msg)
		return true
	}
	return false
}

// PanicIfErr panics if has error
func PanicIfErr(args ...interface{}) {
	for _, v := range args {
		if err, _ := v.(error); err != nil {
			panic(fmt.Errorf("Panicking because of error: %s\nAt:\n%s\n", err, dry.StackTrace(3)))
		}
	}
}

func PanicIfErrWithPause(args ...interface{}) {
	hasErr := false
	for _, v := range args {
		if err, _ := v.(error); err != nil {
			hasErr = true
		}
	}
	if hasErr && uc.pauseInPanic {
		Pause("press any key to quit...")
	}
	PanicIfErr(args...)
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
