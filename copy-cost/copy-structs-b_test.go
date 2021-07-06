package main

import "testing"

var rStruct1b Struct1
func Benchmark_CopyStruct_b1(b *testing.B) {
	var v Struct1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct1b = v
	}
}

var rStruct2b Struct2
func Benchmark_CopyStruct_b2(b *testing.B) {
	var v Struct2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct2b = v
	}
}

var rStruct3b Struct3
func Benchmark_CopyStruct_b3(b *testing.B) {
	var v Struct3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct3b = v
	}
}

var rStruct4b Struct4
func Benchmark_CopyStruct_b4(b *testing.B) {
	var v Struct4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct4b = v
	}
}

var rStruct5b Struct5
func Benchmark_CopyStruct_b5(b *testing.B) {
	var v Struct5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct5b = v
	}
}

var rStruct6b Struct6
func Benchmark_CopyStruct_b6(b *testing.B) {
	var v Struct6
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct6b = v
	}
}

var rStruct7b Struct7
func Benchmark_CopyStruct_b7(b *testing.B) {
	var v Struct7
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct7b = v
	}
}

var rStruct8b Struct8
func Benchmark_CopyStruct_b8(b *testing.B) {
	var v Struct8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct8b = v
	}
}

var rStruct9b Struct9
func Benchmark_CopyStruct_b9(b *testing.B) {
	var v Struct9
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rStruct9b = v
	}
}



