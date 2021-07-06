package main

import (
	"testing"
)

const NLoops = 1024

var p, p2 = new(int), new(int)
var ip interface{} = p
func Benchmark_BoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip = p
	}
}
func Benchmark_PointerAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p = ip.(*int)
	}
}
func Benchmark_PointerAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p = p2
	}
}

var it, it2 int = 66666, 66666
var it_small, it2_small = 255, 255
var ii interface{} = it
func Benchmark_BoxInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ii = it
	}
}
func Benchmark_IntAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it = ii.(int)
	}
}
func Benchmark_BoxSmallInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ii = it_small
	}
}
func Benchmark_SmallIntAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it_small = ii.(int)
	}
}
func Benchmark_IntAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it = it2
	}
}

var by, by2 byte
var iby interface{} = by
func Benchmark_BoxByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iby = by
	}
}
func Benchmark_ByteAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		by = iby.(byte)
	}
}
func Benchmark_ByteAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		by = by2
	}
}

var bl, bl2 bool
var ib interface{} = bl
func Benchmark_BoxBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ib = bl
	}
}
func Benchmark_BoolAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl = ib.(bool)
	}
}
func Benchmark_BoolAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl = bl2
	}
}

var a, a2 [16]int64
var ia interface{} = a
func Benchmark_BoxArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ia = a
	}
}
func Benchmark_ArrayAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = ia.([16]int64)
	}
}
func Benchmark_ArrayAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = a2
	}
}

var st, st2 struct{}
var ist interface{} = st
func Benchmark_BoxBlankStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ist = st
	}
}
func Benchmark_BlankStructAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st = ist.(struct{})
	}
}
func Benchmark_BlankStructAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st = st2
	}
}

var sc, sc2 = []int{}, []int{}
var isc interface{} = sc
func Benchmark_BoxSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isc = sc
	}
}
func Benchmark_SliceAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc = isc.([]int)
	}
}
func Benchmark_SliceAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc = sc2
	}
}

var fc, fc2 = func(){}, func(){}
var ifc interface{} = fc
func Benchmark_BoxFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifc = fc
	}
}
func Benchmark_FunctionAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fc = ifc.(func())
	}
}
func Benchmark_FunctionAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fc = fc2
	}
}

