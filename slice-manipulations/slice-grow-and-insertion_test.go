package main

import (
	"testing"
)

type T = int
var sx = []T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
var sy = []T{111, 222, 333, 444}
var sz = make([]T, 2000)
var index = 6

func SliceUnshift(base []T, t T) []T {
	return append([]T{t}, base...)
}

func SliceInsertOne(base []T, t T, i int) []T {
	return append(base[:i], append([]T{t}, base[i:]...)...)
}

func SliceGrow_OneLine(base []T, newCapacity int) []T {
	return append(base, make([]T, newCapacity-cap(base))...)
}

func SliceGrow_VerboseAppend(base []T, newCapacity int) []T {
	m := make([]T, newCapacity-cap(base))
	return append(base, m...)
}

func SliceGrow_VerboseCopy(base []T, newCapacity int) []T {
	m := make([]T, newCapacity)
	copy(m, base)
	return m
}

func SliceGrow_VerboseOneByOne(base []T, newCapacity int) []T {
	n := newCapacity - len(base)
	for i := 0; i < n; i++ {
		base = append(base, 0)
	}
	return base
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

func SliceMemclr_Loop(s []T, start, end int) {
	s = s[start:end]
	for i := range s {
		s[i] = 0
	}
}

func SliceMemclr_Append(s []T, start, end int) {
	_ = append(s[:start], make([]T, end - start)...)
}

var su []T
func Benchmark_SliceUnshift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		su = SliceUnshift(sx, 123)
	}
}

var si1 []T
func Benchmark_SliceInsertOne_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		si1 = SliceInsertOne(sx, 123, 6)
	}
}

var si2 []T
func Benchmark_SliceInsertOne_Verbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		si2 = SliceInsertion_Verbose(sx, []T{123}, 6)
	}
}

var sa []T
func Benchmark_SliceGrow_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sa = SliceGrow_OneLine(sx, 100)
	}
}

var sb []T
func Benchmark_SliceGrow_VerboseAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb = SliceGrow_VerboseAppend(sx, 100)
	}
}

var sc []T
func Benchmark_SliceGrow_VerboseCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc = SliceGrow_VerboseCopy(sx, 100)
	}
}

var sd []T
func Benchmark_SliceGrow_VerboseOneByOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sd = SliceGrow_VerboseOneByOne(sx, 100)
	}
}

var s1 []T
func Benchmark_SliceInsertion_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s1 = SliceInsertion_OneLine(sx, sy, index)
	}
}

var s2 []T
func Benchmark_SliceInsertion_Verbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 = SliceInsertion_Verbose(sx, sy, index)
	}
}

func Benchmark_SliceMemclr_Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceMemclr_Loop(sz, 1000, 2000)
	}
}

func Benchmark_SliceMemclr_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceMemclr_Append(sz, 1000, 2000)
	}
}


