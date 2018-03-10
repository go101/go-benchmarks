package main

import (
	"testing"
)

var v1a int
func Benchmark_CSE_Case_1a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x, y, z = 123, 456, 789
		v1a = (y+z) / x + (y+z) / x + (y+z) / x
	}
}

var v1b int
func Benchmark_CSE_Case_1b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x, y, z = 123, 456, 789
		t := (y+z) / x
		v1b = t + t + t
	}
}

var v2Array = []int{123, 456, 789}
var v2a int
func Benchmark_CSE_Case_2a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v2a = (v2Array[1]+v2Array[2]) / v2Array[0] +
			(v2Array[1]+v2Array[2]) / v2Array[0] +
			(v2Array[1]+v2Array[2]) / v2Array[0]
	}
}

var v2b int
func Benchmark_CSE_Case_2b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := (v2Array[1]+v2Array[2]) / v2Array[0]
		v2b = t + t + t
	}
}
