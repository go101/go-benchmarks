
### About

This benchmark is to show a weird inline becomes inefficient case.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
Benchmark_GlobalArray_NoInline-4                   	  500000	      3946 ns/op
Benchmark_GlobalArray_Inline-4                     	  300000	      5227 ns/op
Benchmark_GlobalArray_Inline_AnonymousFunction-4   	  300000	      5230 ns/op
Benchmark_LocalArray_NoInline-4                    	  500000	      3963 ns/op
Benchmark_LocalArray_Inline-4                      	  300000	      4144 ns/op
Benchmark_LocalArray_Inline_AnonymousFunction-4    	  500000	      3935 ns/op


```

When to sum of elements of global array, inline is inefficient.


