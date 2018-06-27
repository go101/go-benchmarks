
### About

This benchmark is to compare the costs between one-line and verbose
implementations for slice grow and insertion.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.3 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
Benchmark_SliceGrow_OneLine-4        	 2000000	       678 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceGrow_Verbose-4        	 2000000	       744 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceInsertion_OneLine-4   	 5000000	       353 ns/op	     368 B/op	       2 allocs/op
Benchmark_Verbose-4                  	10000000	       161 ns/op	     160 B/op	       1 allocs/op

$ go version
go version go1.11 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
Benchmark_SliceGrow_OneLine-4        	 3000000	       481 ns/op	     896 B/op	       1 allocs/op
Benchmark_SliceGrow_Verbose-4        	 2000000	       829 ns/op	    1600 B/op	       2 allocs/op
Benchmark_SliceInsertion_OneLine-4   	 5000000	       377 ns/op	     368 B/op	       2 allocs/op
Benchmark_Verbose-4                  	10000000	       187 ns/op	     160 B/op	       1 allocs/op

```

It looks one-line grow is optimized since Go 1.11.


