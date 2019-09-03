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

var it, it2 int
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
func Benchmark_IntAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it = it2
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

