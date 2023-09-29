package calc

import (
	"github.com/google/go-cmp/cmp"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Equal(x, y int) bool {
	return cmp.Equal(x, y)
}
