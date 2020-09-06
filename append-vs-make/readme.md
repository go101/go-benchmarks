
### About

This benchmark is to check the efficiency differences between
using `append` and `make` to create slices.

**_(Please note that, since Go Toolchain 1.15, the benchmark result has changed much.
`make+copy` calls are specially optimized so that they are more efficient than the `append` call.)_**

### Results

All results are for `const N = 1024 * 1024`.

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.11beta2 linux/amd64
```


For `type Element = int64`:

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: AMD Ryzen 3 2200G with Radeon Vega Graphics
$ go version
go version go1.15 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    1921	    613320 ns/op
Benchmark_PureMake-4      	    1305	    858066 ns/op
Benchmark_MakeAndCopy-4   	    1518	    753634 ns/op
Benchmark_Append-4        	    1522	    765014 ns/op
```

**_(The following is not valid since Go Toolchain 1.15.)_**

For `type Element = int64`:

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
It is more surprising that a pure `make` call is even slower than an `append` call.

However, it looks the results vary much for different kinds of element types.
The result differences might related to write barriers.
The more pointer fields in the element type, the slower the append way.
Related issues:
* https://github.com/golang/go/issues/26734
* https://github.com/golang/go/issues/23906

For `type Element = []int`:

```
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    1000	   1610359 ns/op
Benchmark_PureMake-4      	     100	  13603956 ns/op
Benchmark_MakeAndCopy-4   	     100	  21189024 ns/op
Benchmark_Append-4        	     100	  18077256 ns/op
```

For `type Element = func()`:

```
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    2000	    529595 ns/op
Benchmark_PureMake-4      	     300	   4110150 ns/op
Benchmark_MakeAndCopy-4   	     100	  12242049 ns/op
Benchmark_Append-4        	     200	   7151174 ns/op
```

For `type Element = chan int`:

```
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    2000	    529595 ns/op
Benchmark_PureMake-4      	     300	   4110150 ns/op
Benchmark_MakeAndCopy-4   	     100	  12242049 ns/op
Benchmark_Append-4        	     200	   7151174 ns/op
```


For `type Element = string`:

```
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    2000	   1077957 ns/op
Benchmark_PureMake-4      	     200	  10894784 ns/op
Benchmark_MakeAndCopy-4   	     100	  15225871 ns/op
Benchmark_Append-4        	     100	  16427054 ns/op
```


For `type Element = interface{}`:

```
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PureCopy-4      	    2000	   1060068 ns/op
Benchmark_PureMake-4      	     200	   9609967 ns/op
Benchmark_MakeAndCopy-4   	     100	  12714693 ns/op
Benchmark_Append-4        	     100	  13150424 ns/o
```

