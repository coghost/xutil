package xutil

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// TruncateString
//
// this will TruncateString with rune
func TruncateString(s string, max int) string {
	if max <= 0 {
		return ""
	}

	if utf8.RuneCountInString(s) < max {
		return s
	}

	return string([]rune(s)[:max])
}

func TruncateText(s string, max int) string {
	if max <= 0 {
		return ""
	}
	if max >= len(s) {
		return s
	}
	return s[:max]
}

func TruncateWord(s string, max int) string {
	if max <= 0 {
		return ""
	}
	if max >= len(s) {
		return s
	}
	return s[:strings.LastIndexAny(s[:max], " .,:;-")]
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func MustString(obj interface{}) string {
	if v, ok := obj.(string); ok {
		return v
	}
	panic(fmt.Sprintf("obj is %+v, not a valid string", obj))
}
