package xstrings

import (
	"strings"
)

// Dedent removes the common indentation of all non-empty lines in s after removing the opening and closing empty lines.
func Dedent(s string) string {
	// The opening line must be only /\r?\n/.
	// The opening line's newline is removed.
	if strings.HasPrefix(s, "\r\n") {
		s = s[2:]
	} else if strings.HasPrefix(s, "\n") {
		s = s[1:]
	} else {
		return s
	}
	// The closing line's whitespace is removed.
	// The closing line's preceding /\r?\n/ is removed.
	s = strings.TrimRight(s, " \t")
	if strings.HasSuffix(s, "\r\n") {
		s = s[:len(s)-2]
	} else if strings.HasSuffix(s, "\n") {
		s = s[:len(s)-1]
	}
	// Lines containing only whitespace are emptied.
	var b strings.Builder
	for l := range strings.Lines(s) {
		if tl := strings.TrimLeft(l, " \t"); tl == "\r\n" || tl == "\n" {
			b.WriteString(tl)
		} else {
			b.WriteString(l)
		}
	}
	s = b.String()
	// The common indentation of all non-empty lines is calculated.
	var common string
	for l := range strings.Lines(s) {
		// Lines containing only whitespace have already been emptied.
		if l == "\r\n" || l == "\n" {
			continue
		}
		var leading string
		for i, c := range []byte(l) {
			if !(c == ' ' || c == '\t') {
				leading = l[:i]
				break
			}
		}
		if common == "" {
			common = leading
		} else {
			for i := range min(len(common), len(leading)) {
				if common[i] != leading[i] {
					common = common[:i]
					break
				}
			}
		}
	}
	// The common indentation is removed from the beginning of each line.
	b = strings.Builder{}
	for l := range strings.Lines(s) {
		l = strings.TrimPrefix(l, common)
		b.WriteString(l)
	}
	s = b.String()
	return s
}
