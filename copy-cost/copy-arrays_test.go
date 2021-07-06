package main

import "testing"

type Array1 [1]T
var rArray1, sArray1 Array1
func Benchmark_CopyArray1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray1 = sArray1
	}
}

type Array2 [2]T
var rArray2, sArray2 Array2
func Benchmark_CopyArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray2 = sArray2
	}
}

type Array3 [3]T
var rArray3, sArray3 Array3
func Benchmark_CopyArray3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray3 = sArray3
	}
}

type Array4 [4]T
var rArray4, sArray4 Array4
func Benchmark_CopyArray4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray4 = sArray4
	}
}

type Array5 [5]T
var rArray5, sArray5 Array5
func Benchmark_CopyArray5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray5 = sArray5
	}
}

type Array6 [6]T
var rArray6, sArray6 Array6
func Benchmark_CopyArray6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray6 = sArray6
	}
}

type Array7 [7]T
var rArray7, sArray7 Array7
func Benchmark_CopyArray7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray7 = sArray7
	}
}

type Array8 [8]T
var rArray8, sArray8 Array8
func Benchmark_CopyArray8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray8 = sArray8
	}
}

type Array9 [9]T
var rArray9, sArray9 Array9
func Benchmark_CopyArray9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rArray9 = sArray9
	}
}



