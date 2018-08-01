package main

import (
	"testing"
)

const N = 1024 * 1024
type Element = int64
var xForCopy = make([]Element, N)
var xForMake = make([]Element, N)
var xForMakeCopy = make([]Element, N)
var xForAppend = make([]Element, N)
var yForCopy = make([]Element, N)
var yForMake []Element
var yForMakeCopy []Element
var yForAppend []Element

func Benchmark_PureCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(yForCopy, xForCopy)
	}
}

func Benchmark_PureMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMake = make([]Element, N)
	}
}

func Benchmark_MakeAndCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = make([]Element, N)
		copy(yForMakeCopy, xForMakeCopy)
	}
}

func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForAppend = append([]Element(nil), xForAppend...)
	}
}
