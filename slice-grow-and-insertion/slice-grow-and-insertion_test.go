package main

import (
	"testing"
)

type T = int
var sx = []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
var sy = []T{111, 222, 333, 444}
var index = 6

func SliceGrow_OneLine(base []T, newCapacity int) []T {
	return append(base, make([]T, newCapacity-cap(base))...)
}

func SliceGrow_Verbose(base []T, newCapacity int) []T {
	m := make([]T, newCapacity-cap(base))
	return append(base, m...)
}

func SliceInsertion_OneLine(base, inserted []T, i int) []T {
	return append(base[:i], append(inserted, base[i:]...)...)
}

func SliceInsertion_Verbose(base, inserted []T, i int) []T {
	if cap(base)-len(base) >= len(inserted) {
		s := base[:len(base)+len(inserted)]
		copy(s[i+len(inserted):], s[i:])
		copy(s[i:], inserted)
		return s
	} else {
		s := make([]T, 0, len(inserted)+len(base))
		s = append(s, base[:i]...)
		s = append(s, inserted...)
		s = append(s, base[i:]...)
		return s
	}
}

var sa []T
func Benchmark_SliceGrow_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sa = SliceGrow_OneLine(sx, 100)
	}
}

var sb []T
func Benchmark_SliceGrow_Verbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb = SliceGrow_Verbose(sx, 100)
	}
}

var s1 []T
func Benchmark_SliceInsertion_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s1 = SliceInsertion_OneLine(sx, sy, index)
	}
}

var s2 []T
func Benchmark_Verbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 = SliceInsertion_Verbose(sx, sy, index)
	}
}


