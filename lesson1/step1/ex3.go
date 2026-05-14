package step1

import (
	"strings"
)

type UpperWriter struct {
	UpperString string
}

func (uw *UpperWriter) Write(p []byte) (n int, err error) {
	upperStr := strings.ToUpper(string(p))
	uw.UpperString += upperStr

	return len(p), nil
}
