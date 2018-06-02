package main

import (
	"testing"
)

var va int
func ReceiveRemainingValues_CheckLen(c <-chan int) {
	for len(c) > 0 {
		va = <-c
	}
}

var vb int
func ReceiveRemainingValues_Select(c <-chan int) {
	for {
		select {
		case vb = <-c:
		default:
			return
		}
	}
}

const N = 10000
func createBufferChannel() <-chan int {
	c := make(chan int, N)
	for i := 0; i < N; i++ {
		c <- i
	}
	return c
}

func Benchmark_CheckLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := createBufferChannel()
		ReceiveRemainingValues_CheckLen(c)
	}
}

func Benchmark_Select(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := createBufferChannel()
		ReceiveRemainingValues_Select(c)
	}
}


