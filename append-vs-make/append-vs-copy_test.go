package main

import (
	"testing"
)

const N = 1024 * 1024
type Element int64
var x = make([]Element, N)
var y []Element

func Benchmark_AllocWithMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = make([]Element, N)
	}
}

func Benchmark_AllocWithAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = append([]Element(nil), x...)
	}
}
