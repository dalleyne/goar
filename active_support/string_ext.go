package active_support

import (
	"bytes"
	"strings"
)

type String string

func (str String) Underscore() string {
	buf := bytes.NewBufferString("")
	for i, v := range str {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(v)
	}

	return strings.ToLower(buf.String())
}
