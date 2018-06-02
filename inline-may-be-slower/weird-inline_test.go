package main

import (
	"testing"
)

const N = 4096
type T int64
var a [N]T

var globalSum T

func sumByLoopArray(p *[N]T) T {
	var sum T
	for i := 0; i < len(p); i++ {
		sum += T(p[i])
	}
	return sum
}

//============================================

func Benchmark_GlobalArray_NoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalSum = sumByLoopArray(&a)
	}
}

func Benchmark_GlobalArray_Inline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sum T
		for i := 0; i < len(a); i++ {
			sum += T(a[i])
		}
		globalSum = sum
	}
}

func Benchmark_GlobalArray_Inline_AnonymousFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &a
		globalSum = func(p *[N]T) T {
			var sum T
			for i := 0; i < len(a); i++ {
				sum += T(a[i])
			}
			return sum
		}(p)
	}
}

func Benchmark_LocalArray_NoInline(b *testing.B) {
	var a [N]T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		globalSum = sumByLoopArray(&a)
	}
}

func Benchmark_LocalArray_Inline(b *testing.B) {
	var a [N]T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum T
		for i := 0; i < len(a); i++ {
			sum += T(a[i])
		}
		globalSum = sum
	}
}

func Benchmark_LocalArray_Inline_AnonymousFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &a
		globalSum = func(p *[N]T) T {
			var sum T
			for i := 0; i < len(p); i++ {
				sum += T(p[i])
			}
			return sum
		}(p)
	}
}


