package parse

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Add("world")
	f.Add("")
	f.Fuzz(func(t *testing.T, original string) {
		reversed := Reverse(original)
		doubleReversed := Reverse(reversed)
		if original != doubleReversed {
			t.Errorf("double reverse produced different result: %q -> %q", original, doubleReversed)
		}
		if utf8.ValidString(original) && !utf8.ValidString(reversed) {
			t.Errorf("reverse produced invalid UTF-8: %q", reversed)
		}
	})
}
