package main

import (
	"testing"
	"fmt"
)

const N = 1024 * 100
var a [N]int

func init() {
	println("N =", N)
}

func Benchmark_ArrayPtr_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range &a {
			a[i] = 0
		}
	}
}

func Benchmark_Memclr_Array_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range a {
			a[i] = 0
		}
	}
}

func Benchmark_Memclr_Slice_N(b *testing.B) {
	s := a[:]
	for i := 0; i < b.N; i++ {
		for i := range s {
			s[i] = 0
		}
	}
}

func Benchmark_(b *testing.B) {
	const MaxN = 1024 * 1024 * 32
	var a [MaxN]int
	
	for n := MaxN >> 1; n >= 2; n >>= 1 {
		b.Run(fmt.Sprintf("ArrayPtr_%d", n), func(b *testing.B){
			for i := 0; i < b.N; i++ {
				for k := 0; k < n; k++ {
					a[k] = 0
				}
			}
		})
	}
	
	for n := MaxN >> 1; n >= 2; n >>= 1 {
		b.Run(fmt.Sprintf("Memclr_Slice_%d", n), func(b *testing.B){
			for i := 0; i < b.N; i++ {
				s := a[:n]
				for k := range s {
					s[k] = 0
				}
			}
		})
	}
}
