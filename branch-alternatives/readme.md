
### About

This benchmark is to compare different ways to do branch executions.

### Result

```
$ go version
go version go1.16.2 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/branch-alternatives
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_ifelse-4       	1000000000	         0.7660 ns/op
Benchmark_boolmap-4      	19130954	        62.00 ns/op
Benchmark_sliceindex-4   	1000000000	         0.7596 ns/op
```

Conclusion: maps with boolean keys are not specially optimized yet.

(ref: https://github.com/golang/go/issues/44578)

