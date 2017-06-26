package main

import (
	"testing"
)

const N = 32
type T byte
var sum T

func Benchmark_RangeSlice_OneIterationVar(b *testing.B) {
	var s = make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var g T
		for j := range s {
			g += T(s[j])
		}
		sum = g
	}
}

func Benchmark_LoopSlice(b *testing.B) {
	var s = make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var g T
		for j := 0; j < len(s); j++ {
			g += T(s[j])
		}
		sum = g
	}
}

func Benchmark_RangeSlice_TwoIterationVar(b *testing.B) {
	var s = make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var g T
		for _, v := range s {
			g += T(v)
		}
		sum = g
	}
}
