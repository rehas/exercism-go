// Package hamming provides utility methods for strand comparisons
package hamming

import (
	"errors"
)

// Distance measures the count of differences between strands
func Distance(a, b string) (count int, err error) {
	rsa := []rune(a)
	rsb := []rune(b)

	if len(rsa) != len(rsb) {
		err = errors.New("not equal lengths")
		return
	}
	for i := range rsa {
		if a[i] != b[i] {
			count++
		}
	}
	return
}
