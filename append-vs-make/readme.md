
### About

This benchmark is to check the efficiency differences between
using `append` and `make` to create slices.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    2000	    535285 ns/op
Benchmark_PureMake-4      	    1000	   1501547 ns/op
Benchmark_MakeAndCopy-4   	    1000	   2168468 ns/op
Benchmark_Append-4        	    2000	   1001056 ns/op
```

If `N > 10000`,
it looks create a slice by using `append` is more efficient than
by using `copy`. This is some counter-intuitive.
It is more surprised that a pure `make` call is even slower than an `append` call.


