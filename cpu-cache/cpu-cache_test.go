package main

import (
	"testing"
)

const N = 1024 * 64 // number of elements
type Element int64
var sum int

func Benchmark_SumArrayElements_GlobalVar(b *testing.B) {
	s := make([]Element, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = 0
		for _, v := range s {
			sum += int(v)
		}
	}
}

func Benchmark_SumArrayElements_LocalVar(b *testing.B) {
	s := make([]Element, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var g int
		for _, v := range s {
			g += int(v)
		}
		sum = g
	}
}
