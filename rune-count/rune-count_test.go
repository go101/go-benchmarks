package main

import (
	"testing"
	"strings"
	"unicode/utf8"
)

var s = strings.Repeat("你好abc", 16)

var n1 int
func Benchmark_CountStringRunes_Var_ForRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		for range s {
			count++
		}
		n1 = count
	}
}


var n2 int
func Benchmark_CountStringRunes_Var_Utf8Lib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n2 = utf8.RuneCountInString(s)
	}
}

var n3 int
func Benchmark_CountStringRunes_Var_LenRuneSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n3 = len([]rune(s))
	}
}



const c = "你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc你好abc"

var n1c int
func Benchmark_CountStringRunes_Const_ForRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		for range c {
			count++
		}
		n1c = count
	}
}


var n2c int
func Benchmark_CountStringRunes_Const_Utf8Lib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n2c = utf8.RuneCountInString(c)
	}
}

var n3c int
func Benchmark_CountStringRunes_Const_LenRuneSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n3c = len([]rune(c))
	}
}

var n4 int
func Benchmark_CountStringRunes_Var_LenByteSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n4 = len([]byte(s))
	}
}

var n5 int
func Benchmark_CountStringRunes_Var_LenString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n5 = len(s)
	}
}