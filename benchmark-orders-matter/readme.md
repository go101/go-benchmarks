
### About

This benchmark is to prove the benchmark orders matter sontimes.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64

$ cd 1; go test -bench=.
Benchmark_RangeSlice_TwoIterationVar-4   	30000000	        42.2 ns/op
Benchmark_RangeSlice_OneIterationVar-4   	50000000	        31.3 ns/op
Benchmark_LoopSlice-4                    	30000000	        44.1 ns/op

$ cd ../2; go test -bench=.
Benchmark_LoopSlice-4                    	30000000	        42.5 ns/op
Benchmark_RangeSlice_TwoIterationVar-4   	50000000	        29.3 ns/op
Benchmark_RangeSlice_OneIterationVar-4   	30000000	        41.9 ns/op

$ cd ../3; go test -bench=.
Benchmark_RangeSlice_OneIterationVar-4   	30000000	        42.0 ns/op
Benchmark_LoopSlice-4                    	50000000	        30.8 ns/op
Benchmark_RangeSlice_TwoIterationVar-4   	30000000	        42.4 ns/op
```

You will find that the benchmark in the middle always looks most efficient.
But which is really most efficient on earth?

The following are the results by disabling CPU frequency scaling.
The results are a bit different, the middle one is not always the most performant.
But orders are still matter.
It looks, the results of 1 and 3 are the same, but 2 is different.

```
# for i in /sys/devices/system/cpu/cpu[0-3]
> do
> echo performance > $i/cpufreq/scaling_governor
> done
# exit
$
$ go version
go version go1.11beta1 linux/amd64
$
$ cd 1
$ go test -bench=. -benchtime=3s
goos: linux
goarch: amd64
Benchmark_RangeSlice_TwoIterationVar-4   	200000000	        29.3 ns/op
Benchmark_RangeSlice_OneIterationVar-4   	100000000	        30.8 ns/op
Benchmark_LoopSlice-4                    	100000000	        51.0 ns/op
$
$ cd ../2
$ go test -bench=. -benchtime=3s
goos: linux
goarch: amd64
Benchmark_LoopSlice-4                    	100000000	        30.7 ns/op
Benchmark_RangeSlice_TwoIterationVar-4   	100000000	        43.3 ns/op
Benchmark_RangeSlice_OneIterationVar-4   	100000000	        50.9 ns/op
$
$ cd ../3
$ go test -bench=. -benchtime=3s
goos: linux
goarch: amd64
Benchmark_RangeSlice_OneIterationVar-4   	100000000	        30.7 ns/op
Benchmark_LoopSlice-4                    	100000000	        51.0 ns/op
Benchmark_RangeSlice_TwoIterationVar-4   	200000000	        29.3 ns/op
```
