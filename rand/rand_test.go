package main

import (
	"testing"
	mathrand "math/rand"
	cryptorand "crypto/rand"
)

const M = 8
var x1, x2, x3 = make([]byte, M), make([]byte, M), make([]byte, M)
const N = 8192
var y1, y2, y3 = make([]byte, N), make([]byte, N), make([]byte, N)
var err error
var n1, n2 int64

func init() {
	print("M = ", M, ", N = ", N, "\n")
}

var r = mathrand.New(mathrand.NewSource(99))

func Benchmark_MathUnsyncedInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n1 = r.Int63()
	}
}

func Benchmark_MathGlobalInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n2 = mathrand.Int63()
	}
}

func Benchmark_MathUnsyncedRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Read(x1)
	}
}

func Benchmark_MathGlobalRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Read(x2)
	}
}

func Benchmark_CryptoRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptorand.Read(x3)
	}
}

func Benchmark_MathUnsyncedRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Read(y1)
	}
}

func Benchmark_MathGlobalRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Read(y2)
	}
}

func Benchmark_CryptoRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptorand.Read(y3)
	}
}