
### About

This benchmark is to benchmark 
* the `Int63` function/method between new created and global `Rand` instances.
* the `Read` functions between `math/rand` and `crypto/rand` standard packages.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.2 linux/amd64
$ go test -bench=.
M = 8, N = 8192
goos: linux
goarch: amd64
Benchmark_MathUnsyncedInt63-4           	100000000	        14.2 ns/op
Benchmark_MathGlobalInt63-4             	30000000	        35.8 ns/op
Benchmark_MathUnsyncedRandRead_8-4      	30000000	        42.6 ns/op
Benchmark_MathGlobalRandRead_8-4        	20000000	        64.9 ns/op
Benchmark_CryptoRandRead_8-4            	 3000000	       558 ns/op
Benchmark_MathUnsyncedRandRead_8192-4   	   50000	     29164 ns/op
Benchmark_MathGlobalRandRead_8192-4     	   50000	     26535 ns/op
Benchmark_CryptoRandRead_8192-4         	   30000	     41677 ns/oppp
```

Generally, an unsynced `math/rand.Rand` is faster for the global one,
except for calling `Read(p []byte) (n int, err error)` method/function.
The exception is strange.

For `Read(p []byte) (n int, err error)` method/function, `math/rand`
is faster than `crypto/rand`, but the differences become smaller
when the lengths of the input slices become larger.

