package main

import "testing"

var rArray1b Array1
func Benchmark_CopyArray_b1(b *testing.B) {
	var v Array1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray1b = v
	}
}

var rArray2b Array2
func Benchmark_CopyArray_b2(b *testing.B) {
	var v Array2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray2b = v
	}
}

var rArray3b Array3
func Benchmark_CopyArray_b3(b *testing.B) {
	var v Array3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray3b = v
	}
}

var rArray4b Array4
func Benchmark_CopyArray_b4(b *testing.B) {
	var v Array4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray4b = v
	}
}

var rArray5b Array5
func Benchmark_CopyArray_b5(b *testing.B) {
	var v Array5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray5b = v
	}
}

var rArray6b Array6
func Benchmark_CopyArray_b6(b *testing.B) {
	var v Array6
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray6 = v
	}
}

var rArray7b Array7
func Benchmark_CopyArray_b7(b *testing.B) {
	var v Array7
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray7 = v
	}
}

var rArray8b Array8
func Benchmark_CopyArray_b8(b *testing.B) {
	var v Array8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray8 = v
	}
}

var rArray9b Array9
func Benchmark_CopyArray_b9(b *testing.B) {
	var v Array9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rArray9b = v
	}
}



