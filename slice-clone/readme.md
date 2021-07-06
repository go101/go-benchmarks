
### About

This benchmark is to check the efficiency differences between different slice clone implementations.


### Results

For `const N = 33333`:

```
$ go version
go version go1.16.3 linux/amd64
$ go test -bench=. -benchmem -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/slice-clone
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_MakeCopy-4     	   73574	     51496 ns/op	  270336 B/op	       1 allocs/op
Benchmark_MakeAppend-4   	   45998	     73042 ns/op	  270336 B/op	       1 allocs/op
Benchmark_AppendA-4      	   61928	     50016 ns/op	  270336 B/op	       1 allocs/op
Benchmark_AppendB-4      	   71398	     52129 ns/op	  270336 B/op	       1 allocs/op
Benchmark_Verbose-4      	   63339	     53325 ns/op	  270336 B/op	       1 allocs/op
```


For `const N = 99`:

```
$ go version
go version go1.16.3 linux/amd64
$ go test -bench=. -benchmem -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/slice-clone
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_MakeCopy-4     	15266563	       258.9 ns/op	     896 B/op	       1 allocs/op
Benchmark_MakeAppend-4   	13487722	       255.4 ns/op	     896 B/op	       1 allocs/op
Benchmark_AppendA-4      	13766970	       261.5 ns/op	     896 B/op	       1 allocs/op
Benchmark_AppendB-4      	13709266	       255.6 ns/op	     896 B/op	       1 allocs/op
Benchmark_Verbose-4      	15315482	       235.5 ns/op	     896 B/op	       1 allocs/op
```
