
### About

This benchmark is to compare the efficiencies of different slice comparison
implementations, including a way making use of the BCE
(bounds check elimination) optimizations.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_SliceComparison/Reflect-4         	  100000	     21601 ns/op
Benchmark_SliceComparison/General-4         	 5000000	       287 ns/op
Benchmark_SliceComparison/BCE-4             	 5000000	       251 ns/op
Benchmark_SliceComparison/Pointer-4         	 5000000	       257 ns/op
Benchmark_SliceComparison/PointerAndBCE-4   	10000000	       215 ns/op
```

It looks the reflect way is very inefficient,
and the BCE and early pointer comparison really take effect.

NOTE: since Go SDK 1.11, the standard Go compiler becomes more smarter
so that now the General version is equivalent to the BCE version,
and the same is for the Pointer and PointerAndBCE versions.

```
$ go version
go version go1.11 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_SliceComparison/Reflect-4         	  100000	     20585 ns/op
Benchmark_SliceComparison/General-4         	10000000	       214 ns/op
Benchmark_SliceComparison/BCE-4             	10000000	       214 ns/op
Benchmark_SliceComparison/Pointer-4         	10000000	       200 ns/op
Benchmark_SliceComparison/PointerAndBCE-4   	10000000	       199 ns/op
```
