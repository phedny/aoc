package util

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed](n T) T {
	y := n >> (8*unsafe.Sizeof(n) - 1)
	return (n ^ y) - y
}
