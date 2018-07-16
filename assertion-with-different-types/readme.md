
### About

This benchmark is to check the efficiency of type assertion with interface value
with all kinds of types.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.11beta1 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_PointerAssert-4       	2000000000	         1.75 ns/op
Benchmark_PointerAssign-4       	2000000000	         1.31 ns/op
Benchmark_IntAssert-4           	2000000000	         1.31 ns/op
Benchmark_IntAssign-4           	2000000000	         0.88 ns/op
Benchmark_BoolAssert-4          	2000000000	         1.34 ns/op
Benchmark_BoolAssign-4          	2000000000	         0.87 ns/op
Benchmark_ArrayAssert-4         	100000000	        20.2 ns/op
Benchmark_ArrayAssign-4         	300000000	         5.74 ns/op
Benchmark_BlankStructAssert-4   	2000000000	         0.87 ns/op
Benchmark_BlankStructAssign-4   	2000000000	         0.44 ns/op
Benchmark_SliceAssert-4         	1000000000	         2.28 ns/op
Benchmark_SliceAssign-4         	2000000000	         1.75 ns/op
```

For small-size values, an assignment through assertion needs
about two times time of a general assignment. Not bad.


