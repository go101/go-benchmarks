
### About

This benchmark is to benchmark 
* the `Int63` function/method between new created and global `Rand` instances.
* the `Read` functions between `math/rand` and `crypto/rand` standard packages.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.9 linux/amd64
$ go test -bench=.
M = 8, N = 8192
goos: linux
goarch: amd64
Benchmark_MathUnsyncedInt63-4           	100000000	        13.8 ns/op
Benchmark_MathGlobalInt63-4             	30000000	        36.2 ns/op
Benchmark_MathUnsyncedRandRead_8-4      	30000000	        42.5 ns/op
Benchmark_MathGlobalRandRead_8-4        	20000000	        64.3 ns/op
Benchmark_CryptoRandRead_8-4            	 3000000	       580 ns/op
Benchmark_MathUnsyncedRandRead_8192-4   	   50000	     28927 ns/op
Benchmark_MathGlobalRandRead_8192-4     	   50000	     26779 ns/op
Benchmark_CryptoRandRead_8192-4         	   30000	     44432 ns/opp
```

Generally, an unsynced `math/rand.Rand` is faster for the global one,
except for calling `Read(p []byte) (n int, err error)` method/function.

For `Read(p []byte) (n int, err error)` method/function, `math/rand`
is faster than `crypto/rand`, but the differences become smaller
when the lengths of the input slices become larger.

