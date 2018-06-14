package main

import (
	"testing"
)

func Benchmark_TypeMethodAsFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fx()
	}
}

func Benchmark_ValueMethodAsFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fy()
	}
}

func Benchmark_ValueMethodAsFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fy2()
	}
}

func Benchmark_ValueMethodAsFunc_GlobalVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fz()
	}
}

