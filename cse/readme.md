
### About

This benchmark is to check if the SSA CSE (common subexpression elimination)
optimization works.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_CSE_Case_1a-4   	2000000000	         0.88 ns/op
Benchmark_CSE_Case_1b-4   	2000000000	         0.88 ns/op
Benchmark_CSE_Case_2a-4   	100000000	        16.7 ns/op
Benchmark_CSE_Case_2b-4   	100000000	        16.7 ns/op
```

It looks CSE works.


