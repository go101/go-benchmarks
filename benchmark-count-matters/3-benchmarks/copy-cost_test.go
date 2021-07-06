package concat

import (
	"testing"
)

var struct3_0 struct{ a, b, c int }

func Benchmark_CopyStruct_3_fields(b *testing.B) {
	var struct3_1 struct{ a, b, c int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range [10]struct{}{} {
			struct3_0 = struct3_1
		}
	}
}

var struct4_0 struct{ a, b, c, d int }

func Benchmark_CopyStruct_4_fields(b *testing.B) {
	var struct4_1 struct{ a, b, c, d int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range [10]struct{}{} {
			struct4_0 = struct4_1
		}
	}
}

func Benchmark_xxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
