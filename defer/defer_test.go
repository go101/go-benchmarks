package main

import (
	"testing"
	"sync"
)

const N = 1000
var g int

func Benchmark_DefersInLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g = func() (x int) {
			for k := 0; k < N; k++ {
				defer func() {
					x++
				}()
				x += k
			}
			return x
		}()
	}
}

func Benchmark_DefersInFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g = func() (x int) {
			for k := 0; k < N; k++ {
				func() {
					defer func() {
						x++
					}()
					x += k
				}()
			}
			return x
		}()
	}
}

func Benchmark_NoDeferButHasFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g = func() (x int) {
			for k := 0; k < N; k++ {
				func() {
					x += k
					x++
				}()
			}
			return x
		}()
	}
}


func Benchmark_NoDeferNoFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g = func() (x int) {
			for k := 0; k < N; k++ {
				x += k
				x++
			}
			return x
		}()
	}
}

var ma, mb, mc sync.Mutex

func Benchmark_MutexDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			ma.Lock()
			defer ma.Unlock()
			g++
		}()
	}
}

func Benchmark_MutexNoDeferA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			mb.Lock()
			p := &mb
			g++
			p.Unlock()
		}()
	}
}

func Benchmark_MutexNoDeferB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			mc.Lock()
			g++
			mc.Unlock()
		}()
	}
}




/*
$ go test -bench=.
Benchmark_Defer1-4          	   20000	     86562 ns/op
Benchmark_Defer2-4          	   20000	     85333 ns/op
Benchmark_Defer2b-4         	  300000	      4136 ns/op
Benchmark_Defer3-4          	 1000000	      1295 ns/op
Benchmark_MutexDefer-4      	20000000	        89.9 ns/op
Benchmark_MutexNoDeferA-4   	50000000	        27.5 ns/op
Benchmark_MutexNoDeferB-4   	50000000	        27.5 ns/op
*/