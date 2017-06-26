
### About

This benchmark is to prove CPU cache matters much.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
Benchmark_SumArrayElements_GlobalVar-4   	   10000	    204523 ns/op
Benchmark_SumArrayElements_LocalVar-4    	   30000	     57926 ns/op
```

The performance difference is large.


