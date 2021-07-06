
### About

This benchmark is to check the efficiency of type assertion with interface value
with all kinds of types and boxing values into interfaces.

### Result

```
$ go version
go version go1.16.3 linux/amd64
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/box-value-into-interface
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_BoxPointer-4          	1000000000	         1.037 ns/op	       0 B/op	       0 allocs/op
Benchmark_PointerAssert-4       	911917486	         1.235 ns/op	       0 B/op	       0 allocs/op
Benchmark_PointerAssign-4       	1000000000	         0.9384 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxInt-4              	54424342	        23.97 ns/op	       8 B/op	       1 allocs/op
Benchmark_IntAssert-4           	1000000000	         0.7866 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxSmallInt-4         	326547188	         3.735 ns/op	       0 B/op	       0 allocs/op
Benchmark_SmallIntAssert-4      	1000000000	         0.8294 ns/op	       0 B/op	       0 allocs/op
Benchmark_IntAssign-4           	1000000000	         0.5607 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxByte-4             	1000000000	         1.169 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteAssert-4          	1000000000	         0.7922 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteAssign-4          	1000000000	         0.8156 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxBool-4             	962966514	         1.179 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoolAssert-4          	1000000000	         0.7701 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoolAssign-4          	1000000000	         0.7766 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxArray-4            	18964473	        61.61 ns/op	     128 B/op	       1 allocs/op
Benchmark_ArrayAssert-4         	79327922	        14.53 ns/op	       0 B/op	       0 allocs/op
Benchmark_ArrayAssign-4         	285178418	         4.205 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxBlankStruct-4      	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
Benchmark_BlankStructAssert-4   	1000000000	         0.5995 ns/op	       0 B/op	       0 allocs/op
Benchmark_BlankStructAssign-4   	1000000000	         0.3749 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxSlice-4            	26991988	        42.72 ns/op	      24 B/op	       1 allocs/op
Benchmark_SliceAssert-4         	769459327	         1.517 ns/op	       0 B/op	       0 allocs/op
Benchmark_SliceAssign-4         	1000000000	         1.126 ns/op	       0 B/op	       0 allocs/op
Benchmark_BoxFunction-4         	1000000000	         1.013 ns/op	       0 B/op	       0 allocs/op
Benchmark_FunctionAssert-4      	1000000000	         1.124 ns/op	       0 B/op	       0 allocs/op
Benchmark_FunctionAssign-4      	1000000000	         0.9361 ns/op	       0 B/op	       0 allocs/op
```

* Boxing pointers is fast.
* Boxing maps/channels/functions is also fast, for they are implemented as pointers internally.
* Boxing zero-size values, bytes and booleans are also fast.
* Boxing small integers within [0, 255] is not slow.


