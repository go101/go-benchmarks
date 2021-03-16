package main

import (
	"testing"
)

func ifelse(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

var m = map[bool]int{true: 1, false: 0}
func boolmap(x bool) int {
	return m[x]
}

func b2i(x bool) int {
	v := 0
	if x {
		v = 1
	}
	return v
}

var values = [...]int{0, 1}
func sliceindex(x bool) int {
	return values[b2i(x)&1]
}

	

var r1 int
func Benchmark_ifelse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = ifelse(true)
		r1 = ifelse(false)
	}
}

var r2 int
func Benchmark_boolmap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r2 = boolmap(true)
		r2 = boolmap(false)
	}
}

var r3 int
func Benchmark_sliceindex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r3 = sliceindex(true)
		r3 = sliceindex(false)
	}
}


