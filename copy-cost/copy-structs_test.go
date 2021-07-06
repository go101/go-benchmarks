package main

import "testing"

type Struct1 struct{a T}
var rStruct1, sStruct1 Struct1
func Benchmark_CopyStruct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct1 = sStruct1
	}
}

type Struct2 struct{a, b T}
var rStruct2, sStruct2 Struct2
func Benchmark_CopyStruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct2 = sStruct2
	}
}

type Struct3 struct{a, b, c T}
var rStruct3, sStruct3 Struct3
func Benchmark_CopyStruct3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct3 = sStruct3
	}
}

type Struct4 struct{a, b, c, d T}
var rStruct4, sStruct4 Struct4
func Benchmark_CopyStruct4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct4 = sStruct4
	}
}

type Struct5 struct{a, b, c, d, e T}
var rStruct5, sStruct5 Struct5
func Benchmark_CopyStruct5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct5 = sStruct5
	}
}

type Struct6 struct{a, b, c, d, e, f T}
var rStruct6, sStruct6 Struct6
func Benchmark_CopyStruct6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct6 = sStruct6
	}
}

type Struct7 struct{a, b, c, d, e, f, g T}
var rStruct7, sStruct7 Struct7
func Benchmark_CopyStruct7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct7 = sStruct7
	}
}

type Struct8 struct{a, b, c, d, e, f, g, h T}
var rStruct8, sStruct8 Struct8
func Benchmark_CopyStruct8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct8 = sStruct8
	}
}

type Struct9 struct{a, b, c, d, e, f, g, h, i T}
var rStruct9, sStruct9 Struct9
func Benchmark_CopyStruct9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rStruct9 = sStruct9
	}
}



