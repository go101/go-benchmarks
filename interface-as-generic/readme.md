
### About

This benchmark is to check the efficiency of type assertion with interface value
with all kinds of types and boxing values into interfaces.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.11beta1 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_boxPointer-4          	2000000000	         1.33 ns/op
Benchmark_PointerAssert-4       	2000000000	         1.75 ns/op
Benchmark_PointerAssign-4       	2000000000	         1.31 ns/op
Benchmark_BoxInt-4              	300000000	         5.00 ns/op
Benchmark_IntAssert-4           	2000000000	         1.32 ns/op
Benchmark_IntAssign-4           	2000000000	         0.88 ns/op
Benchmark_BoxBool-4             	2000000000	         1.76 ns/op
Benchmark_BoolAssert-4          	2000000000	         1.32 ns/op
Benchmark_BoolAssign-4          	2000000000	         0.88 ns/op
Benchmark_BoxArray-4            	10000000	       127 ns/op
Benchmark_ArrayAssert-4         	50000000	        20.2 ns/op
Benchmark_ArrayAssign-4         	300000000	         5.69 ns/op
Benchmark_BoxBlankStruct-4      	2000000000	         1.31 ns/op
Benchmark_BlankStructAssert-4   	2000000000	         0.87 ns/op
Benchmark_BlankStructAssign-4   	2000000000	         0.44 ns/op
Benchmark_BoxSlice-4            	20000000	        98.5 ns/op
Benchmark_SliceAssert-4         	1000000000	         2.19 ns/op
Benchmark_SliceAssign-4         	2000000000	         1.75 ns/op

```

For small-size values, an assignment through assertion needs
about two times time of a general assignment. Not bad (in fact, quite good).
However, when assign a non-pointer non-interface value to an interface value is quite slow.
So if using interface as generic, try to use pointer values as interface dynamic values.


