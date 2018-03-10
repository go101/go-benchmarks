package main

import (
	"reflect"
	"testing"
)

func CompareSlices_Reflect(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}

func CompareSlices_General(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func CompareSlices_BCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)] // this line is the key
	for i, v := range a {
		if v != b[i] { // here is no bounds checking for b[i]
			return false
		}
	}

	return true
}

func CompareSlices_Pointer(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	if &a[0] == &b[0] {
		return true // early exit
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func CompareSlices_PointerAndBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	b = b[:len(a)]
	if &a[0] == &b[0] {
		return true
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

type _case struct {
	x, y []int
}

var s0a = []int{0, 1, 2, 3, 4}
var s0b = []int{0, 1, 2, 3, 4}
var s1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
var s2 = make([]int, 100)
var s3 = make([]int, 100)

func init() {
	s3[99] = 99
}

var cases = []_case{
	_case{
		x: nil,
		y: nil,
	},
	_case{
		x: nil,
		y: []int{},
	},
	_case{
		x: []int{},
		y: []int{},
	},
	_case{
		x: s0a,
		y: s0a,
	},
	_case{
		x: s0a,
		y: s0b,
	},
	_case{
		x: s1,
		y: s1,
	},
	_case{
		x: s1,
		y: s2,
	},
	_case{
		x: s2,
		y: s3,
	},
}

var r bool
func Benchmark_SliceComparison(b *testing.B) {
	type benchmark struct {
		name string
		f    func(a, b []int) bool
	}
	benchmarks := []benchmark{
		{"Reflect", CompareSlices_Reflect},
		{"General", CompareSlices_General},
		{"BCE", CompareSlices_BCE},
		{"Pointer", CompareSlices_Pointer},
		{"PointerAndBCE", CompareSlices_PointerAndBCE},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, cs := range cases {
					r = bm.f(cs.x, cs.y)
				}
			}
		})
	}
}
