package main

import "testing"

type ArraySlice1 [1]T
var rArraySlice1 T
func Benchmark_RangeArraySlice1(b *testing.B) {
	values := make([]ArraySlice1, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice1 = 0
		for _, v := range values {
			rArraySlice1 = v[0] // += v[0] //+ v[0]
		}
	}
}

type ArraySlice2 [2]T
var rArraySlice2 T
func Benchmark_RangeArraySlice2(b *testing.B) {
	values := make([]ArraySlice2, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice2 = 0
		for _, v := range values {
			rArraySlice2 = v[0] // += v[0] //+ v[1]
		}
	}
}

type ArraySlice3 [3]T
var rArraySlice3 T
func Benchmark_RangeArraySlice3(b *testing.B) {
	values := make([]ArraySlice3, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice3 = 0
		for _, v := range values {
			rArraySlice3 = v[0] // += v[0] //+ v[2]
		}
	}
}

type ArraySlice4 [4]T
var rArraySlice4 T
func Benchmark_RangeArraySlice4(b *testing.B) {
	values := make([]ArraySlice4, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice4 = 0
		for _, v := range values {
			rArraySlice4 = v[0] // += v[0] //+ v[3]
		}
	}
}

type ArraySlice5 [5]T
var rArraySlice5 T
func Benchmark_RangeArraySlice5(b *testing.B) {
	values := make([]ArraySlice5, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice5 = 0
		for _, v := range values {
			rArraySlice5 = v[0] // += v[0] //+ v[4]
		}
	}
}

type ArraySlice6 [6]T
var rArraySlice6 T
func Benchmark_RangeArraySlice6(b *testing.B) {
	values := make([]ArraySlice6, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice6 = 0
		for _, v := range values {
			rArraySlice6 = v[0] // += v[0] //+ v[5]
		}
	}
}

type ArraySlice7 [7]T
var rArraySlice7 T
func Benchmark_RangeArraySlice7(b *testing.B) {
	values := make([]ArraySlice7, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice7 = 0
		for _, v := range values {
			rArraySlice7 = v[0] // += v[0] //+ v[6]
		}
	}
}

type ArraySlice8 [8]T
var rArraySlice8 T
func Benchmark_RangeArraySlice8(b *testing.B) {
	values := make([]ArraySlice8, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice8 = 0
		for _, v := range values {
			rArraySlice8 = v[0] // += v[0] //+ v[7]
		}
	}
}

type ArraySlice9 [9]T
var rArraySlice9 T
func Benchmark_RangeArraySlice9(b *testing.B) {
	values := make([]ArraySlice9, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArraySlice9 = 0
		for _, v := range values {
			rArraySlice9 = v[0] // += v[0] //+ v[8]
		}
	}
}

