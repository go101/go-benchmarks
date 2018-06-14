package main

import (
	"testing"
	"fmt"
	"bytes"
	"strings"
)

var words = []string{"hello", " ", "world", "!"}
var result string

func Benchmark_AddStrings_KnownCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = words[0] + words[1] + words[2] + words[3]
	}
}

func Benchmark_AddStrings_UnnownCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var r = ""
		for _, w := range words {
			result += w
		}
		result = r
	}
}

func Benchmark_FmtSprintf_KnownCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%s%s%s%s",  words[0], words[1], words[2], words[3])
	}
}

func Benchmark_StringsJoin_NoSliceAllocated(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = strings.Join(words, "")
	}
}

func Benchmark_StringsJoin_SliceAllocated(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ws = append([]string(nil), words...)
		result = strings.Join(ws, "")
	}
}

func Benchmark_BytesBuffer_New(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, w := range words {
			buf.WriteString(w)
		}
		result = buf.String()
	}
}

var gBuf bytes.Buffer
func Benchmark_AddStrings_Buffer_Reuse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gBuf.Reset()
		for _, w := range words {
			gBuf.WriteString(w)
		}
		result = gBuf.String()
	}
}

func Benchmark_AddStrings_Builder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, w := range words {
			builder.WriteString(w)
		}
		result = builder.String()
	}
}

// todo: strings.Builder