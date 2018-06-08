package squash

import (
	"unicode"
	"unicode/utf8"
)

func squash(in []byte) string {
	i := 0
	for i < len(in) {
		ru, size := utf8.DecodeRune(in[i:])
		if unicode.IsSpace(ru) {
			copy(in[i:], in[i+size:])
			in = in[:len(in)-size]
		} else {
			i += size
		}
	}
	return string(in[:i])
}
