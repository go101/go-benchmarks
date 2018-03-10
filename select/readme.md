
### About

This benchmark is to compare the cost between two try-send/receive select blocks
and one two-cases select block.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.9 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_TwoTrySends-4             	50000000	        22.3 ns/op
Benchmark_TwoTryReceives-4          	100000000	        18.1 ns/op
Benchmark_OneSelectWithTwoCases-4   	10000000	       203 ns/opp
```

It looks try-send and try-receive select blocks are highly optimized.


