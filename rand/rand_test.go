package main

import (
	"testing"
	mathrand "math/rand"
	cryptorand "crypto/rand"
)

const M = 8
var x = make([]byte, M)
const N = 8192
var y = make([]byte, N)
var err error
var n int64

func init() {
	print("M = ", M, ", N = ", N, "\n")
}

var r = mathrand.New(mathrand.NewSource(99))

func Benchmark_MathUnsyncedInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n = r.Int63()
	}
}

func Benchmark_MathGlobalInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n = mathrand.Int63()
	}
}

func Benchmark_MathUnsyncedRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Read(x)
	}
}

func Benchmark_MathGlobalRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Read(x)
	}
}

func Benchmark_CryptoRandRead_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptorand.Read(x)
	}
}

func Benchmark_MathUnsyncedRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Read(y)
	}
}

func Benchmark_MathGlobalRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathrand.Read(y)
	}
}

func Benchmark_CryptoRandRead_8192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptorand.Read(y)
	}
}