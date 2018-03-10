package main

import (
	"testing"
)

var c1 = make(chan int)
var c2 = make(chan int)

func Benchmark_TwoTrySends(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case c1 <- 1:
		default:
		}
		
		select {
		case c2 <- 1:
		default:
		}
	}
}

func Benchmark_TwoTryReceives(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <- c1:
		default:
		}
		
		select {
		case <- c2:
		default:
		}
	}
}

func Benchmark_OneSelectWithTwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <- c1:
		case c2 <- 1:
		default:
		}
	}
}










