
### About

This benchmark is to prove the benchmark orders matter sontimes.

### Result

```
$ go version
go version go1.17beta1 linux/amd64

$ cd 2-benchmarks; go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/benchmark-count-matters/2-benchmarks
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_CopyStruct_3_fields-4   	129147345	         9.655 ns/op
Benchmark_CopyStruct_4_fields-4   	131043462	         8.973 ns/op
PASS
ok  	github.com/go101/go-benchmarks/benchmark-count-matters/2-benchmarks	4.292s

$ cd ../3-benchmarks; go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/benchmark-count-matters/3-benchmarks
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_CopyStruct_3_fields-4   	71751447	        16.04 ns/op
Benchmark_CopyStruct_4_fields-4   	129615433	         9.951 ns/op
Benchmark_xxx-4                   	1000000000	         0.3997 ns/op
```

You will find that the results of the benchmark `Benchmark_CopyStruct_3` depends
on the benchmark `Benchmark_xxx` exists.

