package main

import (
	"testing"
)

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
	m = f(words)
	check("f")
	m = g(words)
	check("g")
	m = h(words)
	check("h")
}

func f(words [][]byte) map[string]int {
	var m = make(map[string]int)
	for _, w := range words {
		m[string(w)]++
	}
	return m
}

func g(words [][]byte) map[string]int {
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

func h(words [][]byte) map[string]int {
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

var mf map[string]int
func Benchmark_f(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mf = f(words)
	}
}

var mg map[string]int
func Benchmark_g(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mg = g(words)
	}
}

var mh map[string]int
func Benchmark_h(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mh = h(words)
	}
}


