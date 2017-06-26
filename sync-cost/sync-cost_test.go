package main

import (
	"testing"
	"sync/atomic"
	"sync"
)

var g int32

func Benchmark_NoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g++
	}
}

func Benchmark_Atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&g, 1)
	}
}

var m sync.Mutex
func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		g++
		m.Unlock()
	}
}

var ch0 = make(chan struct{}, 1)
func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch0 <- struct{}{}
		g++
		<-ch0
	}
}

var ch1 = make(chan struct{}, 1)
var ch2 = make(chan struct{}, 1)
var ch3 = make(chan struct{}, 1)

func Benchmark_Select_OneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			g++
			<-ch1
		}
	}
}

func Benchmark_Select_TwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			g++
			<-ch1
		case ch2 <- struct{}{}:
			g++
			<-ch2
		}
	}
}

func Benchmark_Select_TwoCases_Plus_TrySent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case chan int(nil) <- 1:
		default:
		}
		
		select {
		case ch1 <- struct{}{}:
			g++
			<-ch1
		case ch2 <- struct{}{}:
			g++
			<-ch2
		}
	}
}

func Benchmark_Select_TwoCases_Plus_TryReceive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-chan int(nil):
		default:
		}
		
		select {
		case ch1 <- struct{}{}:
			g++
			<-ch1
		case ch2 <- struct{}{}:
			g++
			<-ch2
		}
	}
}

func Benchmark_Select_ThreeCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			g++
			<-ch1
		case ch2 <- struct{}{}:
			g++
			<-ch2
		case ch3 <- struct{}{}:
			g++
			<-ch3
		}
	}
}








