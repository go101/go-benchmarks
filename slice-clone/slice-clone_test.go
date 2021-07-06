package main

import (
	"testing"
)

type T = int
const N = 99

var sMakeCopy []T
func Benchmark_MakeCopy(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sMakeCopy = make([]T, len(s))
		copy(sMakeCopy, s)
	}
}

var sMakeAppend []T
func Benchmark_MakeAppend(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sMakeAppend = append(make([]T, 0, len(s)), s...)
	}
}

var sAppendA []T
func Benchmark_AppendA(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sAppendA = append([]T(nil), s...)
	}
}

var sAppendB []T
func Benchmark_AppendB(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sAppendB = append(s[:0:0], s...)
	}
}

var sVerbose []T
func Benchmark_Verbose(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sVerbose = nil
		if s != nil {
			sVerbose = make([]T, len(s))
			copy(sVerbose, s)
		}
	}
}

