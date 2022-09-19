package xutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/coghost/xpretty"
	"github.com/gookit/goutil/dump"
	"github.com/thoas/go-funk"
)

var (
	Green = xpretty.Green
	Red   = xpretty.Red
	faint = xpretty.Faint
	info  = xpretty.Info
)

// DummyLog will print a dummy log with green bg
func DummyLog(msg ...interface{}) {
	if !uc.dummyLog {
		return
	}
	c, l := Caller(2)

	if funk.IsEmpty(msg) {
		fmt.Printf("%v %v(%v) %v\n", faint(StrNow(DFmt2)), Green(c), Green(l), info("dummy log"))
	} else {
		s := AnyJoin(", ", msg...)
		fmt.Printf("%v %v(%v) %v\n", faint(StrNow(DFmt2)), Green(c), Green(l), info(s))
	}
}

// DummyErrorLog will print a dummy log with red bg
func DummyErrorLog(msg ...interface{}) {
	if !uc.dummyLog {
		return
	}
	c, l := Caller(2)

	if funk.IsEmpty(msg) {
		fmt.Printf("%v %v(%v) %v\n", faint(StrNow(DFmt2)), Red(c), Red(l), info("dummy log"))
	} else {
		s := AnyJoin(", ", msg...)
		fmt.Printf("%v %v(%v) %v\n", faint(StrNow(DFmt2)), Red(c), Red(l), info(s))
	}
}

// DumpCallerStack
//
// print the caller tree
func DumpCallerWithKey(args ...string) {
	cont := FirstOrDefaultArgs("", args...)
	start := 2
	for {
		fn, line := Caller(start)
		if cont == "" || strings.HasPrefix(fn, cont) {
			if line != 0 {
				fmt.Println(fn, line)
			}
		}
		if fn == "" {
			break
		}

		start += 1
		if start > 128 {
			break
		}
	}

	fmt.Println("end with", start)
}

func DumpCallerStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	dump.P(string(buf[:n]))
}

func RecoverAndDumpOnly() {
	if !uc.recover {
		return
	}
	if r := recover(); r != nil {
		DumpCallerStack()
		dump.P(r)
	}
}

// PauseToDebug block the normal workflow, and used for debug purpose only
//
//	this is triggered by ctrl.debug = true
func PauseToDebug(msg ...string) bool {
	if !uc.debug {
		return false
	}
	Pause(msg...)
	return true
}

// Pause: press enter to continue
func Pause(args ...string) {
	msg := FirstOrDefaultArgs("Press Enter to continue", args...)

	pc, _, l, _ := runtime.Caller(1)
	c := runtime.FuncForPC(pc).Name()
	fmt.Printf("===> %v(%v) <===\n", c, l)

	xpretty.YellowPrintf("[%v] %v:", StrNow(), msg)
	reader := bufio.NewReader(os.Stdin)
	_, e := reader.ReadString('\n')
	if e != nil {
		log.Println("read line failed")
	}
}
