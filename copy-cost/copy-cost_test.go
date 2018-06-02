package main

import "testing"

type T int

type S2 struct{a, b T}
type S3 struct{a, b, c T}
type S4 struct{a, b, c, d T}
type S5 struct{a, b, c, d, e T}
type S6 struct{a, b, c, d, e, f T}
var s2X, s2Y, s2Z = make([]S2, 1000), make([]S2, 1000), make([]S2, 1000)
var s3X, s3Y, s3Z = make([]S3, 1000), make([]S3, 1000), make([]S3, 1000)
var s4X, s4Y, s4Z = make([]S4, 1000), make([]S4, 1000), make([]S4, 1000)
var s5X, s5Y, s5Z = make([]S5, 1000), make([]S5, 1000), make([]S5, 1000)
var s6X, s6Y, s6Z = make([]S6, 1000), make([]S6, 1000), make([]S6, 1000)
var sum2X, sum2Y, sum2Z T
var sum3X, sum3Y, sum3Z T
var sum4X, sum4Y, sum4Z T
var sum5X, sum5Y, sum5Z T
var sum6X, sum6Y, sum6Z T



// 2

func Benchmark_Loop_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum2X = 0
		for j := 0; j < len(s2X); j++ {
			sum2X += s2X[j].a
		}
	}
}

func Benchmark_Range_OneIterVar_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum2Z = 0
		for j := range s2Y {
			sum2Z += s2Y[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum2Y = 0
		for _, v := range s2Y {
			sum2Y += v.a
		}
	}
}

// 3

func Benchmark_Loop_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum3X = 0
		for j := 0; j < len(s3X); j++ {
			sum3X += s3X[j].a
		}
	}
}

func Benchmark_Range_OneIterVar_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum3Z = 0
		for j := range s3Y {
			sum3Z += s3Y[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum3Y = 0
		for _, v := range s3Y {
			sum3Y += v.a
		}
	}
}

// 4

func Benchmark_Loop_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum4X = 0
		for j := 0; j < len(s4X); j++ {
			sum4X += s4X[j].a
		}
	}
}

func Benchmark_Range_OneIterVar_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum4Z = 0
		for j := range s4Y {
			sum4Z += s4Y[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum4Y = 0
		for _, v := range s4Y {
			sum4Y += v.a
		}
	}
}

// 5

func Benchmark_Loop_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum5X = 0
		for j := 0; j < len(s5X); j++ {
			sum5X += s5X[j].a
		}
	}
}

func Benchmark_Range_OneIterVar_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum5Z = 0
		for j := range s5Y {
			sum5Z += s5Y[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum5Y = 0
		for _, v := range s5Y {
			sum5Y += v.a
		}
	}
}

// 6

func Benchmark_Loop_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum6X = 0
		for j := 0; j < len(s6X); j++ {
			sum6X += s6X[j].a
		}
	}
}

func Benchmark_Range_OneIterVar_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum6Z = 0
		for j := range s6Y {
			sum6Z += s6Y[j].a
		}
	}
}

func Benchmark_Range_TwoIterVar_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum6Y = 0
		for _, v := range s6Y {
			sum6Y += v.a
		}
	}
}