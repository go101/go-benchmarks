package main

import (
	"testing"
)

type T struct {
	value int
	next  *T
}

var ts = [1024]T{}

func init() {
	for i := 1; i < len(ts); i++ {
		ts[i-1].next = &ts[i]
	}
}

var a int
func Benchmark_ByIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < len(ts); i++ {
			sum += ts[i].value
		}
		a = sum
	}
}

var x int
func Benchmark_ByChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for t := &ts[0]; t != nil; t = t.next {
			sum += t.value
		}
		x = sum
	}
}

