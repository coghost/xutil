package xutil

import (
	"strings"
	"unicode/utf8"
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
