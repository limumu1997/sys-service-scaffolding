package util

import (
	"bytes"
)

func FilterArrayByRule(src, rule []byte) []byte {
	i := 0
	for k, v := range src {
		if bytes.Contains(rule, src[k:k+1]) {
			continue
		}
		src[i] = v
		i++
	}
	return src[:i]
}
