
### About

This benchmark is to compare different ways to stat word count.

### Result

```
$ go version
go version go1.16.2 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/count-word-uses
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_IntAdd-4                	     523	   2239815 ns/op	 1928080 B/op	   10009 allocs/op
Benchmark_Increment-4             	     704	   1699259 ns/op	 1928071 B/op	   10009 allocs/op
Benchmark_Pointer-4               	    1862	    614622 ns/op	   32043 B/op	     211 allocs/op
Benchmark_Index-4                 	    1864	    622641 ns/op	   29317 B/op	     117 allocs/op
Benchmark_IntAdd_unsafe-4         	    1036	   1126933 ns/op	    8089 B/op	       9 allocs/op
Benchmark_IntIncrement_unsafe-4   	    1837	    616928 ns/op	    8070 B/op	       9 allocs/op
Benchmark_Pointer_unsafe-4        	    1832	    596916 ns/op	   12851 B/op	     111 allocs/op
Benchmark_Index_unsafe-4          	    1884	    607870 ns/op	   10112 B/op	      17 allocs/op
```

Why is the `Increment` function slower than `Pointer` and `Index`?
The reason is `string(byte_slice)` doesn't make a duplication
when it appears in map element getters (an optimization made by the
official Go compiler), but it makes a duplication if it appears in map element setters.
`m[string(w)]++` is equivalent to `m[string(w)] = m[string(w)] + 1` (from logic.
But another compiler optimization makes their performance different), there is a
map element setter and a map element getter in it (in fact, the getter only exists
in `IntAdd` but not `Increment`). The `Pointer` and `Index` functions
don't use map element setters as frequently as the function `Increment`.

(ref: https://github.com/golang/go/issues/45021)

More `[]byte <-> string` conversion optimizations made by the official Go compiler: https://go101.org/article/string.html#conversion-optimizations


