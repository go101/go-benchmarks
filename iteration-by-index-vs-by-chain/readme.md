
### About

This benchmark is to check the efficiency differences between
iteration by index and by chain.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_ByIndex-4   	 1000000	      1356 ns/op
Benchmark_ByChain-4   	  500000	      2703 ns/op
```

The efficiency of iteration by index is higher than by chain.


