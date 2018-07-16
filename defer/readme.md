
### About

This benchmark is to check how defer will affect program efficiency.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
Benchmark_DefersInLoop-4        	   20000	     85186 ns/op
Benchmark_DefersInFunc-4        	   20000	     83377 ns/op
Benchmark_NoDeferButHasFunc-4   	  300000	      4158 ns/op
Benchmark_NoDeferNoFunc-4       	 1000000	      1301 ns/op
Benchmark_MutexDefer-4          	20000000	        88.7 ns/op
Benchmark_MutexNoDeferA-4       	50000000	        27.4 ns/op
Benchmark_MutexNoDeferB-4       	50000000	        27.6 ns/op
```

`defer` hurts performance, at least at the time of Go 1.11.
However, we should give up using `defer`, for
1. deferred calls make code robust for some scenarios.
1. some defer uses may be optimized later.
