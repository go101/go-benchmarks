package main

import "testing"

type StructSlice1 struct{a T}
var rStructSlice1 T
func Benchmark_RangeStructSlice1(b *testing.B) {
	values := make([]StructSlice1, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice1 = 0
		for _, v := range values {
			rStructSlice1 = v.a // += v.a + v.a
		}
	}
}

type StructSlice2 struct{a, b T}
var rStructSlice2 T
func Benchmark_RangeStructSlice2(b *testing.B) {
	values := make([]StructSlice2, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice2 = 0
		for _, v := range values {
			rStructSlice2 = v.a // += v.a + v.b
		}
	}
}

type StructSlice3 struct{a, b, c T}
var rStructSlice3 T
func Benchmark_RangeStructSlice3(b *testing.B) {
	values := make([]StructSlice3, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice3 = 0
		for _, v := range values {
			rStructSlice3 = v.a // += v.a + v.c
		}
	}
}

type StructSlice4 struct{a, b, c, d T}
var rStructSlice4 T
func Benchmark_RangeStructSlice4(b *testing.B) {
	values := make([]StructSlice4, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice4 = 0
		for _, v := range values {
			rStructSlice4 = v.a // += v.a + v.d
		}
	}
}

type StructSlice5 struct{a, b, c, d, e T}
var rStructSlice5 T
func Benchmark_RangeStructSlice5(b *testing.B) {
	values := make([]StructSlice5, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice5 = 0
		for _, v := range values {
			rStructSlice5 = v.a // += v.a + v.e
		}
	}
}

type StructSlice6 struct{a, b, c, d, e, f T}
var rStructSlice6 T
func Benchmark_RangeStructSlice6(b *testing.B) {
	values := make([]StructSlice6, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice6 = 0
		for _, v := range values {
			rStructSlice6 = v.a // += v.a + v.f
		}
	}
}

type StructSlice7 struct{a, b, c, d, e, f, g T}
var rStructSlice7 T
func Benchmark_RangeStructSlice7(b *testing.B) {
	values := make([]StructSlice7, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice7 = 0
		for _, v := range values {
			rStructSlice7 = v.a // += v.a + v.g
		}
	}
}

type StructSlice8 struct{a, b, c, d, e, f, g, h T}
var rStructSlice8 T
func Benchmark_RangeStructSlice8(b *testing.B) {
	values := make([]StructSlice8, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice8 = 0
		for _, v := range values {
			rStructSlice8 = v.a // += v.a + v.h
		}
	}
}

type StructSlice9 struct{a, b, c, d, e, f, g, h, i T}
var rStructSlice9 T
func Benchmark_RangeStructSlice9(b *testing.B) {
	values := make([]StructSlice9, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStructSlice9 = 0
		for _, v := range values {
			rStructSlice9 = v.a // += v.a + v.i
		}
	}
}

