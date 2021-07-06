package main

import (
	"testing"
	"unsafe"
)

func str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var words [][]byte

func init() {
	var buildWord = func(b byte) []byte {
		w := []byte{b, b, b}
		w = append(w, w...)
		w = append(w, w...)
		w = append(w, w...)
		w = append(w, w...)
		w = append(w, w...)
		w = append(w, w...)
		return w
	}
	for range[100]struct{}{} {
		for j := range[100]struct{}{} {
			w := buildWord(byte(j))
			words = append(words, w)
		}
	}
	
	var m map[string]int
	var check = func(s string) {
		for j := range[100]struct{}{} {
			w := buildWord(byte(j))
			if m[string(w)] != 100 {
				panic(s + "is wrong: " + string(w))
			}
		}
	}
	m = IntAdd(words)
	check("IntAdd")
	m = IntIncrement(words)
	check("IntIncrement")
	m = Pointer(words)
	check("Pointer")
	m = Index(words)
	check("Index")
	m = IntAdd_unsafe(words)
	check("IntAdd_unsafe")
	m = IntIncrement_unsafe(words)
	check("IntIncrement_unsafe")
	m = Pointer_unsafe(words)
	check("Pointer_unsafe")
	m = Index_unsafe(words)
	check("Index_unsafe")
}

func IntAdd(words [][]byte) map[string]int {
	var m = make(map[string]int)
	for _, w := range words {
		m[string(w)] = m[string(w)] + 1
		
	}
	return m
}

func IntIncrement(words [][]byte) map[string]int {
	var m = make(map[string]int)
	for _, w := range words {
		m[string(w)]++
	}
	return m
}

func Pointer(words [][]byte) map[string]int {
	var m2 = make(map[string]*int)
	for _, w := range words {
		if v := m2[string(w)]; v != nil {
			*v++
		} else {
			v = new(int)
			*v = 1
			m2[string(w)] = v
		}
	}
	var m = make(map[string]int, len(m2))
	for k, p := range m2 {
		m[k] = *p
	}
	return m
}

func Index(words [][]byte) map[string]int {
	var m = make(map[string]int)
	var stats []int
	for _, w := range words {
		if v, ok := m[string(w)]; ok {
			stats[v]++
		} else {
			stats = append(stats, 1)
			m[string(w)] = len(m)
		}
	}
	for k, c := range m {
		m[k] = stats[c]
	}
	return m
}

var mAdd map[string]int
func Benchmark_IntAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mAdd = IntAdd(words)
	}
}

var mIncrement map[string]int
func Benchmark_Increment(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mIncrement = IntIncrement(words)
	}
}

var mPointer map[string]int
func Benchmark_Pointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mPointer = Pointer(words)
	}
}

var mIndex map[string]int
func Benchmark_Index(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mIndex = Index(words)
	}
}



func IntAdd_unsafe(words [][]byte) map[string]int {
	var m = make(map[string]int)
	for _, w := range words {
		m[str(w)] = m[str(w)] + 1
	}
	return m
}

func IntIncrement_unsafe(words [][]byte) map[string]int {
	var m = make(map[string]int)
	for _, w := range words {
		m[str(w)]++
	}
	return m
}

func Pointer_unsafe(words [][]byte) map[string]int {
	var m2 = make(map[string]*int)
	for _, w := range words {
		if v := m2[str(w)]; v != nil {
			*v++
		} else {
			v = new(int)
			*v = 1
			m2[str(w)] = v
		}
	}
	var m = make(map[string]int, len(m2))
	for k, p := range m2 {
		m[k] = *p
	}
	return m
}

func Index_unsafe(words [][]byte) map[string]int {
	var m = make(map[string]int)
	var stats []int
	for _, w := range words {
		if v, ok := m[str(w)]; ok {
			stats[v]++
		} else {
			stats = append(stats, 1)
			m[str(w)] = len(m)
		}
	}
	for k, c := range m {
		m[k] = stats[c]
	}
	return m
}

var mAdd_unsafe map[string]int
func Benchmark_IntAdd_unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mAdd_unsafe = IntAdd_unsafe(words)
	}
}

var mIncrement_unsafe map[string]int
func Benchmark_IntIncrement_unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mIncrement_unsafe = IntIncrement_unsafe(words)
	}
}

var mPointer_unsafe map[string]int
func Benchmark_Pointer_unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mPointer_unsafe = Pointer_unsafe(words)
	}
}

var mIndex_unsafe map[string]int
func Benchmark_Index_unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mIndex_unsafe = Index_unsafe(words)
	}
}


