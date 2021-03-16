
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
Benchmark_f-4          	     705	   1715754 ns/op	 1928089 B/op	   10009 allocs/op
Benchmark_g-4          	    1851	    626779 ns/op	   32043 B/op	     211 allocs/op
Benchmark_h-4          	    1826	    633126 ns/op	   29311 B/op	     117 allocs/op
Benchmark_f_unsafe-4   	    1881	    627144 ns/op	    8075 B/op	       9 allocs/op
Benchmark_g_unsafe-4   	    1832	    607550 ns/op	   12845 B/op	     111 allocs/op
Benchmark_h_unsafe-4   	    1872	    616787 ns/op	   10116 B/op	      17 allocs/op
```

Why the `f` function is slow? The reason is `string(byte_slice)` doesn't make a
duplication when it appears in map element getters (an optimization made by the
official Go compiler), but it makes a duplication if it appears in map element setters.
`m[string(w)]++` is equivalent to `m[string(w)] = m[string(w)] + 1`, there is a
map element setter and a map element getter in it. The `g` and `h` functions
doesn't use map element setters as frequently as the function `f`.

(ref: https://github.com/golang/go/issues/45021)

More `[]byte <-> string` conversion optimizations made by the official Go compiler: https://go101.org/article/string.html#conversion-optimizations


