
### About

This benchmark is to compare different ways to stat word count.

### Result

```
$ go version
go version go1.16.2 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/count-word-uses
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_f-4   	     687	   1699257 ns/op
Benchmark_g-4   	    1898	    618261 ns/op
Benchmark_h-4   	    1892	    623414 ns/opp
```

Why the `f` function is slow? The reason is `string(byte_slice)` doesn't make a
duplication when it appears in map element getters, but it makes a duplication
if it appeans in map element setters.
`m[string(w)]++` is equivalent to `m[string(w)] = m[string(w)] + 1`, there is a
map element setter and a map element getter in it. The `g` and `h` functions
doesn't use map element setters as frequently as the function `f`.

(ref: https://github.com/golang/go/issues/45021)


