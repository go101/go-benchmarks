
### About

This benchmark is to compare the costs between one-line and verbose
implementations for some specified slice manipulations.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.3 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
Benchmark_SliceUnshift-4              	10000000	       215 ns/op	     152 B/op	       2 allocs/op
Benchmark_SliceInsertOne_OneLine-4    	 5000000	       384 ns/op	     352 B/op	       2 allocs/op
Benchmark_SliceInsertOne_Verbose-4    	10000000	       165 ns/op	     144 B/op	       1 allocs/op
Benchmark_SliceGrow_OneLine-4         	 2000000	       836 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceGrow_VerboseAppend-4   	 2000000	       832 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceGrow_VerboseCopy-4     	 3000000	       406 ns/op	     896 B/op	       1 allocs/op
Benchmark_SliceInsertion_OneLine-4    	 5000000	       386 ns/op	     368 B/op	       2 allocs/op
Benchmark_SliceInsertion_Verbose-4    	10000000	       169 ns/op	     160 B/op	       1 allocs/op
Benchmark_SliceMemclr_Loop-4          	 5000000	       232 ns/op	       0 B/op	       0 allocs/op
Benchmark_SliceMemclr_Append-4        	  500000	      3472 ns/op	    8192 B/op	       1 allocs/op

$ go version
go version go1.11 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
Benchmark_SliceUnshift-4              	10000000	       212 ns/op	     152 B/op	       2 allocs/op
Benchmark_SliceInsertOne_OneLine-4    	 5000000	       380 ns/op	     352 B/op	       2 allocs/op
Benchmark_SliceInsertOne_Verbose-4    	10000000	       162 ns/op	     144 B/op	       1 allocs/op
Benchmark_SliceGrow_OneLine-4         	 3000000	       488 ns/op	     896 B/op	       1 allocs/op
Benchmark_SliceGrow_VerboseAppend-4   	 2000000	       848 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceGrow_VerboseCopy-4     	 3000000	       414 ns/op	     896 B/op	       1 allocs/op
Benchmark_SliceInsertion_OneLine-4    	 3000000	       389 ns/op	     368 B/op	       2 allocs/op
Benchmark_SliceInsertion_Verbose-4    	10000000	       182 ns/op	     160 B/op	       1 allocs/op
Benchmark_SliceMemclr_Loop-4          	10000000	       228 ns/op	       0 B/op	       0 allocs/op
Benchmark_SliceMemclr_Append-4        	10000000	       224 ns/op	       0 B/op	       0 allocs/op

```

It looks, since Go 1.11, `append(s, make([]T, n)...)` is optimized by avoiding one allocation.
However, this optimization doesn't bring any new benefits  for Go.
For two possible scenarios which may be benefited from this optimization, slice grow and slice segment clear,
there had been already other same-efficency alternative solutions existed before Go 1.11.

In fact, the most needed optimization is `append(sa, append(sb, sc...)...)`,
which has not been optimized in Go 1.11.


